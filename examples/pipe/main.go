package main

import (
	"io"
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

func pipe(ctx *sr.Context) {
	_, _ = io.Copy(ctx, ctx.Body)
}

func main() {
	node := &fiesta.Node{}
	services := map[string]sr.Handler{"pipe": pipe}
	check(node.StartWithKeyAndServiceAndProbeAddrs(sr.GenerateSecretKey(), services, "127.0.0.1:9000"))

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	node.Shutdown()
}
