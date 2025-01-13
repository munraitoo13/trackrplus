package subscriptions

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Subscription struct {
	ID     primitive.ObjectID `json:"id" bson:"_id"`
	UserID primitive.ObjectID `json:"userID" bson:"userID"`

	Name          string    `json:"name" bson:"name"`
	Price         float64   `json:"price" bson:"price"`
	PaymentMethod string    `json:"paymentMethod" bson:"paymentMethod"`
	CardLast4     *string   `json:"cardLast4" bson:"cardLast4"`
	Paid          bool      `json:"paid" bson:"paid"`
	RenewalDate   time.Time `json:"renewalDate" bson:"renewalDate"`
	Notes         *string   `json:"notes" bson:"notes"`

	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
}

type Subscriptions []Subscription
