package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/ArkamFahry/GateGuardian/server/graph/generated"
	"github.com/ArkamFahry/GateGuardian/server/graph/model"
	"github.com/ArkamFahry/GateGuardian/server/resolvers"
)

// Env is the resolver for the _env field.
func (r *queryResolver) Env(ctx context.Context) (*model.Env, error) {
	return resolvers.EnvResolver(ctx)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
