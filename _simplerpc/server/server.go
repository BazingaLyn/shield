package main

import (
	"log"
	"net"
	"net/rpc"
	"shield/_simplerpc/service"
)

func main() {
	rpc.RegisterName("HelloService", new(service.HelloService))

	listen, err := net.Listen("tcp", ":10086")

	if err != nil {
		log.Fatal("ListenTCP error", err)
	}

	conn, err := listen.Accept()

	if err != nil {
		log.Fatal("Accept error:", err)
	}
	rpc.ServeConn(conn)

}
