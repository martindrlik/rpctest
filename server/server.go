package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/martindrlik/rpctest"
)

var (
	network = flag.String("network", "tcp", `net must be a stream-oriented network: "tcp", "tcp4", "tcp6", "unix" or "unixpacket"`)
	address = flag.String("address", ":8088", `listens on the local network address, for TCP and UDP, the syntax of laddr is "host:port", like "127.0.0.1:8080". If host is omitted, as in ":8080"`)
)

type Server int

// Reload replies by multiplying A and B of args.
func (*Server) Reload(args *rpctest.Args, reply *rpctest.Reply) error {
	fmt.Printf("Reload called with arguments %v\n", args)
	reply.C = args.A * args.B
	return nil
}

func main() {
	flag.Parse()
	fmt.Printf("Listen(%v, %v)\n", *network, *address)
	srv := new(Server)
	rpc.Register(srv)
	rpc.HandleHTTP()
	l, err := net.Listen(*network, *address)
	if err != nil {
		log.Fatal(err)
	}
	http.Serve(l, nil)
}
