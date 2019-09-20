package repository

import (
	"context"
	"github.com/ckhungaa/contact/domain/entity"
)

func (r *RespositoryImpl) FindContactById(ctx context.Context, id string, result *entity.Contact) error{
	log.Infof(ctx, "FindContactById begin, id:%s", id)
	query, err := NewContactQuery(ctx, id)
	if err != nil {
		log.Errore(ctx, err, "failed to create query", )
		return err
	}
	log.Infof(ctx, "FindContactById end")
	return r.FindById(ctx, query, result)
}