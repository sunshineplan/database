package driver

import (
	"encoding/json"
	"time"

	"github.com/sunshineplan/database/mongodb"
	"go.mongodb.org/mongo-driver/v2/bson"
)

var (
	_ mongodb.ObjectID = objectID{bson.NilObjectID}
	_ mongodb.Date     = date(time.Time{})
)

func (*Client) ObjectID(s string) (mongodb.ObjectID, error) {
	oid, err := bson.ObjectIDFromHex(s)
	if err != nil {
		return nil, err
	}
	return objectID{oid}, nil
}

type objectID struct{ bson.ObjectID }

func (oid objectID) MarshalBSONValue() (typ byte, data []byte, err error) {
	b, err := oid.MarshalJSON()
	if err != nil {
		return
	}
	return byte(bson.TypeObjectID), b, nil
}

type date time.Time

func (d date) Time() time.Time {
	return time.Time(d)
}

func (d date) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(d))
}

func (d date) MarshalBSONValue() (typ byte, data []byte, err error) {
	b, err := d.MarshalJSON()
	if err != nil {
		return
	}
	return byte(bson.TypeDateTime), b, nil
}

func (*Client) Date(t time.Time) mongodb.Date {
	return date(t)
}
