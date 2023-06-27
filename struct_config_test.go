package main

import (
	"fmt"
	"testing"
)

type OptsFunc func(*Opts)

type Opts struct {
	MaxConn int
	ID      string
	Tls     bool
}

func defaultOpts() Opts {
	return Opts{
		MaxConn: 10,
		ID:      "default",
		Tls:     false,
	}
}

func WitTLS(opts *Opts) {
	opts.Tls = true
}

func WithMaxConn(n int) OptsFunc {
	return func(o *Opts) {
		o.MaxConn = n
	}
}

type Server struct {
	Opts Opts
}

func NewServer(opts ...OptsFunc) *Server {
	defaultOpts := defaultOpts()

	for _, fn := range opts {
		fn(&defaultOpts)
	}

	return &Server{
		Opts: defaultOpts,
	}
}

func TestConfigPattern(t *testing.T) {
	server := NewServer(WitTLS)

	fmt.Printf("%+v\n", server)
}
