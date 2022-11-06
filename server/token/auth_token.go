package token

import (
	"fmt"
	"strings"
	"time"

	"github.com/ArkamFahry/GateGuardian/server/constants"
	"github.com/ArkamFahry/GateGuardian/server/db/maindb/models"
	"github.com/ArkamFahry/GateGuardian/server/db/memorydb"
	"github.com/ArkamFahry/GateGuardian/server/env"
	"github.com/ArkamFahry/GateGuardian/server/parsers"
	"github.com/ArkamFahry/GateGuardian/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type JwtToken struct {
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expires_at"`
}

type Token struct {
	AccessToken  *JwtToken `json:"access_token"`
	RefreshToken *JwtToken `json:"refresh_token"`
}

func CreateAuthTokens(gc *gin.Context, user models.User, roles, scopes []string, loginMethod string) (*Token, error) {
	hostName := parsers.GetHost(gc)
	nonce := uuid.New().String()

	accessToken, accessTokenExpiresAt, err := CreateAccessToken(user, roles, scopes, hostName, nonce, loginMethod)
	if err != nil {
		return nil, err
	}
	refreshToken, refreshTokenExpiresAt, err := CreateRefreshToken(user, roles, scopes, hostName, nonce, loginMethod)
	if err != nil {
		return nil, err
	}

	res := &Token{
		AccessToken:  &JwtToken{Token: accessToken, ExpiresAt: accessTokenExpiresAt},
		RefreshToken: &JwtToken{Token: refreshToken, ExpiresAt: refreshTokenExpiresAt},
	}

	return res, nil
}

func CreateAccessToken(user models.User, roles, scopes []string, hostName, nonce, loginMethod string) (string, int64, error) {
	expireTime, err := env.GetEnvByKey(constants.AccessTokenExpiryTime)
	if err != nil {
		return "", 0, err
	}
	expiryBound, err := utils.ParseDurationInSeconds(expireTime)
	if err != nil {
		expiryBound = time.Minute * 15
	}

	expiresAt := time.Now().Add(expiryBound).Unix()
	clientID, err := env.GetEnvByKey(constants.ClientID)
	if err != nil {
		return "", 0, err
	}
	customClaims := jwt.MapClaims{
		"iss":          hostName,
		"aud":          clientID,
		"nonce":        nonce,
		"sub":          user.ID,
		"exp":          expiresAt,
		"iat":          time.Now().Unix(),
		"token_type":   constants.TokenTypeAccessToken,
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

func CreateRefreshToken(user models.User, roles, scopes []string, hostName, nonce, loginMethod string) (string, int64, error) {
	expiryBound := time.Hour * 168
	expiresAt := time.Now().Add(expiryBound).Unix()
	clientID, err := env.GetEnvByKey(constants.ClientID)
	if err != nil {
		return "", 0, err
	}
	customClaims := jwt.MapClaims{
		"iss":          hostName,
		"aud":          clientID,
		"sub":          user.ID,
		"exp":          expiresAt,
		"iat":          time.Now().Unix(),
		"token_type":   constants.TokenTypeRefreshToken,
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

func GetAccessToken(gc *gin.Context) (string, error) {
	auth := gc.Request.Header.Get("Authorization")
	if auth == "" {
		return "", fmt.Errorf(`unauthorized`)
	}

	authSplit := strings.Split(auth, " ")
	if len(authSplit) != 2 {
		return "", fmt.Errorf(`unauthorized`)
	}

	if strings.ToLower(authSplit[0]) != "bearer" {
		return "", fmt.Errorf(`not a bearer token`)
	}

	token := strings.TrimPrefix(auth, "Bearer ")
	return token, nil
}

func ValidateAccessToken(gc *gin.Context, accessToken string) (map[string]interface{}, error) {
	res := make(map[string]interface{})

	if accessToken == "" {
		return res, fmt.Errorf(`unauthorized`)
	}

	res, err := ParseJWTToken(accessToken)
	if err != nil {
		return res, err
	}

	userID := res["sub"].(string)
	nonce := res["nonce"].(string)
	loginMethod := res["login_method"]
	sessionKey := userID
	if loginMethod != nil && loginMethod != "" {
		sessionKey = loginMethod.(string) + ":" + userID
	}

	token, err := memorydb.Provider.GetSession(sessionKey)
	if nonce == "" || err != nil {
		return res, fmt.Errorf(`unauthorized`)
	}

	if token != accessToken {
		return res, fmt.Errorf(`unauthorized`)
	}

	hostname := parsers.GetHost(gc)
	if ok, err := ValidateJWTClaims(res, hostname, nonce, userID); !ok || err != nil {
		return res, err
	}

	if res["token_type"] != constants.TokenTypeAccessToken {
		return res, fmt.Errorf(`unauthorized: invalid token type`)
	}

	return res, nil
}

func ValidateRefreshToken(gc *gin.Context, refreshToken string) (map[string]interface{}, error) {
	res := make(map[string]interface{})

	if refreshToken == "" {
		return res, fmt.Errorf(`unauthorized`)
	}

	res, err := ParseJWTToken(refreshToken)
	if err != nil {
		return res, err
	}

	userID := res["sub"].(string)
	nonce := res["nonce"].(string)
	loginMethod := res["login_method"]
	sessionKey := userID
	if loginMethod != nil && loginMethod != "" {
		sessionKey = loginMethod.(string) + ":" + userID
	}

	token, err := memorydb.Provider.GetSession(sessionKey)
	if nonce == "" || err != nil {
		return res, fmt.Errorf(`unauthorized`)
	}

	if token != refreshToken {
		return res, fmt.Errorf(`unauthorized`)
	}

	hostname := parsers.GetHost(gc)
	if ok, err := ValidateJWTClaims(res, hostname, nonce, userID); !ok || err != nil {
		return res, err
	}

	if res["token_type"] != constants.TokenTypeRefreshToken {
		return res, fmt.Errorf(`unauthorized: invalid token type`)
	}

	return res, nil
}
