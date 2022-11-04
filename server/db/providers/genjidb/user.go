package genjidb

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/ArkamFahry/GateGuardian/server/constants"
	"github.com/ArkamFahry/GateGuardian/server/db/models"
	"github.com/ArkamFahry/GateGuardian/server/env"
	"github.com/genjidb/genji/document"
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

	query := fmt.Sprintf(`INSERT INTO %s %s VALUES %s`, models.Collections.User, fields, values)
	err = p.db.Exec(query)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (p *provider) UpdateUser() {

}

func (p *provider) DeleteUser() {

}

func (p *provider) ListUser() {

}

func (p *provider) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf(`SELECT * FROM %s WHERE email == '%s'`, models.Collections.User, email)
	res, err := p.db.QueryDocument(query)
	if err != nil {
		return user, err
	}
	document.StructScan(res, &user)

	return user, nil
}

func (p *provider) GetUserByID() {

}

func (p *provider) UpdateUsers() {

}
