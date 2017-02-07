package main

import (
	"flag"
	"fmt"
	"log"
	"net/rpc"

	"github.com/martindrlik/rpctest"
)

var (
	network = flag.String("network", "tcp", `net must be a stream-oriented network: "tcp", "tcp4", "tcp6", "unix" or "unixpacket"`)
	address = flag.String("address", ":8088", `listens on the local network address, for TCP and UDP, the syntax of laddr is "host:port", like "127.0.0.1:8080". If host is omitted, as in ":8080"`)
)

func main() {
	flag.Parse()
	cl, err := rpc.DialHTTP(*network, *address)
	if err != nil {
		log.Fatal(err)
	}
	args := &rpctest.Args{7, 8}
	reply := rpctest.Reply{}
	err = cl.Call("Server.Reload", args, &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Server: %v*%v=%v\n", args.A, args.B, reply.C)
}
