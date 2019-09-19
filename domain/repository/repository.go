package repository

import (
	"context"
	"github.com/ckhungaa/common/domain/repositories"
	"github.com/ckhungaa/common/logs"
	"github.com/ckhungaa/contact/domain/entity"
)

var log  = logs.NewLogger("ContactRepository")

type Repository interface {
	FindContactById(ctx context.Context, id string, result *entity.Contact) error
}

type RespositoryImpl struct {
	*repositories.BaseRepository
}

func ProvideRepository(ctx context.Context, repo *repositories.BaseRepository) (Repository, error) {
	return &RespositoryImpl{BaseRepository: repo}, nil
}