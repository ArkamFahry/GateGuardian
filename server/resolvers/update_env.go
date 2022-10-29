package resolvers

import (
	"context"
	"fmt"

	"github.com/ArkamFahry/GateGuardian/server/constants"
	"github.com/ArkamFahry/GateGuardian/server/db/memorydb"
	"github.com/ArkamFahry/GateGuardian/server/graph/model"
	"github.com/sirupsen/logrus"
)

func UpdateEnvResolver(ctx context.Context, params model.UpdateEnvInput) (*model.Response, error) {
	var res *model.Response

	if params.JwtType == nil && params.JwtSecret == nil && params.ClientID == nil {
		logrus.Debug("No params to update")
		return res, fmt.Errorf("please enter at least one param to update")
	}

	if params.JwtType != nil {
		memorydb.Provider.UpdateEnv(constants.JwtType, *params.JwtType)
	}

	if params.JwtSecret != nil {
		memorydb.Provider.UpdateEnv(constants.JwtSecret, *params.JwtSecret)
	}

	if params.ClientID != nil {
		memorydb.Provider.UpdateEnv(constants.ClientID, *params.ClientID)
	}

	res = &model.Response{
		Message: "configurations updated successfully",
	}

	return res, nil
}
