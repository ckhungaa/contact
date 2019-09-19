package server

import (
	"context"
	"github.com/ckhungaa/contact/domain/service"
)

type ContactServer struct {
	service service.Service
}

func ProvideContact(ctx context.Context, s service.Service) (*ContactServer, error) {
	return &ContactServer{service:s}, nil
}

