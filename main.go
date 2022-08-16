package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	test "test/proto"
)

type AllowListServer struct {
	responseContent test.AllowListProtoResponse
	test.UnimplementedAuditInterfaceServer
}

func main() {
	lis, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		log.Fatalln("cannot create a listener at the address")
	}
	fmt.Println("服务启动，端口5000")
	grpcServer := grpc.NewServer()
	//test.RegisterAllowListInterfaceServer(grpcServer, AllowListServer)
	log.Fatalln(grpcServer.Serve(lis))
}
