package mutations

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/ArkamFahry/GateGuardian/server/constants"
	"github.com/ArkamFahry/GateGuardian/server/crypto"
	"github.com/ArkamFahry/GateGuardian/server/db/maindb"
	"github.com/ArkamFahry/GateGuardian/server/db/memorydb"
	"github.com/ArkamFahry/GateGuardian/server/graph/model"
	"github.com/ArkamFahry/GateGuardian/server/token"
	"github.com/ArkamFahry/GateGuardian/server/utils"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func LoginResolver(ctx context.Context, params model.LoginInput) (*model.AuthResponse, error) {
	var res *model.AuthResponse

	gc, err := utils.GinContextFromContext(ctx)
	if err != nil {
		log.Debug("Failed to get GinContext: ", err)
		return res, err
	}

	params.Email = strings.ToLower(params.Email)
	user, err := maindb.Provider.GetUserByEmail(ctx, params.Email)
	if err != nil {
		log.Debug("Failed to get user by email: ", err)
		return res, fmt.Errorf(`bad user credentials`)
	}

	if user.RevokedTimestamp != nil {
		log.Debug("User access is revoked")
		return res, fmt.Errorf(`user access has been revoked`)
	}

	if !strings.Contains(user.SignUpMethods, constants.AuthRecipeMethodBasicAuth) {
		log.Debug("User signup method is not basic auth")
		return res, fmt.Errorf(`user has not signed up email & password`)
	}

	err = bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(params.Password))

	if err != nil {
		log.Debug("Failed to compare password: ", err)
		return res, fmt.Errorf(`bad user credentials`)
	}

	roles := strings.Split(user.Roles, ",")
	scope := []string{"email", "profile"}
	authToken, err := token.CreateAuthTokens(gc, user, roles, scope, constants.AuthRecipeMethodBasicAuth)
	if err != nil {
		log.Debug("Failed to create auth tokens: ", err)
		return res, err
	}

	expiresIn := authToken.AccessToken.ExpiresAt - time.Now().Unix()
	if expiresIn <= 0 {
		expiresIn = 1
	}

	sessionKey := constants.AuthRecipeMethodBasicAuth + ":" + user.ID
	refreshTokenHash, err := crypto.EncryptPassword(authToken.RefreshToken.Token)
	if err != nil {
		log.Debug("Failed to hash refresh tokens: ", err)
		return res, err
	}
	memorydb.Provider.SetSession(ctx, sessionKey, refreshTokenHash)

	res = &model.AuthResponse{
		Message:      `Logged In successfully.`,
		AccessToken:  &authToken.AccessToken.Token,
		ExpiresIn:    &expiresIn,
		RefreshToken: &authToken.RefreshToken.Token,
		IDToken:      &authToken.IdToken.Token,
		User:         user.AsAPIUser(),
	}

	return res, nil
}
