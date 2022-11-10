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

// Adds session to the database
func (p *provider) AddSession(ctx context.Context, session models.Session) error {
	if session.Id == "" {
		session.Id = uuid.New().String()
	}

	session.CreatedAt = time.Now().Unix()
	session.UpdatedAt = time.Now().Unix()

	bytes, err := json.Marshal(session)
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(strings.NewReader(string(bytes)))
	decoder.UseNumber()
	sessionMap := map[string]interface{}{}
	err = decoder.Decode(&sessionMap)
	if err != nil {
		return err
	}

	fields := "("
	values := "("
	for key, value := range sessionMap {
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

	query := fmt.Sprintf(`INSERT INTO %s %s VALUES %s`, models.Model.Session, fields, values)

	_, err = p.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
