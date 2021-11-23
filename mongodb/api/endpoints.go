package api

import (
	"encoding/json"
)

func (c *Client) FindOne(filter interface{}, opt *FindOneOpt, data interface{}) error {
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

func (c *Client) DeleteOne(filter interface{}) (count int64, err error) {
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

func (c *Client) DeleteMany(filter interface{}) (count int64, err error) {
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

	var res documents
	if err := c.Request(aggregate, aggregateOpt{pipeline}, &res); err != nil {
		return err
	}
	b, _ := json.Marshal(res.Documents)
	return json.Unmarshal(b, data)
}

func (c *Client) CountDocuments(filter interface{}, opt *CountOpt) (n int64, err error) {
	if filter == nil {
		filter = M{}
	}

	pipeline := []M{{"$match": filter}}
	if opt != nil {
		pipeline = append(pipeline, M{"$skip": opt.Skip})
		if opt.Limit != 0 {
			pipeline = append(pipeline, M{"$limit": opt.Limit})
		}
	}
	pipeline = append(pipeline, M{"$group": M{"_id": nil, "n": M{"$sum": 1}}})

	var res []struct{ N int64 }
	if err = c.Aggregate(pipeline, &res); err != nil {
		return
	}

	if len(res) == 0 {
		n = 0
	} else {
		n = res[0].N
	}
	return
}

func (c *Client) FindOneAndDelete(filter interface{}, opt *FindOneOpt, data interface{}) (err error) {
	if filter == nil {
		return ErrNilDocument
	}

	if err = c.FindOne(filter, opt, &data); err != nil {
		return
	}

	_, err = c.DeleteOne(filter)
	return
}

func (c *Client) FindOneAndReplace(filter, replacement interface{}, opt *FindOneOpt, data interface{}) (err error) {
	if filter == nil || replacement == nil {
		return ErrNilDocument
	}

	if err = c.FindOne(filter, opt, &data); err != nil {
		return
	}

	_, err = c.ReplaceOne(filter, replacement, nil)
	return
}

func (c *Client) FindOneAndUpdate(filter, update interface{}, opt *FindOneOpt, data interface{}) (err error) {
	if filter == nil || update == nil {
		return ErrNilDocument
	}

	if err = c.FindOne(filter, opt, &data); err != nil {
		return
	}

	_, err = c.UpdateOne(filter, update, nil)
	return
}
