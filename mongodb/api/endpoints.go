package api

import (
	"encoding/json"
)

type FindOneOpt struct {
	Filter     interface{} `json:"filter,omitempty"`
	Projection interface{} `json:"projection,omitempty"`
}

func (c *Client) FindOne(opt FindOneOpt, data interface{}) (err error) {
	var res document
	if err = c.Request(findOne, opt, &res); err != nil {
		return
	}
	if res.Document == nil {
		return ErrNoDocuments
	}
	b, _ := json.Marshal(res.Document)
	return json.Unmarshal(b, data)
}

type FindOpt struct {
	Filter     interface{} `json:"filter,omitempty"`
	Projection interface{} `json:"projection,omitempty"`
	Sort       interface{} `json:"sort,omitempty"`
	Limit      int         `json:"limit,omitempty"`
	Skip       int         `json:"skip,omitempty"`
}

func (c *Client) Find(opt FindOpt, data interface{}) (err error) {
	var res documents
	if err = c.Request(find, opt, &res); err != nil {
		return
	}
	b, _ := json.Marshal(res.Documents)
	return json.Unmarshal(b, data)
}

func (c *Client) InsertOne(doc interface{}) (id string, err error) {
	if doc == nil {
		return "", ErrNilDocument
	}
	var res insertedId
	if err = c.Request(insertOne, document{doc}, &res); err != nil {
		return
	}
	id = res.InsertedId
	return
}

func (c *Client) InsertMany(docs interface{}) (ids []string, err error) {
	var res insertedIds
	if err = c.Request(insertMany, documents{docs}, &res); err != nil {
		return
	}
	ids = res.InsertedIds
	return
}

type UpdateOpt struct {
	Filter interface{} `json:"filter,omitempty"`
	Update interface{} `json:"update,omitempty"`
	Upsert bool        `json:"upsert,omitempty"`
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
	Filter      interface{} `json:"filter,omitempty"`
	Replacement interface{} `json:"replacement,omitempty"`
	Upsert      bool        `json:"upsert,omitempty"`
}

func (c *Client) ReplaceOne(opt ReplaceOneOpt, data interface{}) (res Result, err error) {
	if err = c.Request(replaceOne, opt, &res); err != nil {
		return
	}
	return
}

type DeleteOpt struct {
	Filter interface{} `json:"filter,omitempty"`
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
	Pipeline interface{} `json:"pipeline,omitempty"`
}

func (c *Client) Aggregate(opt AggregateOpt, data interface{}) (err error) {
	var res documents
	if err = c.Request(aggregate, opt, &res); err != nil {
		return
	}
	b, _ := json.Marshal(res.Documents)
	return json.Unmarshal(b, data)
}
