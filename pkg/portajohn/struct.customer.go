package portajohn

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Customer struct {
	UID                     string             `json:"userRandomId" bson:"userRandomId"`
	Email                   string             `json:"email" bson:"email"`
	PhoneNumber             string             `json:"phoneNumber" bson:"phoneNumber"`
	Name                    string             `json:"name" bson:"name"`
	ProfileImage            string             `json:"profileImage" bson:"profileImage"`
	VerificationCode        string             `json:"verificationCode" bson:"verificationCode"`
	VerificationCodeExpTime int                `json:"verificationCodeExpiryTime" bson:"verificationCodeExpiryTime"`
	Timezone                string             `json:"timezone" bson:"timezone"`
	Zip                     string             `json:"zipCode" bson:"zipCode"`
	AccountType             int                `json:"accountType" bson:"accountType"`
	Status                  int                `json:"status" bson:"status"`
	Permissions             []interface{}      `json:"permissions" bson:"permissions"`
	RefreshAccessToken      string             `json:"refreshAccessToken" bson:"refreshAccessToken"`
	Deleted                 bool               `json:"isDeleted" bson:"isDeleted"`
	Blocked                 bool               `json:"isBlocked" bson:"isBlocked"`
	BanEndTime              int                `json:"banEndTime" bson:"banEndTime"`
	Location                Location           `json:"location" bson:"location"`
	LicenseNumber           string             `json:"licenseNumber" bson:"licenseNumber"`
	SecondaryEmail          string             `json:"secondaryEmail" bson:"secondaryEmail"`
	TaxExempt               bool               `json:"isTaxExempted" bson:"isTaxExempted"`
	BusyInCall              bool               `json:"isBusyInCall" bson:"isBusyInCall"`
	BillingInfo             BillingInfo        `json:"billingInfo" bson:"billingInfo"`
	AccountPayable          AccountPayable     `json:"accountPayable" bson:"accountPayable"`
	CreatedBy               primitive.ObjectID `json:"createdBy" bson:"createdBy"`
	CreatedAt               time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt               time.Time          `json:"updatedAt" bson:"updatedAt"`
}

type Location struct {
	Type   string `json:"type" bson:"type"`
	Coords []int  `json:"coordinates" bson:"coordinates"`
}

type BillingInfo struct {
	Address  string   `json:"address" bson:"adress"`
	City     string   `json:"city" bson:"city"`
	Zip      string   `json:"zipCode" bson:"zipCode"`
	Location Location `json:"location" bson:"location"`
}

type AccountPayable struct {
	Name        string `json:"name" bson:"name"`
	Email       string `json:"email" bson:"email"`
	PhoneNumber string `json:"phoneNumber" bson:"phoneNumber"`
}
