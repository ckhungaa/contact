package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/ckhungaa/common/utils/entities"
	"github.com/ckhungaa/contact/domain/entity"
	"github.com/ckhungaa/proto/proto"
	"github.com/guregu/dynamo"
	"log"
	"time"
)

var data = []*entity.Contact{
	&entity.Contact{
		Audit:       entities.Audit{
			Id:           "ownerId_userId",
			CreatedDate:  time.Now(),
			ModifiedDate: time.Now(),
		},
		OwnerId:     "ownerId",
		UserId:      "userId",
		UserName:    "user name",
		UserAlias:   "user alias",
		AvatarId:    "avatar id",
		ContactType: proto.ContactType_AdminContactType,
		PhoneNumber: "852_12345678",
	},
	&entity.Contact{
		Audit:       entities.Audit{
			Id:           "ownerId_userId2",
			CreatedDate:  time.Now(),
			ModifiedDate: time.Now(),
		},
		OwnerId:     "ownerId",
		UserId:      "userId2",
		UserName:    "user name2",
		UserAlias:   "user alias2",
		AvatarId:    "avatar id2",
		ContactType: proto.ContactType_FriendContactType,
		PhoneNumber: "852_22345678",
	},
}
func main() {
	endPoint := "http://localhost:4569"
	region := "ap-southeast-1"
	db := dynamo.New(session.New(), &aws.Config{Endpoint: &endPoint, Region: aws.String(region)}) //TODO: fix when go prod
	if err := db.CreateTable("Contact", entity.Contact{}).OnDemand(true).Run(); err != nil {
		log.Fatalf("failed to create table:%v\n", err)
	}

	for _, v := range data {
		if err := db.Table("Contact").Put(v).Run(); err != nil {
			log.Fatalf("failed to put contact:%v\n", err)
		}
	}


}