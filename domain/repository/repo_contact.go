package repository

import (
	"context"
	"github.com/ckhungaa/contact/domain/entity"
)

func (r *RespositoryImpl) FindContactById(ctx context.Context, id string, result *entity.Contact) error{
	query, err := NewContactQuery(ctx, id)
	if err != nil {
		log.Errorf(ctx, "failed to parse id value: %v", err)
		return err
	}
	return r.FindById(ctx, query, result)
}