package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/signal"
	"strings"
	"time"

	sr "github.com/TheSmallBoat/carlo/streaming_rpc"
	"github.com/TheSmallBoat/fiesta"
)

const (
	ServicePublish   = "Publish-Service"
	ServiceSubscribe = "Subscribe-Service"

	ActionHeader    = "Action"
	TopicHeader     = "Topic"
	QosHeader       = "Qos"
	BodyTitleHeader = "Body-Title"

	ActionPublish   = "Publish"
	ActionSubscribe = "Subscribe"

	ResponsePublishHeader   = "Pub-Response"
	ResponseSubscribeHeader = "Sub-Response"

	ResponseSuccess = "success"
	ResponseFailure = "failure"
)

var topicMessageChannel = make(chan message, 5)
var topicSubscriber = map[string]func(){}

type message struct {
	topic string
	body  string
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func publishHandler(ctx *sr.Context) {
	now := time.Now()
	buf, err := ioutil.ReadAll(ctx.Body)
	if err != nil {
		return
	}
	fmt.Printf("publish handler => [%s:'%s' // %s:'%s' // %s:'%s' // %s:'%s'] from %s:%d ...\n", TopicHeader, ctx.Headers[TopicHeader], ActionHeader, ctx.Headers[ActionHeader], QosHeader, ctx.Headers[QosHeader], ctx.Headers[BodyTitleHeader], string(buf), ctx.KadId.Host.String(), ctx.KadId.Port)

	if ctx.Headers[ActionHeader] == ActionPublish {
		// process the publish message, and put it to the channel group by topics.
		topicMessageChannel <- message{topic: ctx.Headers[TopicHeader], body: string(buf)}

		if ctx.Headers[QosHeader] == "1" {
			ctx.WriteHeader(ResponsePublishHeader, ResponseSuccess)
			ctx.WriteHeader(BodyTitleHeader, "process-time")
			ctx.Write([]byte(time.Since(now).String()))
		}
	} else {
		ctx.WriteHeader(ResponsePublishHeader, ResponseFailure)
		ctx.WriteHeader(BodyTitleHeader, "action-error")
		ctx.Write([]byte("not a publish operation request."))
	}
}

func subscribeHandler(ctx *sr.Context) {
	now := time.Now()
	buf, err := ioutil.ReadAll(ctx.Body)
	if err != nil {
		return
	}
	fmt.Printf("subscribe handler => [%s:'%s' // %s:'%s' // %s:'%s' // %s:'%s'] from %s:%d ...\n", TopicHeader, ctx.Headers[TopicHeader], ActionHeader, ctx.Headers[ActionHeader], QosHeader, ctx.Headers[QosHeader], ctx.Headers[BodyTitleHeader], string(buf), ctx.KadId.Host.String(), ctx.KadId.Port)

	if ctx.Headers[ActionHeader] == ActionSubscribe {
		// process the subscribe request, append to the channel group by topics, then push the data stream to it.
		topic := ctx.Headers[TopicHeader]
		key := fmt.Sprintf("%s_%s", topic, ctx.KadId.Pub.String())
		_, exist := topicSubscriber[key]
		if !exist {
			topicSubscriber[key] = func() {}
		}

		if ctx.Headers[QosHeader] == "1" {
			ctx.WriteHeader(ResponseSubscribeHeader, ResponseSuccess)
			ctx.WriteHeader(BodyTitleHeader, "process-time")
			ctx.Write([]byte(time.Since(now).String()))
		}
	} else {
		ctx.WriteHeader(ResponseSubscribeHeader, ResponseFailure)
		ctx.WriteHeader(BodyTitleHeader, "action-error")
		ctx.Write([]byte("not a subscribe operation request."))
	}
}

func main() {
	var listenAddr, probeAddr string
	var execPublishTask bool
	flag.BoolVar(&execPublishTask, "f", true, "publish messages, otherwise subscribe")
	flag.StringVar(&listenAddr, "l", ":9000", "address to listen for peers on")
	flag.StringVar(&probeAddr, "p", ":9000", "address to probe")
	flag.Parse()

	services := map[string]sr.Handler{ServicePublish: publishHandler, ServiceSubscribe: subscribeHandler}
	node := &fiesta.Node{PublicAddr: listenAddr, BindAddrs: []string{listenAddr}}
	defer node.Shutdown()

	check(node.StartWithKeyAndServiceAndProbeAddrs(sr.GenerateSecretKey(), services, probeAddr))

	exit := make(chan struct{})
	defer close(exit)

	go func() {
		for {
			select {
			case <-exit:
				return
			case msg, ok := <-topicMessageChannel:
				if ok {
					for k, fwd := range topicSubscriber {
						keys := strings.Split(k, "_")
						if keys[0] == msg.topic {
							fmt.Printf("...... Forward the message to the subscriber ......\n topic:%s,KadId:%s\n", keys[0], keys[1])
							fwd()
						}
					}
				}
			}
		}
	}()

	if execPublishTask {
		go func() {
			publishProviders := node.StreamNode.ProvidersFor(ServicePublish)

			for i := 0; ; i++ {
				select {
				case <-exit:
					return

				case <-time.After(5 * time.Second):
					publishProviders = node.StreamNode.ProvidersFor(ServicePublish)

					for _, provider := range publishProviders {
						topic := fmt.Sprintf("demo/stream/message/%d", i%2)
						body := ioutil.NopCloser(strings.NewReader(fmt.Sprintf("message_%d => public address : %s", i, node.PublicAddr)))
						publishHeader := map[string]string{TopicHeader: topic, ActionHeader: ActionPublish, QosHeader: "1", BodyTitleHeader: "PayLoad"}

						stream, err := provider.Push([]string{ServicePublish}, publishHeader, body)
						if err != nil {
							fmt.Printf("Unable to publish to %s: %s\n", provider.Addr(), err)
						}
						res, err := ioutil.ReadAll(stream.Reader)
						if !errors.Is(err, io.EOF) {
							check(err)
						}
						fmt.Printf("[%d] Got a response from %s! => {%s:%s, %s:'%s'}\n", i, provider.Addr(), ResponsePublishHeader, stream.Header.Headers[ResponsePublishHeader], stream.Header.Headers[BodyTitleHeader], string(res))
					}
				}
			}
		}()
	} else {
		go func() {
			subscribeProviders := node.StreamNode.ProvidersFor(ServiceSubscribe)

			for _, provider := range subscribeProviders {
				ns := time.Now().Nanosecond()
				topic := fmt.Sprintf("demo/stream/message/%d", ns%2)
				body := ioutil.NopCloser(strings.NewReader(fmt.Sprintf("no.%d, public address : %s", ns, node.PublicAddr)))
				subscribeHeader := map[string]string{TopicHeader: topic, ActionHeader: ActionSubscribe, QosHeader: "1", BodyTitleHeader: "PayLoad"}

				stream, err := provider.Push([]string{ServiceSubscribe}, subscribeHeader, body)
				if err != nil {
					fmt.Printf("Unable to subscribe to %s: %s\n", provider.Addr(), err)
				}
				res, err := ioutil.ReadAll(stream.Reader)
				if !errors.Is(err, io.ErrClosedPipe) {
					check(err)
				}
				fmt.Printf("Got a response from %s! => {%s:%s, %s:'%s'}\n", provider.Addr(), ResponseSubscribeHeader, stream.Header.Headers[ResponseSubscribeHeader], stream.Header.Headers[BodyTitleHeader], string(res))
			}

			for {
				select {
				case <-exit:
					return
				}
			}
		}()
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}
