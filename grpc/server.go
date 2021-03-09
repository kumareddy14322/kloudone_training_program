package main

import (
	"log"
	"net"

	"github.com/golang/protobuf/protoc-gen-go/grpc"
)

func main() {
	lis,err:=net.Listen("tcp",":9000")
	if err!=nil{
		log.Fatalf("failed to listen on port9000: %v",err)
	}

	grpcServer := grpc.NewServer()

	if err :=grpcServer.Serve(lis);
	err !=nil{
		log.Fatalf("failed to serer Grpc server over portt 9000: %v",err)
	}
}