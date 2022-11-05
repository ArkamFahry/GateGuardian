package sqlite

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/ArkamFahry/GateGuardian/server/constants"
	"github.com/ArkamFahry/GateGuardian/server/db/maindb/models"
	"github.com/ArkamFahry/GateGuardian/server/env"
	"github.com/google/uuid"
)

func (p *provider) AddUser(user models.User) (models.User, error) {
	if user.ID == "" {
		user.ID = uuid.New().String()
	}

	if user.Roles == "" {
		defaultRole, _ := env.GetEnvByKey(constants.DefaultRoles)
		user.Roles = defaultRole
	}

	user.CreatedAt = time.Now().Unix()
	user.UpdatedAt = time.Now().Unix()
	user.LastLoggedIn = time.Now().Unix()

	bytes, err := json.Marshal(user)
	if err != nil {
		return user, err
	}

	decoder := json.NewDecoder(strings.NewReader(string(bytes)))
	decoder.UseNumber()
	userMap := map[string]interface{}{}
	err = decoder.Decode(&userMap)
	if err != nil {
		return user, err
	}

	fields := "("
	values := "("
	for key, value := range userMap {
		if value != nil {
			if key == "_id" {
				fields += "id,"
			} else {
				fields += key + ","
			}

			valueType := reflect.TypeOf(value)
			if valueType.Name() == "string" {
				values += fmt.Sprintf("'%s',", value.(string))
			} else {
				values += fmt.Sprintf("%v,", value)
			}
		}
	}

	fields = fields[:len(fields)-1] + ")"
	values = values[:len(values)-1] + ")"

	insertUserQuery := fmt.Sprintf(`INSERT INTO %s %s VALUES %s`, models.Collections.User, fields, values)
	_, err = p.db.Exec(insertUserQuery)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (p *provider) UpdateUser(user models.User) (models.User, error) {
	user.UpdatedAt = time.Now().Unix()

	bytes, err := json.Marshal(user)
	if err != nil {
		return user, err
	}
	// use decoder instead of json.Unmarshall, because it converts int64 -> float64 after unmarshalling
	decoder := json.NewDecoder(strings.NewReader(string(bytes)))
	decoder.UseNumber()
	userMap := map[string]interface{}{}
	err = decoder.Decode(&userMap)
	if err != nil {
		return user, err
	}

	updateFields := ""
	for key, value := range userMap {
		if key == "_id" {
			continue
		}

		if key == "_key" {
			continue
		}

		if value == nil {
			updateFields += fmt.Sprintf("%s = null, ", key)
			continue
		}

		valueType := reflect.TypeOf(value)
		if valueType.Name() == "string" {
			updateFields += fmt.Sprintf("%s = '%s', ", key, value.(string))
		} else {
			updateFields += fmt.Sprintf("%s = %v, ", key, value)
		}
	}
	updateFields = strings.Trim(updateFields, " ")
	updateFields = strings.TrimSuffix(updateFields, ",")

	updateUserQuery := fmt.Sprintf("UPDATE %s SET %s WHERE id = '%s'", models.Collections.User, updateFields, user.ID)
	_, err = p.db.Exec(updateUserQuery)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (p *provider) DeleteUser() {

}

func (p *provider) ListUser() {

}

func (p *provider) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	getUserByEmailQuery := fmt.Sprintf(`SELECT * FROM %s WHERE email == '%s'`, models.Collections.User, email)
	err := p.db.QueryRow(getUserByEmailQuery).Scan(&user.ID, &user.Email, &user.EmailVerifiedAt, &user.Password, &user.SignUpMethods, &user.UserName, &user.FamilyName, &user.GivenName, &user.MiddleName, &user.NickName, &user.Gender, &user.BirthDate, &user.PhoneNumber, &user.PhoneNumberVerifiedAt, &user.Picture, &user.Roles, &user.RevokedTimestamp, &user.IsMultiFactorAuthEnabled, &user.UpdatedAt, &user.CreatedAt, &user.LastLoggedIn)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (p *provider) GetUserByID(id string) (models.User, error) {
	var user models.User
	getUserByIdQuery := fmt.Sprintf(`SELECT * FROM %s WHERE id == '%s'`, models.Collections.User, id)
	err := p.db.QueryRow(getUserByIdQuery).Scan(&user.ID, &user.Email, &user.EmailVerifiedAt, &user.Password, &user.SignUpMethods, &user.UserName, &user.FamilyName, &user.GivenName, &user.MiddleName, &user.NickName, &user.Gender, &user.BirthDate, &user.PhoneNumber, &user.PhoneNumberVerifiedAt, &user.Picture, &user.Roles, &user.RevokedTimestamp, &user.IsMultiFactorAuthEnabled, &user.UpdatedAt, &user.CreatedAt, &user.LastLoggedIn)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (p *provider) UpdateUsers() {

}
