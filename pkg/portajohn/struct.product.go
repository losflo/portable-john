package portajohn

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID              primitive.ObjectID  `json:"_id" bson:"_id"`
	Title           string              `json:"title" bson:"title"`
	PID             string              `json:"productRandomId" bson:"productRandomId"`
	Description     string              `json:"description" bson:"description"`
	QuantityInStock int64               `json:"quantityInStock" bson:"quantityInStock"`
	IsDeleted       bool                `json:"isDeleted" bson:"isDeleted"`
	Status          int64               `json:"status" bson:"status"`
	CategoryID      *primitive.ObjectID `json:"categoryId" bson:"categoryId"`
	LocationToken   string              `json:"locationToken" bson:"locationToken"`
	CreatedAt       time.Time           `json:"createdAt" bson:"createdAt"`
	UpdatedAt       time.Time           `json:"updatedAt" bson:"updatedAt"`
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
