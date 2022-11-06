package mutations

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/ArkamFahry/GateGuardian/server/constants"
	"github.com/ArkamFahry/GateGuardian/server/crypto"
	"github.com/ArkamFahry/GateGuardian/server/db/maindb"
	"github.com/ArkamFahry/GateGuardian/server/db/maindb/models"
	"github.com/ArkamFahry/GateGuardian/server/db/memorydb"
	"github.com/ArkamFahry/GateGuardian/server/env"
	"github.com/ArkamFahry/GateGuardian/server/graph/model"
	"github.com/ArkamFahry/GateGuardian/server/token"
	"github.com/ArkamFahry/GateGuardian/server/utils"
	"github.com/ArkamFahry/GateGuardian/server/validators"
	"github.com/sirupsen/logrus"
)

func SignupResolver(ctx context.Context, params model.SignUpInput) (*model.AuthResponse, error) {
	var res *model.AuthResponse

	gc, err := utils.GinContextFromContext(ctx)
	if err != nil {
		logrus.Debug("Failed to get GinContext: ", err)
		return res, err
	}

	if params.ConfirmPassword != params.Password {
		logrus.Debug("Passwords do not match")
		return res, fmt.Errorf(`password and confirm password does not match`)
	}

	if err := validators.IsValidPassword(params.Password); err != nil {
		logrus.Debug("Invalid password")
		return res, err
	}

	if !validators.IsValidEmail(params.Email) {
		logrus.Debug("Invalid email: ", params.Email)
		return res, fmt.Errorf(`invalid email address`)
	}

	params.Email = strings.ToLower(params.Email)

	existingUser, err := maindb.Provider.GetUserByEmail(ctx, params.Email)
	if err == nil {
		logrus.Debug("Failed to get user by email: ", err)
	}

	if existingUser.EmailVerifiedAt != nil {
		logrus.Debug("Email is already verified and signed up.")
		return res, fmt.Errorf(`%s has already signed up`, params.Email)
	} else if existingUser.ID != "" && existingUser.EmailVerifiedAt == nil {
		logrus.Debug("Email is already signed up. Verification pending...")
		return res, fmt.Errorf("%s has already signed up. please complete the email verification process or reset the password", params.Email)
	}

	var roles []string

	if len(params.Roles) > 0 {
		rolesString, err := env.GetEnvByKey(constants.Roles)
		if err != nil {
			logrus.Debug("Error getting roles: ", err)
			return res, err
		} else {
			roles = strings.Split(rolesString, ",")
		}
		if !validators.IsValidRoles(params.Roles, roles) {
			logrus.Debug("Invalid roles: ", params.Roles)
			return res, fmt.Errorf(`invalid roles`)
		} else {
			roles = params.Roles
		}
	} else {
		inputRolesString, err := env.GetEnvByKey(constants.DefaultRoles)
		if err != nil {
			logrus.Debug("Error getting default roles: ", err)
			return res, err
		} else {
			roles = strings.Split(inputRolesString, ",")
		}
	}

	user := models.User{
		Email: params.Email,
	}

	user.Roles = strings.Join(roles, ",")

	password, _ := crypto.EncryptPassword(params.Password)
	user.Password = &password

	if params.UserName != nil {
		user.UserName = params.UserName
	}

	if params.FamilyName != nil {
		user.FamilyName = params.FamilyName
	}

	if params.GivenName != nil {
		user.GivenName = params.GivenName
	}

	if params.MiddleName != nil {
		user.MiddleName = params.MiddleName
	}

	if params.Nickname != nil {
		user.NickName = params.Nickname
	}

	if params.Gender != nil {
		user.Gender = params.Gender
	}

	if params.BirthDate != nil {
		user.BirthDate = params.BirthDate
	}

	if params.PhoneNumber != nil {
		user.PhoneNumber = params.PhoneNumber
	}

	if params.Picture != nil {
		user.Picture = params.Picture
	}

	if params.IsMultiFactorAuthEnabled != nil {
		user.IsMultiFactorAuthEnabled = params.IsMultiFactorAuthEnabled
	}

	user.SignUpMethods = constants.AuthRecipeMethodBasicAuth

	user, err = maindb.Provider.AddUser(ctx, user)
	if err != nil {
		logrus.Debug("Failed to add user: ", err)
		return res, err
	}

	userToReturn := user.AsAPIUser()

	scope := []string{"email", "profile"}

	authToken, err := token.CreateAuthTokens(gc, user, roles, scope, constants.AuthRecipeMethodBasicAuth)
	if err != nil {
		logrus.Debug("Failed to create auth tokens: ", err)
		return res, err
	}

	expiresIn := authToken.AccessToken.ExpiresAt - time.Now().Unix()
	if expiresIn <= 0 {
		expiresIn = 1
	}

	sessionKey := constants.AuthRecipeMethodBasicAuth + ":" + user.ID
	refreshTokenHash, err := crypto.EncryptPassword(authToken.RefreshToken.Token)
	if err != nil {
		logrus.Debug("Failed to hash refresh tokens: ", err)
		return res, err
	}
	memorydb.Provider.SetSession(ctx, sessionKey, refreshTokenHash)

	res = &model.AuthResponse{
		Message:      `Signed up successfully.`,
		AccessToken:  &authToken.AccessToken.Token,
		ExpiresIn:    &expiresIn,
		RefreshToken: &authToken.RefreshToken.Token,
		User:         userToReturn,
	}

	return res, nil
}
