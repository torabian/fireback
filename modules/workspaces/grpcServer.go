package workspaces

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func CreateGrpcServer(port string, server func(*grpc.Server), options ...grpc.ServerOption) {
	config := GetAppConfig()
	p := config.PublicServer.GrpcPort

	if port != "" {
		p = port
	}

	listener, err := net.Listen("tcp", ":"+p)

	fmt.Println("Grpc server is started at", p)

	if err != nil {
		panic(err)
	}

	// s := grpc.NewServer(options...)
	s := grpc.NewServer()

	server(s)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
