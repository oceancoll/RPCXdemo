package main

import (
	"RPCXdemo/onetoone/handler"
	"flag"
	"github.com/smallnest/rpcx/server"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main()  {
	flag.Parse()
	s := server.Server{}
	s.RegisterName("Arith", new(handler.Arith), "")
	go s.Serve("tcp", *addr)
	select {

	}
}