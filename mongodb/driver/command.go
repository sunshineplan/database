package driver

import (
	"context"

	"github.com/sunshineplan/database/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (c *Client) FindOne(filter any, opt *mongodb.FindOneOpt, data any) error {
	if filter == nil {
		filter = mongodb.M{}
	}
	if data == nil {
		return mongodb.ErrDecodeToNil
	}

	option := options.FindOne()
	if opt != nil {
		option.Projection = opt.Projection
	}

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
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
	if data == nil {
		return mongodb.ErrDecodeToNil
	}

	option := options.Find()
	if opt != nil {
		option.Projection = opt.Projection
		option.Sort = opt.Sort
		option.Limit = &opt.Limit
		option.Skip = &opt.Skip
	}

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
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

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	res, err := c.coll.InsertOne(ctx, doc)
	if err != nil {
		return "", err
	}
	if id, ok := res.InsertedID.(primitive.ObjectID); ok {
		return objectID(id), nil
	}
	return res.InsertedID, nil
}

func (c *Client) InsertMany(docs []any) ([]any, error) {
	if docs == nil {
		return nil, mongodb.ErrNilDocument
	}

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	res, err := c.coll.InsertMany(ctx, docs)
	if err != nil {
		return nil, err
	}

	var ids []any
	for _, i := range res.InsertedIDs {
		if id, ok := i.(primitive.ObjectID); ok {
			ids = append(ids, objectID(id))
		} else {
			ids = append(ids, i)
		}
	}
	return ids, nil
}

func (c *Client) UpdateOne(filter, update any, opt *mongodb.UpdateOpt) (*mongodb.UpdateResult, error) {
	if filter == nil || update == nil {
		return nil, mongodb.ErrNilDocument
	}

	option := options.Update()
	if opt != nil {
		option.Upsert = &opt.Upsert
	}

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	res, err := c.coll.UpdateOne(ctx, filter, update, option)
	if err != nil {
		return nil, err
	}
	if id, ok := res.UpsertedID.(primitive.ObjectID); ok {
		res.UpsertedID = objectID(id)
	}

	return (*mongodb.UpdateResult)(res), nil
}

func (c *Client) UpdateMany(filter, update any, opt *mongodb.UpdateOpt) (*mongodb.UpdateResult, error) {
	if filter == nil || update == nil {
		return nil, mongodb.ErrNilDocument
	}

	option := options.Update()
	if opt != nil {
		option.Upsert = &opt.Upsert
	}

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	res, err := c.coll.UpdateMany(ctx, filter, update, option)
	if err != nil {
		return nil, err
	}
	if id, ok := res.UpsertedID.(primitive.ObjectID); ok {
		res.UpsertedID = objectID(id)
	}

	return (*mongodb.UpdateResult)(res), nil
}

func (c *Client) ReplaceOne(filter, replacement any, opt *mongodb.UpdateOpt) (*mongodb.UpdateResult, error) {
	if filter == nil || replacement == nil {
		return nil, mongodb.ErrNilDocument
	}

	option := options.Replace()
	if opt != nil {
		option.Upsert = &opt.Upsert
	}

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	res, err := c.coll.ReplaceOne(ctx, filter, replacement, option)
	if err != nil {
		return nil, err
	}
	if id, ok := res.UpsertedID.(primitive.ObjectID); ok {
		res.UpsertedID = objectID(id)
	}

	return (*mongodb.UpdateResult)(res), nil
}

func (c *Client) DeleteOne(filter any) (int64, error) {
	if filter == nil {
		return 0, mongodb.ErrNilDocument
	}

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
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

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
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
	if data == nil {
		return mongodb.ErrDecodeToNil
	}

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
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
		option.Limit = &opt.Limit
		option.Skip = &opt.Skip
	}

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return c.coll.CountDocuments(ctx, filter, option)
}

func (c *Client) FindOneAndDelete(filter any, opt *mongodb.FindOneOpt, data any) error {
	if filter == nil {
		return mongodb.ErrNilDocument
	}
	if data == nil {
		return mongodb.ErrDecodeToNil
	}

	option := options.FindOneAndDelete()
	if opt != nil {
		option.Projection = opt.Projection
	}

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
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
	if data == nil {
		return mongodb.ErrDecodeToNil
	}

	option := options.FindOneAndReplace()
	if opt != nil {
		option.Projection = opt.Projection
		option.Upsert = &opt.Upsert
	}

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
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
	if data == nil {
		return mongodb.ErrDecodeToNil
	}

	option := options.FindOneAndUpdate()
	if opt != nil {
		option.Projection = opt.Projection
		option.Upsert = &opt.Upsert
	}

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	err := c.coll.FindOneAndUpdate(ctx, filter, update, option).Decode(data)
	if err == mongo.ErrNoDocuments {
		return mongodb.ErrNoDocuments
	}

	return err
}
