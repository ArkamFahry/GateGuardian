package memdb

import (
	"time"

	"github.com/ArkamFahry/GateGuardian/server/db/envdb/models"
	"github.com/sirupsen/logrus"
)

func (p *provider) AddEnv(key string, data string) (string, error) {

	env := models.Env{
		ID:        key,
		Data:      data,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	txn := p.db.Txn(true)

	err := txn.Insert("env", env)
	if err != nil {
		logrus.Info("memdb insert transaction failed: ", err)
	}

	txn.Commit()

	return key, nil
}

func (p *provider) UpdateEnv(key string, data string) (string, error) {

	env := models.Env{
		ID:        key,
		Data:      data,
		UpdatedAt: time.Now().Unix(),
	}

	txn := p.db.Txn(true)

	err := txn.Insert("env", env)
	if err != nil {
		logrus.Info("memdb update transaction failed: ", err)
	}

	txn.Commit()

	return key, nil
}

func (p *provider) DeleteEnv(key string) error {

	env := models.Env{
		ID: key,
	}

	txn := p.db.Txn(true)

	err := txn.Delete("env", env)
	if err != nil {
		logrus.Info("memdb delete transaction failed: ", err)
	}

	txn.Commit()

	return nil
}

func (p *provider) GetEnvByKey(key string) (string, error) {
	txn := p.db.Txn(false)

	raw, err := txn.First("env", "id", key)
	if err != nil {
		logrus.Info("memdb read first transaction failed: ", err)
	}

	env := raw.(models.Env).Data

	return env, nil
}

func (p *provider) ListEnv() ([]models.Env, error) {
	var envs []models.Env
	var env models.Env

	txn := p.db.Txn(false)

	raw, err := txn.Get("env", "id")
	if err != nil {
		logrus.Info("memdb read all transaction failed: ", err)
	}

	for obj := raw.Next(); obj != nil; obj = raw.Next() {

		env.ID = obj.(models.Env).ID
		env.Data = obj.(models.Env).Data
		env.CreatedAt = obj.(models.Env).CreatedAt
		env.UpdatedAt = obj.(models.Env).UpdatedAt

		envs = append(envs, env)
	}

	return envs, nil
}
