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

func clock(ctx *sr.Context) {
	latest := time.Now()
	ours := latest.Format(time.Stamp)

	timestamp, err := ioutil.ReadAll(ctx.Body)
	if err != nil {
		return
	}

	fmt.Printf("Got (%s:%d)'s time ('%s')! Sent back ours ('%s').\n", ctx.KadId.Host.String(), ctx.KadId.Port, string(timestamp), ours)

	ctx.Write([]byte(ours))
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

	services := map[string]sr.Handler{"clock": clock, "chat": chat}
	node := &fiesta.Node{PublicAddr: listenAddr, BindAddrs: []string{listenAddr}}
	defer node.Shutdown()

	check(node.StartWithKeyAndServiceAndProbeAddrs(sr.GenerateSecretKey(), services, flag.Args()...))

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}
