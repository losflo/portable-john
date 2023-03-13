package portajohn

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID              primitive.ObjectID  `json:"_id" bson:"_id"`
	Title           string              `json:"title" bson:"title"`
	PID             string              `json:"productRandomId" bson:"productRandomId"`
	Images          []string            `json:"images" bson:"images"`
	Specs           Specs               `json:"specs" bson:"specs"`
	Features        Features            `json:"features" bson:"features"`
	Requirements    Requirements        `json:"requirements" bson:"requirements"`
	Pricing         Pricing             `json:"pricing" bson:"pricing"`
	Description     string              `json:"description" bson:"description"`
	QuantityInStock int64               `json:"quantityInStock" bson:"quantityInStock"`
	IsDeleted       bool                `json:"isDeleted" bson:"isDeleted"`
	Status          int64               `json:"status" bson:"status"`
	CategoryID      *primitive.ObjectID `json:"categoryId" bson:"categoryId"`
	LocationToken   string              `json:"locationToken" bson:"locationToken"`
	CreatedAt       time.Time           `json:"createdAt" bson:"createdAt"`
	UpdatedAt       time.Time           `json:"updatedAt" bson:"updatedAt"`
}

type Specs struct {
	Height        int32 `json:"height" bson:"height"`
	Length        int32 `json:"length" bson:"length"`
	Width         int32 `json:"width" bson:"width"`
	HoseLenth     int32 `json:"hoseLength" bson:"hoseLength"`
	WasteCapacity int32 `json:"wasteCapacity" bson:"wasteCapacity"`
	WaterCapacity int32 `json:"waterCapacity" bson:"waterCapacity"`
}

type Features struct {
	Towel            bool  `json:"towel" bson:"towel"`
	SecurityLock     bool  `json:"securityLock" bson:"securityLock"`
	ElectricSystem   bool  `json:"electricSystem" bson:"electricSystem"`
	ToiletPaperRolls int32 `json:"toiletPaperRolls" bson:"toiletPaperRolls"`
	HandSanitizers   int32 `json:"handSanitizers" bson:"handSanitizers"`
}

type Requirements struct {
	ServiceIntervalInDays int32 `json:"serviceIntervalInDays" bson:"serviceIntervalInDays"`
	Blu                   int32 `json:"blu" bson:"blu"`
	Chemical              int32 `json:"chemical" bson:"chemical"`
	Gloves                bool  `json:"gloves" bson:"gloves"`
}

type Pricing struct {
	Price              int32 `json:"price" bson:"price"`
	SubTotal           int32 `json:"subTotal" bson:"subTotal"`
	DeliveryFee        int32 `json:"deliveryFee" bson:"deliveryFee"`
	ServiceCharges     int32 `json:"serviceCharges" bson:"serviceCharges"`
	TaxPercentage      int32 `json:"taxPercentage" bson:"taxPercentage"`
	TaxAmount          int32 `json:"taxAmount" bson:"taxAmount"`
	Total              int32 `json:"total" bson:"total"`
	DiscountPercentage int32 `json:"discountPercentage" bson:"discountPercentage"`
	DiscountAmount     int32 `json:"discountAmount" bson:"discountAmount"`
	DiscountedPrice    int32 `json:"discountedPrice" bson:"discountedPrice"`
}

type ProductLocation struct {
	ID           primitive.ObjectID  `json:"_id" bson:"_id"`
	ProductTitle string              `json:"productTitle" bson:"productTitle"`
	ProductID    primitive.ObjectID  `json:"productId" bson:"productId"`
	CustomerID   *primitive.ObjectID `json:"customerId" bson:"customerId"`
	SiteID       *primitive.ObjectID `json:"siteId" bson:"siteId"`

	CreatedAt      time.Time  `json:"createdAt" bson:"createdAt"`
	UpdatedAt      time.Time  `json:"updatedAt" bson:"updatedAt"`
	MaintenancedAt *time.Time `json:"maintenancedAt" bson:"maintenancedAt"`
}
