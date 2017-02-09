package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

var (
	network = flag.String("net", "tcp", "")
	address = flag.String("addr", ":8088", "")
)

type Server struct{}

// Touch just prints touched and returns nil.
func (*Server) Touch(args struct{}, reply *struct{}) error {
	fmt.Println("touched")
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
