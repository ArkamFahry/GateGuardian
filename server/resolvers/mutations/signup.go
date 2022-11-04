package mutations

import (
	"context"

	"github.com/ArkamFahry/GateGuardian/server/graph/model"
)

func SignupResolver(ctx context.Context, params model.SignUpInput) (*model.AuthResponse, error)
