package portajohn

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Site struct {
	ID                primitive.ObjectID `json:"_id" bson:"_id"`
	Title             string             `json:"title" bson:"title"`
	ContactName       string             `json:"contactName" bson:"contactName"`
	Email             string             `json:"email" bson:"email"`
	PhoneNumber       string             `json:"phoneNumber" bson:"phoneNumber"`
	Address           string             `json:"address" bson:"address"`
	City              string             `json:"city" bson:"city"`
	Location          Location           `json:"location"`
	SameAsBillingInfo bool               `json:"isSameAsBillingInfo"`
	Deleted           bool               `json:"isDeleted"`
	Status            int                `json:"status" bson:"status"`
	UID               string             `json:"userId" bson:"userId"`
	CreatedBy         primitive.ObjectID `json:"createdBy" bson:"createdBy"`
	CreatedAt         time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt         time.Time          `json:"updatedAt" bson:"updatedAt"`
}
