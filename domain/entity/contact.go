package entity

import (
	"github.com/ckhungaa/common/utils/entities"
	"github.com/ckhungaa/proto/proto"
)

type Contact struct {
	Audit entities.Audit
	OwnerId     string            `dynamo:"OwnerId,hash"` // partition key
	UserId      string            `dynamo:"UserId,range"` // sort key
	UserName    string
	UserAlias 	string
	AvatarId    string
	ContactType proto.ContactType
	PhoneNumber string
}

func (m *Contact) ToProto() (*proto.ContactResponse, error){

	audit, err := m.Audit.ToProto()
	if err != nil {
		return nil, err
	}
	return &proto.ContactResponse{
		Audit:               audit,
		OwnerId:              m.OwnerId,
		UserId:               m.UserId,
		UserName:             m.UserName,
		UserAlias:            m.UserId,
		AvatarId:             m.UserAlias,
		ContactType:          m.ContactType,
		PhoneNumber:          m.PhoneNumber,
	}, nil
}