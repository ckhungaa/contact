package main

import (
	"context"
	"github.com/ckhungaa/proto/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":8080"
)

func main() {

	ctx := context.TODO()
	log.Printf("run contact server")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	contactService, _ := injectSServer(ctx)
	proto.RegisterContactServerServer(s, contactService)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
