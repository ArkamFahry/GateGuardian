package tokens

import (
	"errors"

	"github.com/ArkamFahry/GateGuardian/server/constants"
	"github.com/ArkamFahry/GateGuardian/server/crypto"
	"github.com/ArkamFahry/GateGuardian/server/memorystore/envstore"
	"github.com/golang-jwt/jwt"
)

func SignJWTToken(claims jwt.MapClaims) (string, error) {
	jwtType, err := envstore.Provider.GetEnv(constants.JwtType)
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
		jwtSecret, err := envstore.Provider.GetEnv(constants.JwtSecret)
		if err != nil {
			return "", err
		}
		return t.SignedString([]byte(jwtSecret))
	case jwt.SigningMethodRS256, jwt.SigningMethodRS384, jwt.SigningMethodRS512:
		jwtPrivateKey, err := envstore.Provider.GetEnv(constants.JwtPrivateKey)
		if err != nil {
			return "", err
		}
		key, err := crypto.ParseRsaPrivateKeyFromPemStr(jwtPrivateKey)
		if err != nil {
			return "", err
		}
		return t.SignedString(key)
	case jwt.SigningMethodES256, jwt.SigningMethodES384, jwt.SigningMethodES512:
		jwtPrivateKey, err := envstore.Provider.GetEnv(constants.JwtPrivateKey)
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

// ParseJWTToken common util to parse jwt token
func ParseJWTToken(token string) (jwt.MapClaims, error) {
	jwtType, err := envstore.Provider.GetEnv(constants.JwtType)
	if err != nil {
		return nil, err
	}
	signingMethod := jwt.GetSigningMethod(jwtType)

	var claims jwt.MapClaims

	switch signingMethod {
	case jwt.SigningMethodHS256, jwt.SigningMethodHS384, jwt.SigningMethodHS512:
		_, err = jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
			jwtSecret, err := envstore.Provider.GetEnv(constants.JwtSecret)
			if err != nil {
				return nil, err
			}
			return []byte(jwtSecret), nil
		})
	case jwt.SigningMethodRS256, jwt.SigningMethodRS384, jwt.SigningMethodRS512:
		_, err = jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
			jwtPublicKey, err := envstore.Provider.GetEnv(constants.JwtPublicKey)
			if err != nil {
				return nil, err
			}
			key, err := crypto.ParseRsaPublicKeyFromPemStr(jwtPublicKey)
			if err != nil {
				return nil, err
			}
			return key, nil
		})
	case jwt.SigningMethodES256, jwt.SigningMethodES384, jwt.SigningMethodES512:
		_, err = jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
			jwtPublicKey, err := envstore.Provider.GetEnv(constants.JwtPublicKey)
			if err != nil {
				return nil, err
			}
			key, err := crypto.ParseEcdsaPublicKeyFromPemStr(jwtPublicKey)
			if err != nil {
				return nil, err
			}
			return key, nil
		})
	default:
		err = errors.New("unsupported signing method")
	}
	if err != nil {
		return claims, err
	}

	// claim parses exp & iat into float 64 with e^10 but we expect int64 hence we need to assert interface and convert to int64
	intExp := int64(claims["exp"].(float64))
	intIat := int64(claims["iat"].(float64))
	claims["exp"] = intExp
	claims["iat"] = intIat

	return claims, nil
}
