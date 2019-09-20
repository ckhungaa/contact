package service

import (
	"context"
	"github.com/ckhungaa/common/utils/logs"
	"github.com/ckhungaa/contact/domain/entity"
	"github.com/ckhungaa/contact/domain/repository"
)

var log  = logs.NewLogger("service")

// Service service
type Service interface {
	FindContactById(ctx context.Context, id string) (*entity.Contact, error)
}

type ServiceImpl struct {
	repo repository.Repository
}

func ProvideService(ctx context.Context, repo repository.Repository) (Service, error) {
	return &ServiceImpl{repo:repo}, nil
}

func (s *ServiceImpl) FindContactById(ctx context.Context, id string) (*entity.Contact, error) {
	log.Infof(ctx, "FindContactById begin, id:%s", id)
	var contact entity.Contact
	if err := s.repo.FindContactById(ctx, id, &contact); err != nil {
		log.Errore(ctx, err, "failed to find contact by id")
		return nil, err
	}
	log.Infof(ctx, "FindContactById end")
	return &contact, nil
}
