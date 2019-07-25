package main

import (
	"RPCXdemo/onetoone/handler"
	"context"
	"flag"
	"github.com/smallnest/rpcx/client"
	"log"
)

var (
	addr = flag.String("addr", "127.0.0.1:8972", "server address")
)

func Peer2Peer()  {
	flag.Parse()
	d := client.NewPeer2PeerDiscovery("tcp@" + *addr, "")
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()
	args := &handler.Args{
		A: 10,
		B: 20,
	}
	reply := &handler.Reply{}
	err := xclient.Call(context.Background(), "Mul", args, reply)
	if err != nil{
		log.Fatalf("failed to call: %v", err)
	}
	log.Printf("%d * %d = %d", args.A, args.B, reply.C)
}

func main()  {
	Peer2Peer()
}