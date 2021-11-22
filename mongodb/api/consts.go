package api

import "errors"

var (
	ErrNilDocument = errors.New("document is nil")
	ErrNoDocuments = errors.New("mongo: no documents in result")
	ErrDecodeToNil = errors.New("cannot Decode to nil value")
)

const base = "https://data.mongodb-api.com/app/%s/endpoint/data/beta"

const (
	findOne    = "/action/findOne"
	find       = "/action/find"
	insertOne  = "/action/insertOne"
	insertMany = "/action/insertMany"
	updateOne  = "/action/updateOne"
	updateMany = "/action/updateMany"
	replaceOne = "/action/replaceOne"
	deleteOne  = "/action/deleteOne"
	deleteMany = "/action/deleteMany"
	aggregate  = "/action/aggregate"
)
