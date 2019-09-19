//+build wireinject

package main

import (
	"context"
	"github.com/ckhungaa/common/config"
	"github.com/ckhungaa/common/domain/repositories"
	"github.com/ckhungaa/contact/domain/repository"
	"github.com/ckhungaa/contact/domain/service"
	"github.com/ckhungaa/contact/server"
	"github.com/google/wire"
)

func injectSServer(ctx context.Context) (*server.ContactServer, error) {
	panic(wire.Build(
		config.WireSet,
		repositories.WireSet,
		repository.ProvideRepository,
		service.ProvideService,
		server.ProvideContact,
	))
}
