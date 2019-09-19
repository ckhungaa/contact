package service

import (
	"context"
	"github.com/ckhungaa/common/logs"
	"github.com/ckhungaa/contact/domain/entity"
	"github.com/ckhungaa/contact/domain/repository"
)

var log  = logs.NewLogger("ContactService")
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
	var contact entity.Contact
	if err := s.repo.FindContactById(ctx, id, &contact); err != nil {
		log.Errorf(ctx, "failed to find contact by id:%v", err)
		return nil, err
	}
	return &contact, nil
}
