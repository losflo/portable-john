package portajohn

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID                      primitive.ObjectID      `bson:"_id" json:"id"`
	QID                     string                  `bson:"quoteRandomId" json:"quoteRandomId"`
	OID                     string                  `bson:"orderRandomId" json:"orderRandomId"`
	StopID                  string                  `bson:"stopRandomId" json:"stopRandomId"`
	ZipCode                 string                  `bson:"zipCode" json:"zipCode"`
	RentalDurationType      int32                   `bson:"rentalDurationType" json:"rentalDurationType"`
	ServiceInterval         int32                   `bson:"serviceInterval" json:"serviceInterval"`
	DeliveryType            int32                   `bson:"deliveryType" json:"deliveryType"`
	DeliveryDate            Date                    `bson:"deliveryDate" json:"deliveryDate"`
	PickupDate              Date                    `bson:"pickupDate" json:"pickupDate"`
	SpecialNeeds            string                  `bson:"specialNeeds" json:"specialNeeds"`
	SiteDetails             SiteDetails             `bson:"siteDetails" json:"siteDetails"`
	IsSameBillingInfo       bool                    `bson:"isSameBillingInfo" json:"isSameBillingInfo"`
	BillingInfo             BillingInfo             `bson:"billingInfo" json:"billingInfo"`
	AccountPayable          AccountPayable          `bson:"accountPayable" json:"accountPayable"`
	Products                []QuoteProduct          `bson:"products" json:"products"`
	Services                []interface{}           `bson:"services" json:"services"`
	Pricing                 Pricing                 `bson:"pricing" json:"pricing"`
	Status                  int32                   `bson:"status" json:"status"`
	OrderStatus             int32                   `bson:"orderStatus" json:"orderStatus"`
	StepNumber              int32                   `bson:"stepNumber" json:"stepNumber"`
	SentTime                int                     `bson:"sentTime" json:"sentTime"`
	ExpectedExpiryTime      int                     `bson:"expectedExpiryTime" json:"expectedExpiryTime"`
	CustomerResponseTime    int                     `bson:"customerResponseTime" json:"customerResponseTime"`
	Token                   string                  `json:"token" bson:"token"`
	UnderTheBookRequestInfo UnderTheBookRequestInfo `bson:"underTheBookRequestInfo" json:"underTheBookRequestInfo"`
	CustomerId              primitive.ObjectID      `bson:"customerId" json:"customerId"`
	CreatedBy               primitive.ObjectID      `bson:"createdBy" json:"createdBy"`
	SentBy                  primitive.ObjectID      `bson:"sentBy" json:"sentBy"`
	SiteId                  primitive.ObjectID      `bson:"siteId" json:"siteId"`
	DeliveryDriverId        primitive.ObjectID      `bson:"deliveryDriverId" json:"deliveryDriverId"`
	ServiceDriverId         primitive.ObjectID      `bson:"serviceDriverId" json:"serviceDriverId"`
	RouteId                 primitive.ObjectID      `bson:"routeId" json:"routeId"`
	IsOrderDeleted          bool                    `bson:"isOrderDeleted" json:"isOrderDeleted"`
	AlgoId                  string                  `bson:"algoId" json:"algoId"`
	CustomerResponses       []interface{}           `bson:"customerResponses" json:"customerResponses"`
	UnassignedInRouting     []interface{}           `bson:"unassignedInRouting" json:"unassignedInRouting"`
	CreatedAt               time.Time               `bson:"createdAt" json:"createdAt"`
	UpdatedAt               time.Time               `bson:"updatedAt" json:"updatedAt"`
}

type UnderTheBookRequestInfo struct {
	SentBy   interface{} `bson:"sentBy" json:"sentBy"`
	SentTime int         `bson:"sentTime" json:"sentTime"`
}

type Date struct {
	Day   int32 `bson:"day" json:"day"`
	Month int32 `bson:"month" json:"month"`
	Year  int32 `bson:"year" json:"year"`
}

type SiteDetails struct {
	Title       string   `bson:"title" json:"title"`
	ContactName string   `bson:"contactName" json:"contactName"`
	Email       string   `bson:"email" json:"email"`
	PhoneNumber string   `bson:"phoneNumber" json:"phoneNumber"`
	Address     string   `bson:"address" json:"address"`
	City        string   `bson:"city" json:"city"`
	ZipCode     string   `bson:"zipCode" json:"zipCode"`
	Location    Location `bson:"location" json:"location"`
}

type QuoteProduct struct {
	ID                   primitive.ObjectID `bson:"_id" json:"id"`
	PID                  string             `bson:"productId" json:"productId"`
	Title                string             `bson:"title" json:"title"`
	Image                string             `bson:"image" json:"image"`
	CategoryName         string             `bson:"categoryName" json:"categoryName"`
	Quantity             int32              `bson:"quantity" json:"quantity"`
	PerItemActualCost    float64            `bson:"perItemActualCost" json:"perItemActualCost"`
	PerItemOfferedAmount float64            `bson:"perItemOfferedAmount" json:"perItemOfferedAmount"`
}
