package dataloader

// import graph gophers with your other imports
import (
	"context"
	"fmt"
	dl "github.com/graph-gophers/dataloader"
	"github.com/kaito2/gqlgen-sample/internal/adapter/db"
	generated "github.com/kaito2/gqlgen-sample/internal/adapter/graph/model"
	"log"
	"net/http"
)

type ctxKey string

const (
	loadersKey = ctxKey("dataloaders")
)

// UserReader reads Users from a database
type UserReader struct {
	db *db.DB
}

// GetUsers implements a batch function that can retrieve many users by ID,
// for use in a dl
func (u *UserReader) GetUsers(ctx context.Context, keys dl.Keys) []*dl.Result {
	// read all requested users in a single query
	userIDs := make([]string, len(keys))
	for ix, key := range keys {
		userIDs[ix] = key.String()
	}

	users, err := u.db.FindUsersByIDs(userIDs)
	if err != nil {
		// REVIEW: エラー処理どうする?
		log.Println(fmt.Errorf("failed to db.FindUsersByIDs: %w", err))
		return nil
	}

	// return User records into a map by ID
	userById := map[string]*generated.User{}
	for _, user := range users {
		userById[user.ID] = &generated.User{
			ID:   user.ID,
			Name: user.Name,
		}
	}

	// return users in the same order requested
	output := make([]*dl.Result, len(keys))
	for i, userKey := range keys {
		record, ok := userById[userKey.String()]
		if ok {
			output[i] = &dl.Result{Data: record, Error: nil}
		} else {
			err := fmt.Errorf("user not found (id: %s)", userKey.String())
			output[i] = &dl.Result{Data: nil, Error: err}
		}
	}
	return output
}

// Loaders wrap your data loaders to inject via middleware
type Loaders struct {
	UserLoader *dl.Loader
}

// NewLoaders instantiates data loaders for the middleware
func NewLoaders(db *db.DB) *Loaders {
	// define the data loader
	userReader := &UserReader{db: db}
	loaders := &Loaders{
		UserLoader: dl.NewBatchedLoader(userReader.GetUsers),
	}
	return loaders
}

// Middleware injects data loaders into the context
func Middleware(loaders *Loaders, next http.Handler) http.Handler {
	// return a middleware that injects the loader to the request context
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextCtx := context.WithValue(r.Context(), loadersKey, loaders)
		r = r.WithContext(nextCtx)
		next.ServeHTTP(w, r)
	})
}

// For returns the dl for a given context
func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}

// GetUser wraps the User dl for efficient retrieval by user ID
func GetUser(ctx context.Context, userID string) (*generated.User, error) {
	loaders := For(ctx)
	thunk := loaders.UserLoader.Load(ctx, dl.StringKey(userID))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result.(*generated.User), nil
}
