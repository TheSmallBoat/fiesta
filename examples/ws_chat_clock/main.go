package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"time"

	sr "github.com/TheSmallBoat/carlo/streaming_rpc"
	"github.com/TheSmallBoat/fiesta"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var listenAddr string
	flag.StringVar(&listenAddr, "l", ":9000", "address to listen for peers on")
	flag.Parse()

	node := &fiesta.Node{PublicAddr: listenAddr, BindAddrs: []string{listenAddr}}
	defer node.Shutdown()

	services := map[string]sr.Handler{
		"clock": func(ctx *sr.Context) {
			latest := time.Now()
			ours := latest.Format(time.Stamp)

			timestamp, err := ioutil.ReadAll(ctx.Body)
			if err != nil {
				return
			}
			kid := node.StreamNode.KadId
			fmt.Printf("Clock Service => Got (%s:%d)'s time ('%s')! Sent back ours ('%s').\n", ctx.KadId.Host.String(), ctx.KadId.Port, timestamp, ours)
			ctx.Write([]byte(fmt.Sprintf("Clock '%s' From %s(%s:%d)", ours, kid.Pub.String(), kid.Host.String(), kid.Port)))
		},
		"chat": func(ctx *sr.Context) {
			buf, err := ioutil.ReadAll(ctx.Body)
			if err != nil {
				return
			}
			kid := node.StreamNode.KadId
			fmt.Printf("Chat Service => Got '%s' from %s:%d!\n", buf, ctx.KadId.Host.String(), ctx.KadId.Port)
			ctx.Write([]byte(fmt.Sprintf("Echo %s From %s(%s:%d)", buf, kid.Pub.String(), kid.Host.String(), kid.Port)))
		}}
	check(node.StartWithKeyAndServiceAndProbeAddrs(sr.GenerateSecretKey(), services, flag.Args()...))

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}
