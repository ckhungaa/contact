package server

import (
	"context"
	"github.com/ckhungaa/common/utils/logs"
	"github.com/ckhungaa/contact/domain/service"
)

var (
	log = logs.NewLogger("server")
)

type ContactServer struct {
	service service.Service
}

func ProvideContact(ctx context.Context, s service.Service) (*ContactServer, error) {
	return &ContactServer{service:s}, nil
}

