package main

import (
	"RPCXdemo/onetoone/handler"
	"flag"
	"github.com/smallnest/rpcx/server"
)

var (
	addr1 = flag.String("addr1", "localhost:8972", "server1 address")
	addr2 = flag.String("addr2", "localhost:8973", "server2 address")
)

func createServer(addr string)  {
	s := server.NewServer()
	s.RegisterName("Arith", new(handler.Arith), "")
	s.Serve("tcp", addr)
}

func main()  {
	flag.Parse()
	go createServer(*addr1)
	go createServer(*addr2)
	select {

	}
}