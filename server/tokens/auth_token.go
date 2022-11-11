package tokens

import (
	"time"

	"github.com/ArkamFahry/GateGuardian/server/constants"
	"github.com/ArkamFahry/GateGuardian/server/db/models"
	"github.com/ArkamFahry/GateGuardian/server/memorystore/envstore"
	"github.com/ArkamFahry/GateGuardian/server/utils"
	"github.com/golang-jwt/jwt"
)

func CreateAccessToken(user models.User, roles, scopes []string, hostName, nonce, loginMethod string) (string, int64, error) {
	expireTime, err := envstore.Provider.GetEnv(constants.ACCESS_TOKEN_EXPIRY_TIME)
	if err != nil {
		return "", 0, err
	}
	expiryBound, err := utils.ParseDurationInSeconds(expireTime)
	if err != nil {
		expiryBound = time.Minute * 30
	}

	expiresAt := time.Now().Add(expiryBound).Unix()

	clientID, err := envstore.Provider.GetEnv(constants.CLIENT_ID)
	if err != nil {
		return "", 0, err
	}
	customClaims := jwt.MapClaims{
		"iss":          hostName,
		"aud":          clientID,
		"nonce":        nonce,
		"sub":          user.Id,
		"exp":          expiresAt,
		"iat":          time.Now().Unix(),
		"token_type":   constants.AccessToken,
		"scope":        scopes,
		"roles":        roles,
		"login_method": loginMethod,
	}

	token, err := SignJWTToken(customClaims)
	if err != nil {
		return "", 0, err
	}

	return token, expiresAt, nil
}

func CreateRefreshToken(user models.User, roles, scopes []string, hostname, nonce, loginMethod string) (string, int64, error) {
	expiryBound := time.Hour * 8760
	expiresAt := time.Now().Add(expiryBound).Unix()
	clientID, err := envstore.Provider.GetEnv(constants.CLIENT_ID)
	if err != nil {
		return "", 0, err
	}
	customClaims := jwt.MapClaims{
		"iss":          hostname,
		"aud":          clientID,
		"sub":          user.Id,
		"exp":          expiresAt,
		"iat":          time.Now().Unix(),
		"token_type":   constants.RefreshToken,
		"roles":        roles,
		"scope":        scopes,
		"nonce":        nonce,
		"login_method": loginMethod,
	}

	token, err := SignJWTToken(customClaims)
	if err != nil {
		return "", 0, err
	}

	return token, expiresAt, nil
}
