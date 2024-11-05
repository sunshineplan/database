package driver

import (
	"time"

	"github.com/sunshineplan/database/mongodb"
	"go.mongodb.org/mongo-driver/v2/bson"
)

var (
	_ mongodb.ObjectID = bson.NilObjectID
	_ mongodb.Time     = bson.DateTime(0)
)

func (*Client) ObjectID(s string) (mongodb.ObjectID, error) {
	return bson.ObjectIDFromHex(s)
}

func (*Client) Time(t time.Time) mongodb.Time {
	return bson.NewDateTimeFromTime(t)
}
