package main

import (
	"github.com/shibumi/kubernetes-avionics-playground/internal/navigation"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":7777")
	if err != nil {
		log.Println(err)
	}
	grpcServer := grpc.NewServer()

	server, err := navigation.NewServer("localhost", "5000")
	if err != nil {
		log.Println(err)
	}
	navigation.RegisterNavigationServer(grpcServer, server)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Println(err)
	}
}
