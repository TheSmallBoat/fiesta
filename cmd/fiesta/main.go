package main

import (
	"io/ioutil"
	"log"
	"net"
	"strconv"

	"github.com/BurntSushi/toml"
	"github.com/TheSmallBoat/carlo/rpc"
	gateway "github.com/TheSmallBoat/fiesta/http_gateway"
	"github.com/caddyserver/certmagic"
	"github.com/julienschmidt/httprouter"
	"github.com/spf13/pflag"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func hostOnly(hostPort string) string {
	host, _, err := net.SplitHostPort(hostPort)
	if err != nil {
		return hostPort
	}
	return host
}

func main() {
	var configPath string
	var bindHost net.IP
	var bindPort uint16

	pflag.StringVarP(&configPath, "config", "c", "config.toml", "path to config file")
	pflag.IPVarP(&bindHost, "host", "h", net.ParseIP("127.0.0.1"), "bind host")
	pflag.Uint16VarP(&bindPort, "port", "p", 9000, "bind port")
	pflag.Parse()

	var cfg gateway.Config

	buf, err := ioutil.ReadFile(configPath)
	if err == nil {
		check(toml.Unmarshal(buf, &cfg))
	} else {
		log.Printf("Unable to find a configuration file '%s'.", configPath)
	}
	check(cfg.Validate())

	if cfg.Addr != "" {
		host, port, err := net.SplitHostPort(cfg.Addr)
		check(err)

		bindHost = net.ParseIP(host)

		{
			port, err := strconv.ParseUint(port, 10, 16)
			check(err)

			bindPort = uint16(port)
		}
	}

	addr := rpc.HostAddr(bindHost, bindPort)
	node := &rpc.Node{PublicAddr: addr, SecretKey: rpc.GenerateSecretKey()}
	check(node)
	defer node.Shutdown()

}
