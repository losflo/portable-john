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
	Deleted           bool               `json:"isDeleted" bson:"isDeleted"`
	Status            int                `json:"status" bson:"status"`
	AccountID         primitive.ObjectID `json:"userId" bson:"userId"`
	CreatedBy         primitive.ObjectID `json:"createdBy" bson:"createdBy"`
	CreatedAt         time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt         time.Time          `json:"updatedAt" bson:"updatedAt"`

	// added fields
	Address2    string `json:"address2" bson:"address2"`
	State       string `json:"state" bson:"state"`
	Zip         string `json:"zipCode" bson:"zipCode"`
	TacMaster   string `json:"tacMaster" bson:"tacMaster"`
	TacMasterId int64  `json:"tacMasterId" bson:"tacMasterId"`
}
