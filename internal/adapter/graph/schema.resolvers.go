package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/kaito2/gqlgen-sample/internal/adapter/graph/dataloader"
	"github.com/kaito2/gqlgen-sample/internal/adapter/graph/generated"
	"github.com/kaito2/gqlgen-sample/internal/adapter/graph/transformer"

	generated_model "github.com/kaito2/gqlgen-sample/internal/adapter/graph/model"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input generated_model.NewTodo) (*generated_model.Todo, error) {
	panic(fmt.Errorf("not implemented: CreateTodo - createTodo"))
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*generated_model.Todo, error) {
	return transformer.TodosFromRecords(r.DB.GetTodos()), nil
}

// User is the resolver for the user field.
func (r *todoResolver) User(ctx context.Context, obj *generated_model.Todo) (*generated_model.User, error) {
	return dataloader.GetUser(ctx, obj.UserID)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
