package repository

import (
	"context"
	"github.com/ckhungaa/common/component/repos"
	"github.com/ckhungaa/common/utils/logs"
	"github.com/ckhungaa/contact/domain/entity"
)

var log  = logs.NewLogger("repository")

type Repository interface {
	FindContactById(ctx context.Context, id string, result *entity.Contact) error
}

type RespositoryImpl struct {
	*repos.BaseRepository
}

func ProvideRepository(ctx context.Context, repo *repos.BaseRepository) (Repository, error) {
	return &RespositoryImpl{BaseRepository: repo}, nil
}