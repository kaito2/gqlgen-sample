package transformer

import (
	"github.com/kaito2/gqlgen-sample/internal/adapter/db"
	generated "github.com/kaito2/gqlgen-sample/internal/adapter/graph/model"
)

func UsersFromRecords(recs []*db.User) []*generated.User {
	users := make([]*generated.User, len(recs))
	for i, rec := range recs {
		users[i] = UserFromRecord(rec)
	}
	return users
}

func UserFromRecord(rec *db.User) *generated.User {
	return &generated.User{
		ID:   rec.ID,
		Name: rec.Name,
	}
}
