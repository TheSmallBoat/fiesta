package fiesta

import (
	"fmt"
	"log"
	"math"
	"net"
	"sync"

	sr "github.com/TheSmallBoat/carlo/streaming_rpc"
	"github.com/lithdew/kademlia"
)

type Node struct {
	// A reachable, public address which peers may reach you on.
	// The format of the address must be [host]:[port].
	PublicAddr string

	// A list of IPv4/IPv6 addresses and ports assembled as [host]:[port] which
	// your fiesta node will listen for other nodes from.
	BindAddrs []string

	lns []net.Listener

	StreamNode *sr.StreamNode

	wg sync.WaitGroup
}

func (n *Node) StartWithKeyAndServiceAndProbeAddrs(sk kademlia.PrivateKey, services map[string]sr.Handler, probeAddrs ...string) error {
	var (
		publicHost net.IP
		publicPort uint16

		kid *kademlia.ID
		tab *kademlia.Table
	)

	if sk != kademlia.ZeroPrivateKey {
		if n.PublicAddr != "" { // resolve the address
			addr, err := net.ResolveTCPAddr("tcp", n.PublicAddr)
			if err != nil {
				return err
			}

			publicHost = addr.IP

			if addr.Port <= 0 || addr.Port >= math.MaxUint16 {
				return fmt.Errorf("'%d' is an invalid port", addr.Port)
			}

			publicPort = uint16(addr.Port)
		} else { // get a random public address
			ln, err := net.Listen("tcp", ":0")
			if err != nil {
				return fmt.Errorf("unable to listen on any port: %w", err)
			}
			bindAddr := ln.Addr().(*net.TCPAddr)
			publicHost = bindAddr.IP
			publicPort = uint16(bindAddr.Port)
			if err := ln.Close(); err != nil {
				return fmt.Errorf("failed to close listener for getting avaialble port: %w", err)
			}
		}

		if publicHost == nil {
			publicHost = net.ParseIP("0.0.0.0")
		}

		kid = &kademlia.ID{
			Pub:  sk.Public(),
			Host: publicHost,
			Port: publicPort,
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

	if kid != nil && n.StreamNode.NetProtocol == sr.NetProtocolTCP && len(n.BindAddrs) == 0 {
		ln, err := net.Listen("tcp", sr.HostAddr(kid.Host, kid.Port))
		if err != nil {
			return err
		}

		log.Printf("Listening for Fiesta nodes on '%s'.", ln.Addr().String())

		n.wg.Add(1)
		go func() {
			defer n.wg.Done()
			_ = n.StreamNode.Srv.Serve(ln)
		}()

		n.lns = append(n.lns, ln)
	}

	for _, addr := range n.BindAddrs {
		ln, err := net.Listen("tcp", addr)
		if err != nil {
			for _, ln := range n.lns {
				_ = ln.Close()
			}
			return err
		}

		log.Printf("Listening for Fiesta nodes on '%s'.", ln.Addr().String())

		n.wg.Add(1)
		go func() {
			defer n.wg.Done()
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
	n.wg.Wait()
}
