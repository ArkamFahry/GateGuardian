package resolvers

import (
	"context"
	"fmt"

	"github.com/ArkamFahry/GateGuardian/server/constants"
	"github.com/ArkamFahry/GateGuardian/server/crypto"
	"github.com/ArkamFahry/GateGuardian/server/db/memorydb"
	"github.com/ArkamFahry/GateGuardian/server/graph/model"
	"github.com/ArkamFahry/GateGuardian/server/validators"
	"github.com/sirupsen/logrus"
)

func UpdateEnvResolver(ctx context.Context, params model.UpdateEnvInput) (*model.Response, error) {
	var res *model.Response

	if params.JwtType == nil && params.JwtSecret == nil && params.ClientID == nil {
		logrus.Debug("No params to update")
		return res, fmt.Errorf("please enter at least one param to update")
	}

	isJWTUpdated := false

	// update jwt type
	algo, _ := memorydb.Provider.GetEnvByKey(constants.JwtType)
	if params.JwtType != nil {
		jwtAlgo := *params.JwtType
		if !crypto.IsHMACA(jwtAlgo) && !crypto.IsECDSA(jwtAlgo) && !crypto.IsRSA(jwtAlgo) {
			logrus.Debug("Invalid JWT type: ", jwtAlgo)
			return res, fmt.Errorf("invalid jwt type")
		}

		memorydb.Provider.UpdateEnv(constants.JwtType, jwtAlgo)
		algo, _ = memorydb.Provider.GetEnvByKey(constants.JwtType)
	}

	// update jwt secret
	if err := validators.IsValidJwtSecret(*params.JwtSecret); err != nil {
		logrus.Debug("Invalid Jwt Secret")
		return res, err
	} else {
		memorydb.Provider.UpdateEnv(constants.JwtSecret, *params.JwtSecret)
	}

	// use to reset when type is changed from rsa, edsa -> hmac or vice a versa
	if isJWTUpdated {
		defaultSecret := ""
		defaultPublicKey := ""
		defaultPrivateKey := ""
		// check if jwt secret is provided
		if crypto.IsHMACA(algo) {
			if params.JwtSecret == nil {
				logrus.Debug("JWT secret is required for HMAC")
				return res, fmt.Errorf("jwt secret is required for HMAC algorithm")
			}

			// reset public key and private key
			params.JwtPrivateKey = &defaultPrivateKey
			params.JwtPublicKey = &defaultPublicKey
		}

		if crypto.IsRSA(algo) {
			if params.JwtPrivateKey == nil || params.JwtPublicKey == nil {
				logrus.Debug("JWT private key and public key are required for RSA: ", *params.JwtPrivateKey, *params.JwtPublicKey)
				return res, fmt.Errorf("jwt private and public key is required for RSA (PKCS1) / ECDSA algorithm")
			}

			// reset the jwt secret
			params.JwtSecret = &defaultSecret
			_, err := crypto.ParseRsaPrivateKeyFromPemStr(*params.JwtPrivateKey)
			if err != nil {
				logrus.Debug("Invalid JWT private key: ", err)
				return res, err
			} else {
				memorydb.Provider.UpdateEnv(constants.JwtPrivateKey, *params.JwtPrivateKey)
			}

			_, err = crypto.ParseRsaPublicKeyFromPemStr(*params.JwtPublicKey)
			if err != nil {
				logrus.Debug("Invalid JWT public key: ", err)
				return res, err
			} else {
				memorydb.Provider.UpdateEnv(constants.JwtPublicKey, *params.JwtPublicKey)
			}
		}

		if crypto.IsECDSA(algo) {
			if params.JwtPrivateKey == nil || params.JwtPublicKey == nil {
				logrus.Debug("JWT private key and public key are required for ECDSA: ", *params.JwtPrivateKey, *params.JwtPublicKey)
				return res, fmt.Errorf("jwt private and public key is required for RSA (PKCS1) / ECDSA algorithm")
			}

			// reset the jwt secret
			params.JwtSecret = &defaultSecret
			_, err := crypto.ParseEcdsaPrivateKeyFromPemStr(*params.JwtPrivateKey)
			if err != nil {
				logrus.Debug("Invalid JWT private key: ", err)
				return res, err
			} else {
				memorydb.Provider.UpdateEnv(constants.JwtPrivateKey, *params.JwtPrivateKey)
			}

			_, err = crypto.ParseEcdsaPublicKeyFromPemStr(*params.JwtPublicKey)
			if err != nil {
				logrus.Debug("Invalid JWT public key: ", err)
				return res, err
			} else {
				memorydb.Provider.UpdateEnv(constants.JwtPublicKey, *params.JwtPublicKey)
			}
		}

	}

	res = &model.Response{
		Message: "configurations updated successfully",
	}

	return res, nil
}
