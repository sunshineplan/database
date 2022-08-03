package api

import (
	"encoding/hex"
	"errors"
	"strconv"
	"time"

	"github.com/sunshineplan/database/mongodb"
)

var (
	_ mongodb.ObjectID = objectID("")
	_ mongodb.Date     = date(time.Time{})
)

type objectID string

func (id objectID) Hex() string {
	return string(id)
}

func (id objectID) Interface() any {
	return mongodb.M{"$oid": id}
}

func (*Client) ObjectID(s string) (mongodb.ObjectID, error) {
	if !isValidObjectID(s) {
		return nil, errors.New("the provided string is not a valid ObjectID")
	}
	return objectID(s), nil
}

func isValidObjectID(s string) bool {
	if len(s) != 24 {
		return false
	}
	_, err := hex.DecodeString(s)
	return err == nil
}

type date time.Time

func (d date) Time() time.Time {
	return time.Time(d)
}

func (d date) Interface() any {
	return mongodb.M{"$date": mongodb.M{"$numberLong": strconv.FormatInt(time.Time(d).UnixMilli(), 10)}}
}

func (*Client) Date(t time.Time) mongodb.Date {
	return date(t)
}
