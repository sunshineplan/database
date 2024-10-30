package driver

import (
	"encoding/json"
	"time"

	"github.com/sunshineplan/database/mongodb"
	"go.mongodb.org/mongo-driver/v2/bson"
)

var (
	_ mongodb.ObjectID = bson.NilObjectID
	_ mongodb.Date     = date(time.Time{})
)

func (*Client) ObjectID(s string) (mongodb.ObjectID, error) {
	return bson.ObjectIDFromHex(s)
}

type date time.Time

func (d date) Time() time.Time {
	return time.Time(d)
}

func (d date) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(d))
}

func (*Client) Date(t time.Time) mongodb.Date {
	return date(t)
}
