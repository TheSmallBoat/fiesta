package http_gateway

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var HttpMethods = map[string]struct{}{
	http.MethodGet:    {},
	http.MethodPost:   {},
	http.MethodPut:    {},
	http.MethodDelete: {},
	http.MethodPatch:  {},
}

const DefaultShutdownTimeout = 10 * time.Second

type Duration struct {
	time.Duration
}

func (d *Duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}

type Config struct {
	Addr     string
	ConfHttp []ConfigHttp `toml:"http"`
}

func (c Config) Validate() error {
	for _, srv := range c.ConfHttp {
		err := srv.Validate()
		if err != nil {
			return err
		}
	}

	return nil
}

type ConfigHttp struct {
	Domain  string
	Domains []string

	Addr  string
	Addrs []string

	EnableHttps bool

	RedirectTrailingSlash *bool `toml:"redirect_trailing_slash"`
	RedirectFixedPath     *bool `toml:"redirect_fixed_path"`

	Timeout struct {
		Read       Duration
		ReadHeader Duration
		Idle       Duration
		Write      Duration
		Shutdown   Duration
	}

	Min struct {
		BodySize *int `toml:"body_size"`
	}

	Max struct {
		HeaderSize int  `toml:"header_size"`
		BodySize   *int `toml:"body_size"`
	}

	ConfRoutes []ConfigRoute `toml:"routes"`
}

func (h ConfigHttp) GetShutdownTimeout() time.Duration {
	if h.Timeout.Shutdown.Duration < 0 {
		return DefaultShutdownTimeout
	}
	return h.Timeout.Shutdown.Duration
}

func (h ConfigHttp) GetDomains() []string {
	if h.Domain != "" {
		return []string{h.Domain}
	}
	return h.Domains
}

func (h ConfigHttp) GetAddrs() []string {
	if len(h.Addrs) > 0 {
		return h.Addrs
	}
	if h.Addr != "" {
		return []string{h.Addr}
	}
	if h.EnableHttps {
		return []string{net.JoinHostPort("", "443")}
	}
	return []string{net.JoinHostPort("", "80")}
}

func (h ConfigHttp) Validate() error {
	if h.Domain != "" && h.Domains != nil {
		return errors.New("'domain' and 'domains' cannot both be non-nil at the same time")
	}

	if h.Addr != "" && h.Addrs != nil {
		return errors.New("'addr' and 'addrs' cannot both be non-nil at the same time")
	}

	for _, route := range h.ConfRoutes {
		err := route.Validate()
		if err != nil {
			return err
		}
	}

	return nil
}

type ConfigRoute struct {
	Path     string
	Dispatch string
	Static   string // static files
	Service  string
	Services []string

	NoCache   bool
	WebSocket bool

	Min struct {
		BodySize *int `toml:"body_size"`
	}

	Max struct {
		BodySize *int `toml:"body_size"`
	}
}

func (r ConfigRoute) GetServices() []string {
	if r.Service == "" {
		return r.Services
	}
	return []string{r.Service}
}

func (r ConfigRoute) Validate() error {
	if r.Service != "" && r.Services != nil {
		return errors.New("'service' and 'services' cannot both be non-nil at the same time")
	}

	fields := strings.Fields(r.Path)
	if len(fields) != 2 {
		return fmt.Errorf("invalid number of fields in route path '%s' (format: 'HTTP_METHOD /path/here')",
			r.Path)
	}

	method := strings.ToUpper(fields[0])
	_, exists := HttpMethods[method]
	if !exists {
		return fmt.Errorf("unknown http method '%s'", method)
	}

	if len(fields[1]) < 1 || fields[1][0] != '/' {
		return fmt.Errorf("path must begin with '/' in path '%s'", fields[1])
	}

	_, err := url.ParseRequestURI(fields[1])
	if err != nil {
		return fmt.Errorf("invalid http path '%s': %w", fields[1], err)
	}

	if r.Static != "" && method != http.MethodGet {
		return fmt.Errorf("path '%s' method must be 'GET' to serve files", r.Path)
	}

	return nil
}
