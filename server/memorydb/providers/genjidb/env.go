package genjidb

import (
	"strings"
	"time"

	"github.com/ArkamFahry/GateGuardian/server/memorydb/models"
	"github.com/genjidb/genji/document"
	"github.com/genjidb/genji/types"
	"github.com/sirupsen/logrus"
)

func (p *provider) AddEnv(key string, data string) (string, error) {
	err := p.memorydb.Exec(`INSERT INTO env (id, data, created_at, updated_at) VALUES (?, ?, ?, ?);`, key, data, time.Now().Unix(), time.Now().Unix())
	if err != nil {
		logrus.Debug("Failed to insert env: ", err)
	}

	return key, err
}

func (p *provider) UpdateEnv(key string, data string) (string, error) {
	err := p.memorydb.Exec(`UPDATE env SET data = ?, updated_at = ? WHERE id == ?`, data, time.Now().Unix(), key)
	if err != nil {
		logrus.Debug("Failed to update env: ", err)
	}

	return key, err
}

func (p *provider) DeleteEnv(key string) error {
	err := p.memorydb.Exec(`DELETE FROM env WHERE id == ?`, key)
	if err != nil {
		logrus.Info("Failed to delete env: ", err)
	}

	return err
}

func (p *provider) GetEnvByKey(key string) (string, error) {
	var env string
	res, err := p.memorydb.QueryDocument(`SELECT data FROM env WHERE id == ?;`, key)
	if err != nil {
		logrus.Debug("No such env present in db: ", err)
	} else {
		data, err := res.GetByField("data")
		if err != nil {
			logrus.Error("No Such Field: ", err)
		}

		env = strings.Trim(data.String(), `"`)
	}

	return env, err
}

func (p *provider) ListEnv() ([]models.Env, error) {
	res, err := p.memorydb.Query(`SELECT * FROM env;`)
	if err != nil {
		logrus.Debug("Failed to get envs: ", err)
	}

	var envs []models.Env
	var env models.Env

	res.Iterate(
		func(d types.Document) error {
			err = document.StructScan(d, &env)
			if err != nil {
				return err
			}

			envs = append(envs, env)

			return nil
		})

	return envs, err
}
