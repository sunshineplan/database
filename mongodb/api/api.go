package api

const base = "https://data.mongodb-api.com/app/%s/endpoint/data/%s"

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
		Document any
	}
	documents struct {
		Documents any
	}
	insertedID   struct{ InsertedID string }
	insertedIDs  struct{ InsertedIDs []string }
	deletedCount struct{ DeletedCount int64 }
)

type (
	findOneOpt struct {
		Filter     any
		Projection any
	}

	findOpt struct {
		Filter     any
		Projection any
		Sort       any
		Limit      int64
		Skip       int64
	}

	updateOpt struct {
		Filter any
		Update any
		Upsert bool
	}

	replaceOneOpt struct {
		Filter      any
		Replacement any
		Upsert      bool
	}

	deleteOpt struct {
		Filter any
	}

	aggregateOpt struct {
		Pipeline any
	}
)
