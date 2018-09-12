package model

import (
	"gopkg.in/mgo.v2/bson"
)

// Represents a user, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type User struct {
	ID         bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Name       string        `bson:"name" json:"name"`
	Password   string        `bson:"password" json:"password"`
	Role       string        `bson:"role" json:"role"`
	Created_At string        `bson:"created_at" json:"created_at"`
}
