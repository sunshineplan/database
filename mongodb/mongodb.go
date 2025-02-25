package mongodb

import (
	"errors"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type (
	M   = bson.M
	OID = bson.ObjectID
)

func OIDFromHex(s string) (OID, error) { return bson.ObjectIDFromHex(s) }

type (
	FindOneOpt struct {
		Projection any
	}

	FindOpt struct {
		Projection any
		Sort       any
		Limit      int64
		Skip       int64
	}

	UpdateOpt struct {
		Upsert bool
	}

	CountOpt struct {
		Limit int64
		Skip  int64
	}

	FindAndUpdateOpt struct {
		Projection any
		Upsert     bool
	}

	UpdateResult struct {
		MatchedCount  int64
		ModifiedCount int64
		UpsertedCount int64
		UpsertedID    any
		Acknowledged  bool
	}
)

type ObjectID interface {
	Hex() string
}

type Time interface {
	Time() time.Time
}

type Client interface {
	SetTimeout(time.Duration)
	Connect() error
	Close() error

	FindOne(filter any, opt *FindOneOpt, data any) error
	Find(filter any, opt *FindOpt, data any) error
	InsertOne(doc any) (id any, err error)
	InsertMany(docs any) (ids []any, err error)
	UpdateOne(filter, update any, opt *UpdateOpt) (*UpdateResult, error)
	UpdateMany(filter, update any, opt *UpdateOpt) (*UpdateResult, error)
	ReplaceOne(filter, replacement any, opt *UpdateOpt) (*UpdateResult, error)
	DeleteOne(filter any) (count int64, err error)
	DeleteMany(filter any) (count int64, err error)
	Aggregate(pipeline, data any) error
	CountDocuments(filter any, opt *CountOpt) (n int64, err error)
	FindOneAndDelete(filter any, opt *FindOneOpt, data any) error
	FindOneAndReplace(filter, replacement any, opt *FindAndUpdateOpt, data any) error
	FindOneAndUpdate(filter, update any, opt *FindAndUpdateOpt, data any) error

	ObjectID(string) (ObjectID, error)
	Time(time.Time) Time
}

var (
	ErrNilDocument = errors.New("mongo: document is nil")
	ErrNoDocuments = errors.New("mongo: no documents in result")
)

type InvalidDecodeError struct {
	Type reflect.Type
}

func (e *InvalidDecodeError) Error() string {
	if e.Type == nil {
		return "mongo: Decode(nil)"
	}
	if e.Type.Kind() != reflect.Pointer {
		return "mongo: Decode(non-pointer " + e.Type.String() + ")"
	}
	return "mongo: Decode(nil " + e.Type.String() + ")"
}
