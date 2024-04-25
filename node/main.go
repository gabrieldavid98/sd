package main

import (
	"log"
	"net"
	"node/server"

	pb "idl/gen/node/v1"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":7001")

	if err != nil {
		log.Fatalf("failed connection: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterLeibnizPiServiceServer(s, server.New())

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
