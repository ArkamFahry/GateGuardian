package sqlite

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/ArkamFahry/GateGuardian/server/db/models"
	"github.com/google/uuid"
)

func (p *provider) AddUser(ctx context.Context, user models.User) (models.User, error) {

	if user.Id == "" {
		user.Id = uuid.New().String()
	}

	user.CreatedAt = time.Now().Unix()
	user.UpdatedAt = time.Now().Unix()

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

	insertUserQuery := fmt.Sprintf(`INSERT INTO %s %s Values %s`, models.Models.User, fields, values)
	p.db.Exec(insertUserQuery)

	return user, nil
}

func (p *provider) UpdateUser(ctx context.Context, user models.User) (models.User, error) {
	user.UpdatedAt = time.Now().Unix()

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

	queryUpdateUser := fmt.Sprintf("UPDATE %s SET %s WHERE id = '%s'", models.Models.User, updateFields, user.Id)

	_, err = p.db.Exec(queryUpdateUser)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (p *provider) DeleteUser(ctx context.Context, user models.User) error {

	return nil
}

func (p *provider) ListUsers(ctx context.Context) ([]models.User, error) {
	var users []models.User
	var user models.User

	getUsersQuery := fmt.Sprintf(`SELECT * FROM %s`, models.Models.User)

	rows, err := p.db.Query(getUsersQuery)
	if err != nil {
		return users, err
	}

	for rows.Next() {
		rows.Scan(&user.Id, &user.Email, &user.Password, &user.GivenName, &user.FamilyName, &user.MiddleName, &user.NickName, &user.Gender, &user.CreatedAt, &user.UpdatedAt)

		users = append(users, user)
	}

	return users, nil
}

func (p *provider) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	var user models.User

	getUserByEmailQuery := fmt.Sprintf(`SELECT * FROM %s WHERE email == '%s'`, models.Models.User, email)

	err := p.db.QueryRow(getUserByEmailQuery).Scan(&user.Id, &user.Email, &user.Password, &user.GivenName, &user.FamilyName, &user.MiddleName, &user.NickName, &user.Gender, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (p *provider) GetUserByID(ctx context.Context, id string) (models.User, error) {
	var user models.User

	getUserByIdQuery := fmt.Sprintf(`SELECT * FROM %s WHERE id == '%s'`, models.Models.User, id)

	err := p.db.QueryRow(getUserByIdQuery).Scan(&user.Id, &user.Email, &user.Password, &user.GivenName, &user.FamilyName, &user.MiddleName, &user.NickName, &user.Gender, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (p *provider) UpdateUsers(ctx context.Context, data map[string]interface{}, ids []string) error {

	return nil
}
