package driver

import (
	"time"

	"github.com/sunshineplan/database/mongodb"
	"go.mongodb.org/mongo-driver/v2/bson"
)

var (
	_ mongodb.ObjectID = objectID(bson.NilObjectID)
	_ mongodb.Date     = date(time.Time{})
)

type objectID bson.ObjectID

func (id objectID) Hex() string {
	return bson.ObjectID(id).Hex()
}

func (id objectID) Interface() any {
	return bson.ObjectID(id)
}

func (*Client) ObjectID(s string) (mongodb.ObjectID, error) {
	id, err := bson.ObjectIDFromHex(s)
	if err != nil {
		return nil, err
	}
	return objectID(id), nil
}

type date time.Time

func (d date) Time() time.Time {
	return time.Time(d)
}

func (d date) Interface() any {
	return time.Time(d)
}

func (*Client) Date(t time.Time) mongodb.Date {
	return date(t)
}
