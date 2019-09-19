package server

import (
	"context"
	"github.com/ckhungaa/proto/proto"
)

func (s *ContactServer) FindContactById(ctx context.Context, in *proto.FindContactByIdRequest) (*proto.ContactResponse, error) {
	contact, err := s.service.FindContactById(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	rsp , err := contact.ToProto()
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func (s *ContactServer) FindContactsByOwnerId(ctx context.Context, in *proto.FindContactsByOwnerIdRequest) (*proto.FindContactsByOwnerIdResponse, error) {
	return nil, nil
}
