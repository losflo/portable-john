package portajohn

import (
	"fmt"
	"math/rand"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Customer struct {
	ID                      primitive.ObjectID `json:"_id" bson:"_id"`
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
	InvalidLoginAttempTimes []interface{}      `json:"invalidLoginAttempTimes" bson:"invalidLoginAttempTimes"`
	BanEndTime              int                `json:"banEndTime" bson:"banEndTime"`
	Location                Location           `json:"location,omitempty" bson:"location,omitempty"`
	LicenseNumber           string             `json:"licenseNumber" bson:"licenseNumber"`
	SecondaryEmail          string             `json:"secondaryEmail" bson:"secondaryEmail"`
	TaxExempt               bool               `json:"isTaxExempted" bson:"isTaxExempted"`
	BusyInCall              bool               `json:"isBusyInCall" bson:"isBusyInCall"`
	BillingInfo             BillingInfo        `json:"billingInfo" bson:"billingInfo"`
	AccountPayable          AccountPayable     `json:"accountPayable" bson:"accountPayable"`
	CreatedBy               primitive.ObjectID `json:"createdBy" bson:"createdBy"`
	Sessions                []interface{}      `json:"sessions" bson:"sessions"`
	CreatedAt               time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt               time.Time          `json:"updatedAt" bson:"updatedAt"`

	// added fields
	Zip4        string `json:"zip4" bson:"zip4"`
	TacMaster   string `json:"tacMaster" bson:"tacMaster"`
	TacMasterId int64  `json:"tacMasterId" bson:"tacMasterId"`
}

func NewUID() string {
	a := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	s := ""
	amin := 1
	amax := 26
	rand.Seed(time.Now().UnixNano())
	for range [3]struct{}{} {
		rint := rand.Intn(amax-amin) + amin
		s += string(a[rint])
		rint = rand.Intn(9)
		s += fmt.Sprintf("%d", rint)
	}
	return s
} // ./NewUID

type Location struct {
	Type   string    `json:"type" bson:"type"`
	Coords []float32 `json:"coordinates" bson:"coordinates"`
}

type BillingInfo struct {
	Address  string   `json:"address" bson:"address"`
	Address2 string   `json:"address2" bson:"address2"`
	City     string   `json:"city" bson:"city"`
	Zip      string   `json:"zipCode" bson:"zipCode"`
	Zip4     string   `json:"zip4" bson:"zip4"`
	Location Location `json:"location" bson:"location"`
}

type AccountPayable struct {
	Name        string `json:"name" bson:"name"`
	Email       string `json:"email" bson:"email"`
	PhoneNumber string `json:"phoneNumber" bson:"phoneNumber"`
}
