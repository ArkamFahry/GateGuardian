package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/ArkamFahry/GateGuardian/server/graph/generated"
	"github.com/ArkamFahry/GateGuardian/server/graph/model"
)

// Env is the resolver for the _env field.
func (r *queryResolver) Env(ctx context.Context) (*model.Env, error) {
	panic(fmt.Errorf("not implemented: Env - _env"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
