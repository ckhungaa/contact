package repository

import (
	"context"
	"github.com/ckhungaa/common/component/repos"
)

type ContactQuery struct {
	*repos.PartitionDBKey
}

func (q *ContactQuery) table() string {
	return "Contact"
}

func (q *ContactQuery) PartitionKeyName() string {
	return "OwnerId"
}

func (q *ContactQuery) SortKeyName() string {
	return "UserId"
}

func (q *ContactQuery) TableName() string {
	return "Contact"
}

func NewContactQuery(ctx context.Context, id string) (*ContactQuery, error) {
	key, err := repos.PartitionDBKeyFromId(ctx, id)
	if err != nil {
		log.Errore(ctx, err, "failed to parse id(%s) into PartitionDBKey", id)
		return nil, err
	}
	return &ContactQuery{PartitionDBKey: key}, nil
}
