package api

import (
	"encoding/json"

	"github.com/sunshineplan/database/mongodb"
)

func (c *Client) FindOne(filter interface{}, opt *mongodb.FindOneOpt, data interface{}) error {
	option := findOneOpt{Filter: filter}
	if opt != nil {
		option.Projection = opt.Projection
	}

	var res document
	if err := c.Request(findOne, option, &res); err != nil {
		return err
	}
	if res.Document == nil {
		return mongodb.ErrNoDocuments
	}
	b, _ := json.Marshal(res.Document)
	return json.Unmarshal(b, data)
}

func (c *Client) Find(filter interface{}, opt *mongodb.FindOpt, data interface{}) error {
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

func (c *Client) InsertOne(doc interface{}) (interface{}, error) {
	if doc == nil {
		return "", mongodb.ErrNilDocument
	}
	var res insertedID
	if err := c.Request(insertOne, document{doc}, &res); err != nil {
		return nil, err
	}
	if isValidObjectID(res.InsertedID) {
		id, _ := c.ObjectID(res.InsertedID)
		return id, nil
	}
	return res.InsertedID, nil
}

func (c *Client) InsertMany(docs []interface{}) (ids []interface{}, err error) {
	if docs == nil {
		return nil, mongodb.ErrNilDocument
	}

	var res insertedIDs
	if err = c.Request(insertMany, documents{docs}, &res); err != nil {
		return
	}
	for _, i := range res.InsertedIDs {
		if isValidObjectID(i) {
			id, _ := c.ObjectID(i)
			ids = append(ids, id)
		} else {
			ids = append(ids, i)
		}
	}
	return
}

func (c *Client) UpdateOne(filter, update interface{}, opt *mongodb.UpdateOpt) (res *mongodb.UpdateResult, err error) {
	if filter == nil || update == nil {
		err = mongodb.ErrNilDocument
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

func (c *Client) UpdateMany(filter, update interface{}, opt *mongodb.UpdateOpt) (res *mongodb.UpdateResult, err error) {
	if filter == nil || update == nil {
		err = mongodb.ErrNilDocument
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

func (c *Client) ReplaceOne(filter, replacement interface{}, opt *mongodb.UpdateOpt) (res *mongodb.UpdateResult, err error) {
	if filter == nil || replacement == nil {
		err = mongodb.ErrNilDocument
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
		err = mongodb.ErrNilDocument
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
		err = mongodb.ErrNilDocument
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
		return mongodb.ErrNilDocument
	}

	var res documents
	if err := c.Request(aggregate, aggregateOpt{pipeline}, &res); err != nil {
		return err
	}
	b, _ := json.Marshal(res.Documents)
	return json.Unmarshal(b, data)
}

func (c *Client) CountDocuments(filter interface{}, opt *mongodb.CountOpt) (n int64, err error) {
	if filter == nil {
		filter = mongodb.M{}
	}

	pipeline := []mongodb.M{{"$match": filter}}
	if opt != nil {
		pipeline = append(pipeline, mongodb.M{"$skip": opt.Skip})
		if opt.Limit != 0 {
			pipeline = append(pipeline, mongodb.M{"$limit": opt.Limit})
		}
	}
	pipeline = append(pipeline, mongodb.M{"$group": mongodb.M{"_id": nil, "n": mongodb.M{"$sum": 1}}})

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

func (c *Client) FindOneAndDelete(filter interface{}, opt *mongodb.FindOneOpt, data interface{}) (err error) {
	if filter == nil {
		return mongodb.ErrNilDocument
	}

	if err = c.FindOne(filter, opt, &data); err != nil {
		return
	}

	_, err = c.DeleteOne(filter)
	return
}

func (c *Client) FindOneAndReplace(filter, replacement interface{}, opt *mongodb.FindAndUpdateOpt, data interface{}) (err error) {
	if filter == nil || replacement == nil {
		return mongodb.ErrNilDocument
	}

	findOneOpt := new(mongodb.FindOneOpt)
	if opt != nil {
		findOneOpt.Projection = opt.Projection
	}
	if err = c.FindOne(filter, findOneOpt, &data); err != nil {
		return
	}

	updateOpt := new(mongodb.UpdateOpt)
	if opt != nil {
		updateOpt.Upsert = opt.Upsert
	}
	_, err = c.ReplaceOne(filter, replacement, updateOpt)
	return
}

func (c *Client) FindOneAndUpdate(filter, update interface{}, opt *mongodb.FindAndUpdateOpt, data interface{}) (err error) {
	if filter == nil || update == nil {
		return mongodb.ErrNilDocument
	}

	findOneOpt := new(mongodb.FindOneOpt)
	if opt != nil {
		findOneOpt.Projection = opt.Projection
	}
	if err = c.FindOne(filter, findOneOpt, &data); err != nil {
		return
	}

	updateOpt := new(mongodb.UpdateOpt)
	if opt != nil {
		updateOpt.Upsert = opt.Upsert
	}
	_, err = c.UpdateOne(filter, update, updateOpt)
	return
}
