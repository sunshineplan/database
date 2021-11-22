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

type Result struct {
	MatchedCount  int
	ModifiedCount int
	UpsertedId    string
}
