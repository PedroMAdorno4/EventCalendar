package create

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	ID          primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	OwnerID     primitive.ObjectID   `json:"ownerID,omitempty" bson:"ownerID,omitempty"`
	Guests      []primitive.ObjectID `json:"guests,omitempty" bson:"guests,omitempty"`
	Title       string               `json:"title" bson:"title"`
	Description string               `json:"description" bson:"description"`
	StartDate   int64                `json:"start" bson:"start"`
	EndDate     int64                `json:"end" bson:"end"`
	// Permissions auth.Permission    `json:"permissions" bson:"permissions"`
	// GroupID     primitive.ObjectID `json:"groupID" bson:"groupID"`
}
