package mongodb

import (
	"errors"
	"time"
)

type M map[string]interface{}

type (
	FindOneOpt struct {
		Projection interface{}
	}

	FindOpt struct {
		Projection interface{}
		Sort       interface{}
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
		Projection interface{}
		Upsert     bool
	}

	UpdateResult struct {
		MatchedCount  int64
		ModifiedCount int64
		UpsertedCount int64
		UpsertedID    interface{}
	}
)

type ObjectID interface {
	Hex() string
	Interface() interface{}
}

type Date interface {
	Time() time.Time
	Interface() interface{}
}

type Client interface {
	SetTimeout(time.Duration)
	Connect() error
	Close() error

	FindOne(filter interface{}, opt *FindOneOpt, data interface{}) error
	Find(filter interface{}, opt *FindOpt, data interface{}) error
	InsertOne(doc interface{}) (id interface{}, err error)
	InsertMany(docs []interface{}) (ids []interface{}, err error)
	UpdateOne(filter, update interface{}, opt *UpdateOpt) (*UpdateResult, error)
	UpdateMany(filter, update interface{}, opt *UpdateOpt) (*UpdateResult, error)
	ReplaceOne(filter, replacement interface{}, opt *UpdateOpt) (*UpdateResult, error)
	DeleteOne(filter interface{}) (count int64, err error)
	DeleteMany(filter interface{}) (count int64, err error)
	Aggregate(pipeline, data interface{}) error
	CountDocuments(filter interface{}, opt *CountOpt) (n int64, err error)
	FindOneAndDelete(filter interface{}, opt *FindOneOpt, data interface{}) error
	FindOneAndReplace(filter, replacement interface{}, opt *FindAndUpdateOpt, data interface{}) error
	FindOneAndUpdate(filter, update interface{}, opt *FindAndUpdateOpt, data interface{}) error

	ObjectID(string) (ObjectID, error)
	Date(time.Time) (Date, error)
}

var (
	ErrNilDocument = errors.New("document is nil")
	ErrNoDocuments = errors.New("mongo: no documents in result")
	ErrDecodeToNil = errors.New("cannot Decode to nil value")
)
