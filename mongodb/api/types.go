package api

type M map[string]interface{}

type (
	document struct {
		Document interface{}
	}
	documents struct {
		Documents interface{}
	}
	insertedId   struct{ InsertedId string }
	insertedIds  struct{ InsertedIds []string }
	deletedCount struct{ DeletedCount int64 }
)

type (
	findOneOpt struct {
		Filter     interface{}
		Projection interface{}
	}
	FindOneOpt struct {
		Projection interface{}
	}

	findOpt struct {
		Filter     interface{}
		Projection interface{}
		Sort       interface{}
		Limit      int64
		Skip       int64
	}
	FindOpt struct {
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
	UpdateOpt struct {
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

	CountOpt struct {
		Limit int64
		Skip  int64
	}
)

type Result struct {
	MatchedCount  int64
	ModifiedCount int64
	UpsertedId    string
}
