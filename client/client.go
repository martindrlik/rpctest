package main

import (
	"flag"
	"fmt"
	"log"
	"net/rpc"
)

var (
	network = flag.String("net", "tcp", "")
	address = flag.String("addr", ":8088", "")
)

func main() {
	flag.Parse()
	client, err := rpc.DialHTTP(*network, *address)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Call("Server.Touch", nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("touch")
}
