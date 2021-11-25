package driver

import (
	"time"

	"github.com/sunshineplan/database/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	_ mongodb.ObjectID = objectID(primitive.NilObjectID)
	_ mongodb.Date     = date(time.Time{})
)

type objectID primitive.ObjectID

func (id objectID) Hex() string {
	return primitive.ObjectID(id).Hex()
}

func (id objectID) Interface() interface{} {
	return primitive.ObjectID(id)
}

func (*Client) ObjectID(s string) (mongodb.ObjectID, error) {
	id, err := primitive.ObjectIDFromHex(s)
	if err != nil {
		return nil, err
	}
	return objectID(id), nil
}

type date time.Time

func (d date) Time() time.Time {
	return time.Time(d)
}

func (d date) Interface() interface{} {
	return time.Time(d)
}

func (*Client) Date(t time.Time) mongodb.Date {
	return date(t)
}
