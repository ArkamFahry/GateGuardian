package resolvers

import (
	"context"
	"fmt"

	"github.com/ArkamFahry/GateGuardian/server/constants"
	"github.com/ArkamFahry/GateGuardian/server/crypto"
	"github.com/ArkamFahry/GateGuardian/server/db/memorydb"
	"github.com/ArkamFahry/GateGuardian/server/graph/model"
	"github.com/ArkamFahry/GateGuardian/server/validators"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func UpdateEnvResolver(ctx context.Context, params model.UpdateEnvInput) (*model.Response, error) {
	var res *model.Response

	if params.JwtType == nil && params.JwtSecret == nil && params.ClientID == nil {
		logrus.Debug("No params to update")
		return res, fmt.Errorf("please enter at least one param to update")
	}

	isJwtUpdated := false

	if params.JwtType != nil {
		algo := *params.JwtType
		if !crypto.IsHMACA(algo) && !crypto.IsECDSA(algo) && !crypto.IsRSA(algo) {
			logrus.Debug("Invalid JWT type: ", algo)
			return res, fmt.Errorf("invalid jwt type")
		}

		memorydb.Provider.UpdateEnv(constants.JwtType, algo)
	}

	if params.JwtSecret != nil || params.JwtType != nil || params.ClientID != nil {
		isJwtUpdated = true
	}

	// update jwt secret
	if params.JwtSecret != nil {
		if err := validators.IsValidJwtSecret(*params.JwtSecret); err != nil {
			logrus.Debug("Invalid Jwt Secret")
			return res, fmt.Errorf("jwt secret is not valid. It needs to be at least 32 characters long and needs contain number, uppercase letter, lowercase letter and special character")
		} else {
			memorydb.Provider.UpdateEnv(constants.JwtSecret, *params.JwtSecret)
		}
	}

	if params.ClientID != nil {
		if *params.ClientID == "" {
			*params.ClientID = uuid.New().String()
		}
		memorydb.Provider.UpdateEnv(constants.ClientID, *params.ClientID)
		isJwtUpdated = true
	}

	if isJwtUpdated {
		algo, _ := memorydb.Provider.GetEnvByKey(constants.JwtType)
		clientId, _ := memorydb.Provider.GetEnvByKey(constants.ClientID)
		if crypto.IsRSA(algo) {
			_, jwtPrivateKey, jwtPublicKey, _, _ := crypto.NewRSAKey(algo, clientId)
			memorydb.Provider.UpdateEnv(constants.JwtPrivateKey, jwtPrivateKey)
			memorydb.Provider.UpdateEnv(constants.JwtPrivateKey, jwtPublicKey)
		} else if crypto.IsECDSA(algo) {
			_, jwtPrivateKey, jwtPublicKey, _, _ := crypto.NewECDSAKey(algo, clientId)
			memorydb.Provider.UpdateEnv(constants.JwtPrivateKey, jwtPrivateKey)
			memorydb.Provider.UpdateEnv(constants.JwtPrivateKey, jwtPublicKey)
		}
	}

	res = &model.Response{
		Message: "configurations updated successfully",
	}

	return res, nil
}
