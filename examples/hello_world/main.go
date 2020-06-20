package main

import (
	"os"
	"os/signal"

	sr "github.com/TheSmallBoat/carlo/streaming_rpc"
	"github.com/TheSmallBoat/fiesta"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func helloWorld(ctx *sr.Context) {
	ctx.WriteHeader("Content-Type", "text/plain; charset=utf-8")
	ctx.Write([]byte("Hello world!"))
}

func main() {
	node := &fiesta.Node{}
	services := map[string]sr.Handler{"hello_world": helloWorld}
	check(node.StartWithKeyAndServiceAndProbeAddrs(sr.GenerateSecretKey(), services, "127.0.0.1:9000"))

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	node.Shutdown()
}
