package api

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

type (
	document struct {
		Document interface{}
	}
	documents struct {
		Documents interface{}
	}
	insertedID   struct{ InsertedID string }
	insertedIDs  struct{ InsertedIDs []string }
	deletedCount struct{ DeletedCount int64 }
)

type (
	findOneOpt struct {
		Filter     interface{}
		Projection interface{}
	}

	findOpt struct {
		Filter     interface{}
		Projection interface{}
		Sort       interface{}
		Limit      int64
		Skip       int64
	}

	updateOpt struct {
		Filter interface{}
		Update interface{}
		Upsert bool
	}

	replaceOneOpt struct {
		Filter      interface{}
		Replacement interface{}
		Upsert      bool
	}

	deleteOpt struct {
		Filter interface{}
	}

	aggregateOpt struct {
		Pipeline interface{}
	}
)
