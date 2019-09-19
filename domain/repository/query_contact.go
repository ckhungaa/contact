package repository

import (
	"context"
	"github.com/ckhungaa/common/domain/repositories"
)

type ContactQuery struct {
	*repositories.PartitionDBKey
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
	key, err := repositories.PartitionDBKeyFromId(ctx, id)
	if err != nil {
		log.Errorf(ctx, "failed to parse id into PartitionDBKey: %v", err)
		return nil, err
	}
	return &ContactQuery{PartitionDBKey: key}, nil
}
