package main

import (
	"RPCXdemo/onetoone/handler"
	"context"
	"flag"
	"github.com/smallnest/rpcx/client"
	"log"
)

var (
	addr1 = flag.String("addr1", "localhost:8972", "server1 address")
	addr2 = flag.String("addr2", "localhost:8973", "server2 address")
)

func Peer2Many()  {
	flag.Parse()
	d := client.NewMultipleServersDiscovery([]*client.KVPair{{Key:*addr1}, {Key:*addr2}})
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
	Peer2Many()
}