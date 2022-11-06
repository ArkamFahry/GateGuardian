package token

import (
	"encoding/json"
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
	"github.com/robertkrimen/otto"
	log "github.com/sirupsen/logrus"
)

type JwtToken struct {
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expires_at"`
}

type Token struct {
	AccessToken  *JwtToken `json:"access_token"`
	RefreshToken *JwtToken `json:"refresh_token"`
	IdToken      *JwtToken `json:"id_token"`
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
	idToken, idTokenExpiresAt, err := CreateIdToken(user, roles, hostName, nonce, loginMethod)
	if err != nil {
		return nil, err
	}

	res := &Token{
		AccessToken:  &JwtToken{Token: accessToken, ExpiresAt: accessTokenExpiresAt},
		RefreshToken: &JwtToken{Token: refreshToken, ExpiresAt: refreshTokenExpiresAt},
		IdToken:      &JwtToken{Token: idToken, ExpiresAt: idTokenExpiresAt},
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

func CreateIdToken(user models.User, roles []string, hostname, nonce, loginMethod string) (string, int64, error) {
	expireTime, err := env.GetEnvByKey(constants.AccessTokenExpiryTime)
	if err != nil {
		return "", 0, err
	}
	expiryBound, err := utils.ParseDurationInSeconds(expireTime)
	if err != nil {
		expiryBound = time.Minute * 30
	}

	expiresAt := time.Now().Add(expiryBound).Unix()

	resUser := user.AsAPIUser()
	userBytes, _ := json.Marshal(&resUser)
	var userMap map[string]interface{}
	json.Unmarshal(userBytes, &userMap)

	claimKey, err := env.GetEnvByKey(constants.JwtRoleClaim)
	if err != nil {
		claimKey = "roles"
	}

	clientID, err := env.GetEnvByKey(constants.ClientID)
	if err != nil {
		return "", 0, err
	}
	customClaims := jwt.MapClaims{
		"iss":           hostname,
		"aud":           clientID,
		"nonce":         nonce,
		"sub":           user.ID,
		"exp":           expiresAt,
		"iat":           time.Now().Unix(),
		"token_type":    constants.TokenTypeIdentityToken,
		"allowed_roles": strings.Split(user.Roles, ","),
		"login_method":  loginMethod,
		claimKey:        roles,
	}

	for k, v := range userMap {
		if k != "roles" {
			customClaims[k] = v
		}
	}

	accessTokenScript, err := env.GetEnvByKey(constants.CustomAccessTokenScript)
	if err != nil {
		log.Debug("Failed to get custom access token script: ", err)
		accessTokenScript = ""
	}
	if accessTokenScript != "" {
		vm := otto.New()

		claimBytes, _ := json.Marshal(customClaims)
		vm.Run(fmt.Sprintf(`
			var user = %s;
			var tokenPayload = %s;
			var customFunction = %s;
			var functionRes = JSON.stringify(customFunction(user, tokenPayload));
		`, string(userBytes), string(claimBytes), accessTokenScript))

		val, err := vm.Get("functionRes")
		if err != nil {
			log.Debug("error getting custom access token script: ", err)
		} else {
			extraPayload := make(map[string]interface{})
			err = json.Unmarshal([]byte(fmt.Sprintf("%s", val)), &extraPayload)
			if err != nil {
				log.Debug("error converting accessTokenScript response to map: ", err)
			} else {
				for k, v := range extraPayload {
					customClaims[k] = v
				}
			}
		}
	}

	token, err := SignJWTToken(customClaims)
	if err != nil {
		return "", 0, err
	}

	return token, expiresAt, nil
}

func GetIDToken(gc *gin.Context) (string, error) {
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

	token, err := memorydb.Provider.GetSession(gc, sessionKey)
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

	token, err := memorydb.Provider.GetSession(gc, sessionKey)
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
