package api

import (
	"encoding/json"
	"unsafe"
)

type FindOneOpt struct {
	Filter     M `json:"filter,omitempty"`
	Projection M `json:"projection,omitempty"`
}

func (c *Client) FindOne(opt FindOneOpt, data interface{}) (err error) {
	var res document
	if err = c.Request(findOne, opt, &res); err != nil {
		return
	}
	b, _ := json.Marshal(res.Document)
	return json.Unmarshal(b, data)
}

type FindOpt struct {
	Filter     M   `json:"filter,omitempty"`
	Projection M   `json:"projection,omitempty"`
	Sort       M   `json:"sort,omitempty"`
	Limit      int `json:"limit,omitempty"`
	Skip       int `json:"skip,omitempty"`
}

func (c *Client) Find(opt FindOpt, data interface{}) (err error) {
	var res documents
	if err = c.Request(find, opt, &res); err != nil {
		return
	}
	b, _ := json.Marshal(res.Documents)
	return json.Unmarshal(b, data)
}

func (c *Client) InsertOne(doc Document) (id string, err error) {
	var res insertedId
	if err = c.Request(insertOne, document{M(doc)}, &res); err != nil {
		return
	}
	id = res.InsertedId
	return
}

func (c *Client) InsertMany(docs []Document) (ids []string, err error) {
	var res insertedIds
	if err = c.Request(insertMany, documents{*(*[]M)(unsafe.Pointer(&docs))}, &res); err != nil {
		return
	}
	ids = res.InsertedIds
	return
}

type UpdateOpt struct {
	Filter M    `json:"filter,omitempty"`
	Update M    `json:"update,omitempty"`
	Upsert bool `json:"upsert,omitempty"`
}

func (c *Client) UpdateOne(opt UpdateOpt) (res Result, err error) {
	if err = c.Request(updateOne, opt, &res); err != nil {
		return
	}
	return
}

func (c *Client) UpdateMany(opt UpdateOpt) (res Result, err error) {
	if err = c.Request(updateMany, opt, &res); err != nil {
		return
	}
	return
}

type ReplaceOneOpt struct {
	Filter      M    `json:"filter,omitempty"`
	Replacement M    `json:"replacement,omitempty"`
	Upsert      bool `json:"upsert,omitempty"`
}

func (c *Client) ReplaceOne(opt ReplaceOneOpt, data interface{}) (res Result, err error) {
	if err = c.Request(replaceOne, opt, &res); err != nil {
		return
	}
	return
}

type DeleteOpt struct {
	Filter M `json:"filter,omitempty"`
}

func (c *Client) DeleteOne(opt DeleteOpt) (count int, err error) {
	var res deletedCount
	if err = c.Request(deleteOne, opt, &res); err != nil {
		return
	}
	count = res.DeletedCount
	return
}

func (c *Client) DeleteMany(opt DeleteOpt) (count int, err error) {
	var res deletedCount
	if err = c.Request(deleteMany, opt, &res); err != nil {
		return
	}
	count = res.DeletedCount
	return
}

type AggregateOpt struct {
	Pipeline []M `json:"pipeline,omitempty"`
}

func (c *Client) Aggregate(opt AggregateOpt, data interface{}) (err error) {
	var res documents
	if err = c.Request(aggregate, opt, &res); err != nil {
		return
	}
	b, _ := json.Marshal(res.Documents)
	return json.Unmarshal(b, data)
}
