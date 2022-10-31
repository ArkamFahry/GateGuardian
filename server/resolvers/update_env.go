package resolvers

import (
	"context"
	"fmt"
	"strings"

	"github.com/ArkamFahry/GateGuardian/server/constants"
	"github.com/ArkamFahry/GateGuardian/server/crypto"
	"github.com/ArkamFahry/GateGuardian/server/graph/model"
	"github.com/ArkamFahry/GateGuardian/server/memorydb"
	"github.com/ArkamFahry/GateGuardian/server/validators"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func UpdateEnvResolver(ctx context.Context, params model.UpdateEnvInput) (*model.Response, error) {
	var res *model.Response

	if params.JwtType == nil && params.JwtSecret == nil && params.ClientID == nil && params.Roles == nil && params.DefaultRoles == nil {
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

		algo, _ = crypto.EncryptAES(*params.JwtType)
		memorydb.Provider.UpdateEnv(constants.JwtType, algo)
	}

	if params.JwtSecret != nil || params.JwtType != nil || params.ClientID != nil {
		isJwtUpdated = true
	}

	if params.JwtSecret != nil {
		if err := validators.IsValidJwtSecret(*params.JwtSecret); err != nil {
			logrus.Debug("Invalid Jwt Secret")
			return res, fmt.Errorf("jwt secret is not valid. It needs to be at least 32 characters long and needs contain number, uppercase letter, lowercase letter and special character")
		} else {
			jwtSecret, _ := crypto.EncryptAES(*params.JwtSecret)
			memorydb.Provider.UpdateEnv(constants.JwtSecret, jwtSecret)
		}
	}

	if params.ClientID != nil {
		if *params.ClientID == "" {
			*params.ClientID = uuid.New().String()
		}

		clientID, _ := crypto.EncryptAES(*params.ClientID)
		memorydb.Provider.UpdateEnv(constants.ClientID, clientID)
		isJwtUpdated = true
	}

	if isJwtUpdated {
		algo, _ := memorydb.Provider.GetEnvByKey(constants.JwtType)
		clientId, _ := memorydb.Provider.GetEnvByKey(constants.ClientID)
		if crypto.IsRSA(algo) {
			_, jwtPrivateKey, jwtPublicKey, _, _ := crypto.NewRSAKey(algo, clientId)
			jwtPrivateKey, _ = crypto.EncryptAES(jwtPrivateKey)
			jwtPublicKey, _ = crypto.EncryptAES(jwtPublicKey)
			memorydb.Provider.UpdateEnv(constants.JwtPrivateKey, jwtPrivateKey)
			memorydb.Provider.UpdateEnv(constants.JwtPrivateKey, jwtPublicKey)
		} else if crypto.IsECDSA(algo) {
			_, jwtPrivateKey, jwtPublicKey, _, _ := crypto.NewECDSAKey(algo, clientId)
			jwtPrivateKey, _ = crypto.EncryptAES(jwtPrivateKey)
			jwtPublicKey, _ = crypto.EncryptAES(jwtPublicKey)
			memorydb.Provider.UpdateEnv(constants.JwtPrivateKey, jwtPrivateKey)
			memorydb.Provider.UpdateEnv(constants.JwtPrivateKey, jwtPublicKey)
		}
	}

	if params.Roles != nil {
		if len(params.Roles) > 0 {
			roles, _ := crypto.EncryptAES(strings.Join(params.Roles, ","))
			memorydb.Provider.UpdateEnv(constants.Roles, roles)
		}
	}

	if params.DefaultRoles != nil {
		if len(params.DefaultRoles) > 0 {
			roles, _ := memorydb.Provider.GetEnvByKey(constants.Roles)
			roles, _ = crypto.DecryptAES(roles)
			if !validators.IsValidRoles(params.DefaultRoles, strings.Split(roles, ",")) {
				return res, fmt.Errorf("invalid list of default roles")
			} else {
				defaultRoles, _ := crypto.EncryptAES(strings.Join(params.DefaultRoles, ","))
				memorydb.Provider.UpdateEnv(constants.DefaultRoles, defaultRoles)
			}
		}
	}

	res = &model.Response{
		Message: "configurations updated successfully",
	}

	return res, nil
}
