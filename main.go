package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc_serivce/service"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	fmt.Println("Serve...")
	service.RegisterServiceServer(srv, &service.Service{})
	if err = srv.Serve(listener); err != nil {
		log.Fatalf("failed to servce: %v", err)
	}
}
