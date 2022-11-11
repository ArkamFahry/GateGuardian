package tokens

import (
	"errors"

	"github.com/ArkamFahry/GateGuardian/server/constants"
	"github.com/ArkamFahry/GateGuardian/server/crypto"
	"github.com/ArkamFahry/GateGuardian/server/memorystore/envstore"
	"github.com/golang-jwt/jwt"
)

func SignJWTToken(claims jwt.MapClaims) (string, error) {
	jwtType, err := envstore.Provider.GetEnv(constants.JWT_TYPE)
	if err != nil {
		return "", err
	}
	signingMethod := jwt.GetSigningMethod(jwtType)
	if signingMethod == nil {
		return "", errors.New("unsupported signing method")
	}
	t := jwt.New(signingMethod)
	if t == nil {
		return "", errors.New("unsupported signing method")
	}
	t.Claims = claims

	switch signingMethod {
	case jwt.SigningMethodHS256, jwt.SigningMethodHS384, jwt.SigningMethodHS512:
		jwtSecret, err := envstore.Provider.GetEnv(constants.JWT_SECRET)
		if err != nil {
			return "", err
		}
		return t.SignedString([]byte(jwtSecret))
	case jwt.SigningMethodRS256, jwt.SigningMethodRS384, jwt.SigningMethodRS512:
		jwtPrivateKey, err := envstore.Provider.GetEnv(constants.JWT_PRIVATE_KEY)
		if err != nil {
			return "", err
		}
		key, err := crypto.ParseRsaPrivateKeyFromPemStr(jwtPrivateKey)
		if err != nil {
			return "", err
		}
		return t.SignedString(key)
	case jwt.SigningMethodES256, jwt.SigningMethodES384, jwt.SigningMethodES512:
		jwtPrivateKey, err := envstore.Provider.GetEnv(constants.JWT_PRIVATE_KEY)
		if err != nil {
			return "", err
		}
		key, err := crypto.ParseEcdsaPrivateKeyFromPemStr(jwtPrivateKey)
		if err != nil {
			return "", err
		}

		return t.SignedString(key)
	default:
		return "", errors.New("unsupported signing method")
	}
}
