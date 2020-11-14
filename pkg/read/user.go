package read

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	Pending   []primitive.ObjectID `json:"pending" bson:"pending"`
	Firstname string               `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname  string               `json:"lastname,omitempty" bson:"lastname,omitempty"`
	Username  string               `json:"username" bson:"username"`
	// GroupID   primitive.ObjectID `json:"groupID" bson:"groupID"`
}
