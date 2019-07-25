.PHONY: all dep server client

ROOT_DIR = $(CURDIR)
BIN_DIR = $(ROOT_DIR)/bin

all: server client

dep:
	go get -u -v -tags "reuseport quic kcp zookeeper etcd consul ping rudp utp" github.com/smallnest/rpcx/...

server:
	mkdir -p $(BIN_DIR)
	cd $(BIN_DIR) && go build -tags consul mysamples/rpcxsample/server

client:
	mkdir -p $(BIN_DIR)
	cd $(BIN_DIR) && go build -tags consul mysamples/rpcxsample/client