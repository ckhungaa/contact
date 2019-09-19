package entity

import (
	"github.com/ckhungaa/common/domain/entities"
	"github.com/ckhungaa/proto/proto"
	"time"
)

type CallHistory struct {
	entities.Audit
	FromUserId string
	ToUserId   string
	proto.CallType
	time.Duration
	ReferenceId string
}
