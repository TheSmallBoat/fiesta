package fiesta

import (
	"net"
)

type BindFunc func() (net.Listener, error)

func BindTCPAnyPort() BindFunc {
	return func() (net.Listener, error) { return net.Listen("tcp", ":0") }
}

func BindTCP(addr string) BindFunc {
	return func() (net.Listener, error) { return net.Listen("tcp", addr) }
}

func BindTCPv4(addr string) BindFunc {
	return func() (net.Listener, error) { return net.Listen("tcp4", addr) }
}

func BindTCPv6(addr string) BindFunc {
	return func() (net.Listener, error) { return net.Listen("tcp6", addr) }
}
