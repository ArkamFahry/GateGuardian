package tokens

import (
	"strings"
	"time"

	"github.com/ArkamFahry/GateGuardian/server/api/model"
	"github.com/ArkamFahry/GateGuardian/server/constants"
	"github.com/ArkamFahry/GateGuardian/server/db/models"
	"github.com/ArkamFahry/GateGuardian/server/memorystore/envstore"
	"github.com/ArkamFahry/GateGuardian/server/utils"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func CreateAuthTokens(user models.User, hostName string) (model.AuthResponse, error) {
	nonce := uuid.New().String()
	access_token, access_token_expiry, err := CreateAccessToken(user, hostName, nonce)
	if err != nil {
		return model.AuthResponse{}, err
	}
	refresh_token, refresh_token_expiry, err := CreateRefreshToken(user, hostName, nonce)
	if err != nil {
		return model.AuthResponse{}, err
	}

	return model.AuthResponse{
		AccessToken:         access_token,
		AccessTokenExpires:  access_token_expiry,
		RefreshToken:        refresh_token,
		RefreshTokenExpires: refresh_token_expiry,
	}, nil
}

func CreateAccessToken(user models.User, hostName, nonce string) (string, int64, error) {
	expireTime, err := envstore.Provider.GetEnv(constants.AccessTokenExpiryTime)
	if err != nil {
		return "", 0, err
	}
	expiryBound, err := utils.ParseDurationInSeconds(expireTime)
	if err != nil {
		expiryBound = time.Minute * 30
	}

	expiresAt := time.Now().Add(expiryBound).Unix()

	AllowedRoles, _ := envstore.Provider.GetEnv(constants.AllowedRoles)
	allowedRoles := strings.Split(AllowedRoles, ",")

	roles := strings.Split(user.Roles, ",")

	clientID, err := envstore.Provider.GetEnv(constants.ClientId)
	if err != nil {
		return "", 0, err
	}
	customClaims := jwt.MapClaims{
		"iss":        hostName,
		"aud":        clientID,
		"nonce":      nonce,
		"sub":        user.Id,
		"exp":        expiresAt,
		"iat":        time.Now().Unix(),
		"token_type": constants.AccessToken,
		"https://hasura.io/jwt/claims": map[string]interface{}{
			"x-hasura-allowed-roles": allowedRoles,
			"x-hasura-default-role":  roles,
			"x-hasura-user-id":       user.Id,
			"x-hasura-user-email":    user.Email,
		},
	}

	token, err := SignJWTToken(customClaims)
	if err != nil {
		return "", 0, err
	}

	return token, expiresAt, nil
}

func CreateRefreshToken(user models.User, hostname, nonce string) (string, int64, error) {
	expiryBound := time.Hour * 8760
	expiresAt := time.Now().Add(expiryBound).Unix()
	clientID, err := envstore.Provider.GetEnv(constants.ClientId)
	if err != nil {
		return "", 0, err
	}
	customClaims := jwt.MapClaims{
		"iss":        hostname,
		"aud":        clientID,
		"sub":        user.Id,
		"exp":        expiresAt,
		"iat":        time.Now().Unix(),
		"token_type": constants.RefreshToken,
		"nonce":      nonce,
	}

	token, err := SignJWTToken(customClaims)
	if err != nil {
		return "", 0, err
	}

	return token, expiresAt, nil
}
