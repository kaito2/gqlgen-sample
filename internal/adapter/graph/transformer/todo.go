package transformer

import (
	"github.com/kaito2/gqlgen-sample/internal/adapter/db"
	generated "github.com/kaito2/gqlgen-sample/internal/adapter/graph/model"
)

func TodosFromRecords(recs []*db.Todo) []*generated.Todo {
	todos := make([]*generated.Todo, len(recs))
	for i, rec := range recs {
		todos[i] = TodoFromRecord(rec)
	}
	return todos
}

func TodoFromRecord(rec *db.Todo) *generated.Todo {
	return &generated.Todo{
		ID:     rec.ID,
		Text:   rec.Text,
		Done:   rec.Done,
		UserID: rec.UserID,
		User:   nil,
	}
}
