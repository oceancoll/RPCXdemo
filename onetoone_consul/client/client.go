package main

import (
	"github.com/smallnest/rpcx/client"
	"log"
	"context"
	"flag"
	"RPCXdemo/onetoone_consul/handler"
)

var (
	consulAddr = flag.String("consulAddr", "localhost:8500", "consul address")
	basePath   = flag.String("base", "/rpcx_test/Arith", "prefix path")
)

func ConsulServer() {
	flag.Parse()

	d := client.NewConsulDiscovery(*basePath, "", []string{*consulAddr}, nil)
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	args := &handler.Args{
		A: 10,
		B: 20,
	}

	reply := &handler.Reply{}
	err := xclient.Call(context.Background(), "Mul", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Printf("%d * %d = %d", args.A, args.B, reply.C)

}
func main() {
	ConsulServer()
}