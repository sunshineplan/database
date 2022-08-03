package mongodb

import (
	"errors"
	"time"
)

type M map[string]any

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
	}
)

type ObjectID interface {
	Hex() string
	Interface() any
}

type Date interface {
	Time() time.Time
	Interface() any
}

type Client interface {
	SetTimeout(time.Duration)
	Connect() error
	Close() error

	FindOne(filter any, opt *FindOneOpt, data any) error
	Find(filter any, opt *FindOpt, data any) error
	InsertOne(doc any) (id any, err error)
	InsertMany(docs []any) (ids []any, err error)
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
	Date(time.Time) Date
}

var (
	ErrNilDocument = errors.New("document is nil")
	ErrNoDocuments = errors.New("mongo: no documents in result")
	ErrDecodeToNil = errors.New("cannot Decode to nil value")
)
