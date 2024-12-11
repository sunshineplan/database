package driver

import (
	"reflect"

	"github.com/sunshineplan/database/mongodb"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func (c *Client) FindOne(filter any, opt *mongodb.FindOneOpt, data any) error {
	if filter == nil {
		filter = mongodb.M{}
	}
	if rv := reflect.ValueOf(data); rv.Kind() != reflect.Pointer {
		return &mongodb.InvalidDecodeError{Type: reflect.TypeOf(data)}
	}

	option := options.FindOne()
	if opt != nil {
		option.SetProjection(opt.Projection)
	}

	ctx, cancel := c.context()
	defer cancel()

	err := c.coll.FindOne(ctx, filter, option).Decode(data)
	if err == mongo.ErrNoDocuments {
		return mongodb.ErrNoDocuments
	}

	return err
}

func (c *Client) Find(filter any, opt *mongodb.FindOpt, data any) error {
	if filter == nil {
		filter = mongodb.M{}
	}
	if rv := reflect.ValueOf(data); rv.Kind() != reflect.Pointer {
		return &mongodb.InvalidDecodeError{Type: reflect.TypeOf(data)}
	}

	option := options.Find()
	if opt != nil {
		option.SetProjection(opt.Projection)
		option.SetSort(opt.Sort)
		option.SetLimit(opt.Limit)
		option.SetSkip(opt.Skip)
	}

	ctx, cancel := c.context()
	defer cancel()

	cur, err := c.coll.Find(ctx, filter, option)
	if err != nil {
		return err
	}
	return cur.All(ctx, data)
}

func (c *Client) InsertOne(doc any) (any, error) {
	if doc == nil {
		return "", mongodb.ErrNilDocument
	}

	ctx, cancel := c.context()
	defer cancel()

	res, err := c.coll.InsertOne(ctx, doc)
	if err != nil {
		return "", err
	}
	return res.InsertedID, nil
}

func (c *Client) InsertMany(docs any) ([]any, error) {
	if docs == nil {
		return nil, mongodb.ErrNilDocument
	}

	ctx, cancel := c.context()
	defer cancel()

	res, err := c.coll.InsertMany(ctx, docs)
	if err != nil {
		return nil, err
	}
	return res.InsertedIDs, nil
}

func (c *Client) UpdateOne(filter, update any, opt *mongodb.UpdateOpt) (*mongodb.UpdateResult, error) {
	if filter == nil || update == nil {
		return nil, mongodb.ErrNilDocument
	}

	option := options.UpdateOne()
	if opt != nil {
		option.SetUpsert(opt.Upsert)
	}

	ctx, cancel := c.context()
	defer cancel()

	res, err := c.coll.UpdateOne(ctx, filter, update, option)
	if err != nil {
		return nil, err
	}
	return (*mongodb.UpdateResult)(res), nil
}

func (c *Client) UpdateMany(filter, update any, opt *mongodb.UpdateOpt) (*mongodb.UpdateResult, error) {
	if filter == nil || update == nil {
		return nil, mongodb.ErrNilDocument
	}

	option := options.UpdateMany()
	if opt != nil {
		option.SetUpsert(opt.Upsert)
	}

	ctx, cancel := c.context()
	defer cancel()

	res, err := c.coll.UpdateMany(ctx, filter, update, option)
	if err != nil {
		return nil, err
	}
	return (*mongodb.UpdateResult)(res), nil
}

func (c *Client) ReplaceOne(filter, replacement any, opt *mongodb.UpdateOpt) (*mongodb.UpdateResult, error) {
	if filter == nil || replacement == nil {
		return nil, mongodb.ErrNilDocument
	}

	option := options.Replace()
	if opt != nil {
		option.SetUpsert(opt.Upsert)
	}

	ctx, cancel := c.context()
	defer cancel()

	res, err := c.coll.ReplaceOne(ctx, filter, replacement, option)
	if err != nil {
		return nil, err
	}
	return (*mongodb.UpdateResult)(res), nil
}

func (c *Client) DeleteOne(filter any) (int64, error) {
	if filter == nil {
		return 0, mongodb.ErrNilDocument
	}

	ctx, cancel := c.context()
	defer cancel()

	res, err := c.coll.DeleteOne(ctx, filter)
	if err != nil {
		return 0, err
	}
	return res.DeletedCount, nil
}

func (c *Client) DeleteMany(filter any) (int64, error) {
	if filter == nil {
		return 0, mongodb.ErrNilDocument
	}

	ctx, cancel := c.context()
	defer cancel()

	res, err := c.coll.DeleteMany(ctx, filter)
	if err != nil {
		return 0, err
	}
	return res.DeletedCount, nil
}

func (c *Client) Aggregate(pipeline, data any) error {
	if pipeline == nil {
		return mongodb.ErrNilDocument
	}
	if rv := reflect.ValueOf(data); rv.Kind() != reflect.Pointer {
		return &mongodb.InvalidDecodeError{Type: reflect.TypeOf(data)}
	}

	ctx, cancel := c.context()
	defer cancel()

	cur, err := c.coll.Aggregate(ctx, pipeline)
	if err != nil {
		return err
	}
	return cur.All(ctx, data)
}

func (c *Client) CountDocuments(filter any, opt *mongodb.CountOpt) (int64, error) {
	if filter == nil {
		filter = mongodb.M{}
	}

	option := options.Count()
	if opt != nil {
		option.SetLimit(opt.Limit)
		option.SetSkip(opt.Skip)
	}

	ctx, cancel := c.context()
	defer cancel()

	return c.coll.CountDocuments(ctx, filter, option)
}

func (c *Client) FindOneAndDelete(filter any, opt *mongodb.FindOneOpt, data any) error {
	if filter == nil {
		return mongodb.ErrNilDocument
	}
	if rv := reflect.ValueOf(data); rv.Kind() != reflect.Pointer {
		return &mongodb.InvalidDecodeError{Type: reflect.TypeOf(data)}
	}

	option := options.FindOneAndDelete()
	if opt != nil {
		option.SetProjection(opt.Projection)
	}

	ctx, cancel := c.context()
	defer cancel()

	err := c.coll.FindOneAndDelete(ctx, filter, option).Decode(data)
	if err == mongo.ErrNoDocuments {
		return mongodb.ErrNoDocuments
	}

	return err
}

func (c *Client) FindOneAndReplace(filter, replacement any, opt *mongodb.FindAndUpdateOpt, data any) error {
	if filter == nil || replacement == nil {
		return mongodb.ErrNilDocument
	}
	if rv := reflect.ValueOf(data); rv.Kind() != reflect.Pointer {
		return &mongodb.InvalidDecodeError{Type: reflect.TypeOf(data)}
	}

	option := options.FindOneAndReplace()
	if opt != nil {
		option.SetProjection(opt.Projection)
		option.SetUpsert(opt.Upsert)
	}

	ctx, cancel := c.context()
	defer cancel()

	err := c.coll.FindOneAndReplace(ctx, filter, replacement, option).Decode(data)
	if err == mongo.ErrNoDocuments {
		return mongodb.ErrNoDocuments
	}

	return err
}

func (c *Client) FindOneAndUpdate(filter, update any, opt *mongodb.FindAndUpdateOpt, data any) error {
	if filter == nil || update == nil {
		return mongodb.ErrNilDocument
	}
	if rv := reflect.ValueOf(data); rv.Kind() != reflect.Pointer {
		return &mongodb.InvalidDecodeError{Type: reflect.TypeOf(data)}
	}

	option := options.FindOneAndUpdate()
	if opt != nil {
		option.SetProjection(opt.Projection)
		option.SetUpsert(opt.Upsert)
	}

	ctx, cancel := c.context()
	defer cancel()

	err := c.coll.FindOneAndUpdate(ctx, filter, update, option).Decode(data)
	if err == mongo.ErrNoDocuments {
		return mongodb.ErrNoDocuments
	}

	return err
}
