package api

import (
	"encoding/json"
)

func (c *Client) FindOne(filter interface{}, opt *FindOneOpt, data interface{}) error {
	if filter == nil {
		return ErrNilDocument
	}
	if data == nil {
		return ErrDecodeToNil
	}

	option := findOneOpt{Filter: filter}
	if opt != nil {
		option.Projection = opt.Projection
	}

	var res document
	if err := c.Request(findOne, option, &res); err != nil {
		return err
	}
	if res.Document == nil {
		return ErrNoDocuments
	}
	b, _ := json.Marshal(res.Document)
	return json.Unmarshal(b, data)
}

func (c *Client) Find(filter interface{}, opt *FindOpt, data interface{}) error {
	if filter == nil {
		return ErrNilDocument
	}
	if data == nil {
		return ErrDecodeToNil
	}

	option := findOpt{Filter: filter}
	if opt != nil {
		option.Projection = opt.Projection
		option.Sort = opt.Sort
		option.Limit = opt.Limit
		option.Skip = opt.Skip
	}

	var res documents
	if err := c.Request(find, option, &res); err != nil {
		return err
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
	if docs == nil {
		return nil, ErrNilDocument
	}

	var res insertedIds
	if err = c.Request(insertMany, documents{docs}, &res); err != nil {
		return
	}
	ids = res.InsertedIds
	return
}

func (c *Client) UpdateOne(filter, update interface{}, opt *UpdateOpt) (res Result, err error) {
	if filter == nil || update == nil {
		err = ErrNilDocument
		return
	}

	option := updateOpt{Filter: filter, Update: update}
	if opt != nil {
		option.Upsert = opt.Upsert
	}

	if err = c.Request(updateOne, option, &res); err != nil {
		return
	}
	return
}

func (c *Client) UpdateMany(filter, update interface{}, opt *UpdateOpt) (res Result, err error) {
	if filter == nil || update == nil {
		err = ErrNilDocument
		return
	}

	option := updateOpt{Filter: filter, Update: update}
	if opt != nil {
		option.Upsert = opt.Upsert
	}

	if err = c.Request(updateMany, option, &res); err != nil {
		return
	}
	return
}

func (c *Client) ReplaceOne(filter, replacement interface{}, opt *UpdateOpt) (res Result, err error) {
	if filter == nil || replacement == nil {
		err = ErrNilDocument
		return
	}

	option := replaceOneOpt{Filter: filter, Replacement: replacement}
	if opt != nil {
		option.Upsert = opt.Upsert
	}

	if err = c.Request(replaceOne, option, &res); err != nil {
		return
	}
	return
}

func (c *Client) DeleteOne(filter interface{}) (count int, err error) {
	if filter == nil {
		err = ErrNilDocument
		return
	}

	var res deletedCount
	if err = c.Request(deleteOne, deleteOpt{filter}, &res); err != nil {
		return
	}
	count = res.DeletedCount
	return
}

func (c *Client) DeleteMany(filter interface{}) (count int, err error) {
	if filter == nil {
		err = ErrNilDocument
		return
	}

	var res deletedCount
	if err = c.Request(deleteMany, deleteOpt{filter}, &res); err != nil {
		return
	}
	count = res.DeletedCount
	return
}

func (c *Client) Aggregate(pipeline, data interface{}) error {
	if pipeline == nil {
		return ErrNilDocument
	}
	if data == nil {
		return ErrDecodeToNil
	}

	var res documents
	if err := c.Request(aggregate, aggregateOpt{pipeline}, &res); err != nil {
		return err
	}
	b, _ := json.Marshal(res.Documents)
	return json.Unmarshal(b, data)
}
