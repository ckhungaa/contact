package main

import (
	"github.com/ckhungaa/common/utils/contexts"
	"github.com/ckhungaa/common/utils/logs"
	"github.com/ckhungaa/proto/proto"
	"google.golang.org/grpc"
	"net"
)

const (
	port = ":8888"
)

var (
	log = logs.NewLogger("main")
)

func main() {
	ctx := contexts.NewContext("main")
	log.Infof(ctx, "main begin 123")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatale(ctx, err, "failed to listen", )
	}
	s := grpc.NewServer()
	contactService, _ := injectSServer(ctx)
	proto.RegisterContactServerServer(s, contactService)
	if err := s.Serve(lis); err != nil {
		log.Fatale(ctx, err, "failed to serve")
	}
}
