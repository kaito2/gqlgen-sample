package graph

import "github.com/kaito2/gqlgen-sample/internal/adapter/db"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB *db.DB
}
