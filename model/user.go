package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Represents a user, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type User struct {
	ID         bson.ObjectId `bson:"_id" json:"id"`
	Name       string        `bson:"name" json:"name"`
	Password   string        `bson:"password" json:"password"`
	Role       string        `bson:"role" json:"role"`
	Created_At *time.Time    `bson:"created_at" json:"created_at"`
}
