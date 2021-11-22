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
	deletedCount struct{ DeletedCount int }
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
		Limit      int         `json:"limit,omitempty"`
		Skip       int         `json:"skip,omitempty"`
	}
	FindOpt struct {
		Projection interface{}
		Sort       interface{}
		Limit      int
		Skip       int
	}

	updateOpt struct {
		Filter interface{} `json:"filter,omitempty"`
		Update interface{} `json:"update,omitempty"`
		Upsert bool        `json:"upsert,omitempty"`
	}
	UpdateOpt struct {
		Update interface{}
		Upsert bool
	}

	replaceOneOpt struct {
		Filter      interface{} `json:"filter,omitempty"`
		Replacement interface{} `json:"replacement,omitempty"`
		Upsert      bool        `json:"upsert,omitempty"`
	}
	ReplaceOneOpt struct {
		Replacement interface{}
		Upsert      bool
	}

	deleteOpt struct {
		Filter interface{} `json:"filter,omitempty"`
	}

	aggregateOpt struct {
		Pipeline interface{} `json:"pipeline,omitempty"`
	}
)

type Result struct {
	MatchedCount  int
	ModifiedCount int
	UpsertedId    string
}
