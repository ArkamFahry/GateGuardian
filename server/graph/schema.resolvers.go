package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/ArkamFahry/GateGuardian/server/graph/generated"
	"github.com/ArkamFahry/GateGuardian/server/graph/model"
	"github.com/ArkamFahry/GateGuardian/server/resolvers/mutations"
	"github.com/ArkamFahry/GateGuardian/server/resolvers/queries"
)

// Signup is the resolver for the signup field.
func (r *mutationResolver) Signup(ctx context.Context, params model.SignUpInput) (*model.AuthResponse, error) {
	return mutations.SignupResolver(ctx, params)
}

// UpdateEnv is the resolver for the _update_env field.
func (r *mutationResolver) UpdateEnv(ctx context.Context, params model.UpdateEnvInput) (*model.Response, error) {
	return mutations.UpdateEnvResolver(ctx, params)
}

// Profile is the resolver for the profile field.
func (r *queryResolver) Profile(ctx context.Context) (*model.User, error) {
	panic(fmt.Errorf("not implemented: Profile - profile"))
}

// Env is the resolver for the _env field.
func (r *queryResolver) Env(ctx context.Context) (*model.Env, error) {
	return queries.EnvResolver(ctx)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
