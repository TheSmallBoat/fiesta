package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	sr "github.com/TheSmallBoat/carlo/streaming_rpc"
	"github.com/TheSmallBoat/fiesta"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func chat(ctx *sr.Context) {
	buf, err := ioutil.ReadAll(ctx.Body)
	if err != nil {
		return
	}
	fmt.Printf("Got '%s' from %s:%d!\n", string(buf), ctx.KadId.Host.String(), ctx.KadId.Port)
}

func main() {
	var listenAddr string
	flag.StringVar(&listenAddr, "l", ":9000", "address to listen for peers on")
	flag.Parse()

	services := map[string]sr.Handler{"chat": chat}
	node := &fiesta.Node{PublicAddr: listenAddr, BindAddrs: []string{listenAddr}}
	defer node.Shutdown()

	check(node.StartWithKeyAndServiceAndProbeAddrs(sr.GenerateSecretKey(), services, flag.Args()...))

	br := bufio.NewReader(os.Stdin)
	for {
		line, _, err := br.ReadLine()
		if err != nil {
			break
		}

		line = bytes.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		providers := node.StreamNode.ProvidersFor("chat")
		for _, provider := range providers {
			_, err := provider.Push([]string{"chat"}, nil, ioutil.NopCloser(bytes.NewReader(line)))
			if err != nil {
				fmt.Printf("Unable to broadcast to %s: %s\n", provider.Addr(), err)
			}
		}
	}
}
