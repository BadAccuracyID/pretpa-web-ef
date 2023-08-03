package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.35

import (
	"context"

	"github.com/badaccuracyid/pretpa-web-ef/internal/graph"
	"github.com/badaccuracyid/pretpa-web-ef/internal/graph/model"
	"github.com/badaccuracyid/pretpa-web-ef/internal/service"
)

// UpdateCurrentUser is the resolver for the updateCurrentUser field.
func (r *mutationResolver) UpdateCurrentUser(ctx context.Context, input model.UserInput) (*model.User, error) {
	userService := service.NewUserService(ctx, r.DB)
	return userService.UpdateCurrentUser(&input)
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input model.UserInput) (*model.User, error) {
	userService := service.NewUserService(ctx, r.DB)
	return userService.UpdateUser(id, &input)
}

// DeleteCurrentUser is the resolver for the deleteCurrentUser field.
func (r *mutationResolver) DeleteCurrentUser(ctx context.Context) (*model.User, error) {
	userService := service.NewUserService(ctx, r.DB)
	return userService.DeleteCurrentUser()
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*model.User, error) {
	userService := service.NewUserService(ctx, r.DB)
	return userService.DeleteUser(id)
}

// GetCurrentUser is the resolver for the getCurrentUser field.
func (r *queryResolver) GetCurrentUser(ctx context.Context) (*model.User, error) {
	userService := service.NewUserService(ctx, r.DB)
	return userService.GetCurrentUser()
}

// GetUser is the resolver for the getUser field.
func (r *queryResolver) GetUser(ctx context.Context, id string) (*model.User, error) {
	userService := service.NewUserService(ctx, r.DB)
	return userService.GetUser(id)
}

// GetAllUsers is the resolver for the getAllUsers field.
func (r *queryResolver) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	userService := service.NewUserService(ctx, r.DB)
	return userService.GetAllUsers()
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
