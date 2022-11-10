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

// Inserts a user into the database
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

	query := fmt.Sprintf(`INSERT INTO %s %s VALUES %s`, models.Model.User, fields, values)
	p.db.Exec(query)

	return user, nil
}

// Updates a user by id and before the update omits the null value in the model and selectively updates only the felids that are not null
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

	query := fmt.Sprintf(`UPDATE %s SET %s WHERE id = '%s'`, models.Model.User, updateFields, user.Id)

	_, err = p.db.Exec(query)
	if err != nil {
		return user, err
	}

	return user, nil
}

// Deletes a user by id
func (p *provider) DeleteUser(ctx context.Context, user models.User) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = '%s'`, models.Model.User, user.Id)

	_, err := p.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

// Gets a list of users
func (p *provider) ListUsers(ctx context.Context, pagination models.Pagination) ([]models.User, error) {
	var users []models.User
	var user models.User

	query := fmt.Sprintf(`SELECT * FROM %s LIMIT %d OFFSET %d`, models.Model.User, pagination.Limit, pagination.Offset)

	rows, err := p.db.Query(query)
	if err != nil {
		return users, err
	}

	for rows.Next() {
		rows.Scan(&user.Id, &user.Email, &user.Password, &user.GivenName, &user.FamilyName, &user.MiddleName, &user.NickName, &user.Gender, &user.BirthDate, &user.Picture, &user.CreatedAt, &user.UpdatedAt)

		users = append(users, user)
	}

	return users, nil
}

// Gets a single the user by email
func (p *provider) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	var user models.User

	query := fmt.Sprintf(`SELECT * FROM %s WHERE email = '%s'`, models.Model.User, email)

	err := p.db.QueryRow(query).Scan(&user.Id, &user.Email, &user.Password, &user.GivenName, &user.FamilyName, &user.MiddleName, &user.NickName, &user.Gender, &user.BirthDate, &user.Picture, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return user, err
	}

	return user, nil
}

// Gets a single user by id
func (p *provider) GetUserByID(ctx context.Context, id string) (models.User, error) {
	var user models.User

	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = '%s'`, models.Model.User, id)

	err := p.db.QueryRow(query).Scan(&user.Id, &user.Email, &user.Password, &user.GivenName, &user.FamilyName, &user.MiddleName, &user.NickName, &user.Gender, &user.BirthDate, &user.Picture, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return user, err
	}

	return user, nil
}

// Updates a list of users by ids or if ids not given updates all the users
func (p *provider) UpdateUsers(ctx context.Context, data map[string]interface{}, ids []string) error {

	return nil
}
