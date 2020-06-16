package jsonrpcserver

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"strconv"

	"github.com/898anil/gotest/configs"
	"github.com/898anil/gotest/pkg/common/logger"
)

func Start() {
	user := new(User)
	server := rpc.NewServer()
	server.Register(user)
	server.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(configs.Config.Rpc.Port))
	if err != nil {
		logger.Log.Fatal("listen error:", err)
	}

	for {
		if conn, err := listener.Accept(); err != nil {
			logger.Log.Fatal("accept error: " + err.Error())
		} else {
			logger.Log.Print("new connection established\n")
			go server.ServeCodec(jsonrpc.NewServerCodec(conn))
		}
	}
}
