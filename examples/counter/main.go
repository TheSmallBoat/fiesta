package main

import (
	"os"
	"os/signal"
	"strconv"
	"sync/atomic"

	sr "github.com/TheSmallBoat/carlo/streaming_rpc"
	"github.com/TheSmallBoat/fiesta"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	counter := uint64(0)

	services := map[string]sr.Handler{
		"count": func(ctx *sr.Context) {
			current := atomic.AddUint64(&counter, 1) - 1
			ctx.Write(strconv.AppendUint(nil, current, 10))
		},
	}

	node := &fiesta.Node{}
	check(node.StartWithKeyAndServiceAndProbeAddrs(sr.GenerateSecretKey(), services, "127.0.0.1:9000"))

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	node.Shutdown()
}
