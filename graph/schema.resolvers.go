package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.55

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/nozoo3/go-template/cmd/hello"
	"github.com/nozoo3/go-template/graph/model"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	randNumber, _ := rand.Int(rand.Reader, big.NewInt(100))
	todo := &model.Todo{
		Text: input.Text,
		ID:   fmt.Sprintf("T%d", randNumber),
		User: &model.User{ID: input.UserID, Name: "user " + input.UserID},
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

// CreateGreeting is the resolver for the createGreeting field.
func (r *mutationResolver) CreateGreeting(ctx context.Context, input model.NewGreeting) (*model.Greeting, error) {
	message := hello.Hello()
	greeting := &model.Greeting{
		Message: message + input.Message,
	}
	r.greetings = append(r.greetings, greeting)
	return greeting, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

// Greetings is the resolver for the greetings field.
func (r *queryResolver) Greetings(ctx context.Context) ([]*model.Greeting, error) {
	return r.greetings, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
