package fiesta

import (
	"fmt"
	"log"
	"math"
	"net"

	sr "github.com/TheSmallBoat/carlo/streaming_rpc"
	"github.com/lithdew/kademlia"
)

type Node struct {
	PublicAddr string
	BindAddrs  []BindFunc
	lns        []net.Listener

	StreamNode *sr.StreamNode
}

func (n *Node) Start(sk kademlia.PrivateKey, services map[string]sr.Handler, probeAddrs ...string) error {
	var (
		bindHost net.IP
		bindPort uint16

		kid *kademlia.ID
		tab *kademlia.Table
	)

	if sk != kademlia.ZeroPrivateKey {
		if n.PublicAddr != "" { // resolve the address
			addr, err := net.ResolveTCPAddr("tcp", n.PublicAddr)
			if err != nil {
				return err
			}

			bindHost = addr.IP
			if bindHost == nil {
				return fmt.Errorf("'%s' is an invalid host: it must be an IPv4 or IPv6 address", addr.IP)
			}

			if addr.Port <= 0 || addr.Port >= math.MaxUint16 {
				return fmt.Errorf("'%d' is an invalid port", addr.Port)
			}

			bindPort = uint16(addr.Port)
		} else { // get a random public address
			ln, err := net.Listen("tcp", ":0")
			if err != nil {
				return fmt.Errorf("unable to listen on any port: %w", err)
			}
			bindAddr := ln.Addr().(*net.TCPAddr)
			bindHost = bindAddr.IP
			bindPort = uint16(bindAddr.Port)
			if err := ln.Close(); err != nil {
				return fmt.Errorf("failed to close listener for getting avaialble port: %w", err)
			}
		}

		kid = &kademlia.ID{
			Pub:  sk.Public(),
			Host: bindHost,
			Port: bindPort,
		}

		tab = kademlia.NewTable(kid.Pub)
	} else {
		kid = nil
		tab = kademlia.NewTable(kademlia.ZeroPublicKey)
	}

	n.StreamNode = sr.NewStreamNode(sk, kid, tab)
	if services != nil {
		n.StreamNode.Services = services
	}

	err := n.StreamNode.Start()
	if err != nil {
		return err
	}

	if n.StreamNode.KadId != nil && n.StreamNode.NetProtocol == sr.NetProtocolTCP && len(n.BindAddrs) == 0 {
		ln, err := BindTCP(sr.HostAddr(n.StreamNode.KadId.Host, n.StreamNode.KadId.Port))()
		if err != nil {
			return err
		}

		log.Printf("Listening for Fiesta nodes on '%s'.", ln.Addr().String())

		n.StreamNode.Wg.Add(1)
		go func() {
			defer n.StreamNode.Wg.Done()
			_ = n.StreamNode.Srv.Serve(ln)
		}()

		n.lns = append(n.lns, ln)
	}

	for _, fn := range n.BindAddrs {
		ln, err := fn()
		if err != nil {
			for _, ln := range n.lns {
				_ = ln.Close()
			}
			return err
		}

		log.Printf("Listening for Fiesta nodes on '%s'.", ln.Addr().String())

		n.StreamNode.Wg.Add(1)
		go func() {
			defer n.StreamNode.Wg.Done()
			_ = n.StreamNode.Srv.Serve(ln)
		}()

		n.lns = append(n.lns, ln)
	}

	for _, addr := range probeAddrs {
		err := n.StreamNode.ProbeWithAddr(addr)
		if err != nil {
			return fmt.Errorf("failed to probe '%s': %w", addr, err)
		}
	}

	n.StreamNode.Bootstrap()

	return nil
}

func (n *Node) Shutdown() {
	n.StreamNode.Shutdown()

	for _, ln := range n.lns {
		_ = ln.Close()
	}
}
