package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Note struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Title     string        `bson:"title" json:"title"`
	Done      bool          `bson:"done" json:"done"`
	CreatedAt time.Time     `bson:"created_at json:"created_at"`
}
