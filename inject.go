//+build wireinject

package main

import (
	"context"
	"github.com/ckhungaa/common/component/configs"
	"github.com/ckhungaa/common/component/repos"
	"github.com/ckhungaa/contact/domain/repository"
	"github.com/ckhungaa/contact/domain/service"
	"github.com/ckhungaa/contact/server"
	"github.com/google/wire"
)

func injectSServer(ctx context.Context) (*server.ContactServer, error) {
	panic(wire.Build(
		configs.WireSet,
		repos.WireSet,
		repository.ProvideRepository,
		service.ProvideService,
		server.ProvideContact,
	))
}
