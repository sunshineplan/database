package api

type M map[string]interface{}

type (
	document struct {
		Document interface{} `json:"document"`
	}
	documents struct {
		Documents interface{} `json:"documents"`
	}
	insertedId   struct{ InsertedId string }
	insertedIds  struct{ InsertedIds []string }
	deletedCount struct{ DeletedCount int64 }
)

type (
	findOneOpt struct {
		Filter     interface{} `json:"filter,omitempty"`
		Projection interface{} `json:"projection,omitempty"`
	}
	FindOneOpt struct {
		Projection interface{}
	}

	findOpt struct {
		Filter     interface{} `json:"filter,omitempty"`
		Projection interface{} `json:"projection,omitempty"`
		Sort       interface{} `json:"sort,omitempty"`
		Limit      int64       `json:"limit,omitempty"`
		Skip       int64       `json:"skip,omitempty"`
	}
	FindOpt struct {
		Projection interface{}
		Sort       interface{}
		Limit      int64
		Skip       int64
	}

	updateOpt struct {
		Filter interface{} `json:"filter,omitempty"`
		Update interface{} `json:"update,omitempty"`
		Upsert bool        `json:"upsert,omitempty"`
	}
	UpdateOpt struct {
		Upsert bool
	}

	replaceOneOpt struct {
		Filter      interface{} `json:"filter,omitempty"`
		Replacement interface{} `json:"replacement,omitempty"`
		Upsert      bool        `json:"upsert,omitempty"`
	}

	deleteOpt struct {
		Filter interface{} `json:"filter,omitempty"`
	}

	aggregateOpt struct {
		Pipeline interface{} `json:"pipeline,omitempty"`
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
