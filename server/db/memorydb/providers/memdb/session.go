package memdb

import (
	"github.com/ArkamFahry/GateGuardian/server/db/memorydb/models"
	"github.com/sirupsen/logrus"
)

func (p *provider) SetSession(key string, value string) (string, error) {

	session := models.Session{
		Key:   key,
		Value: value,
	}

	txn := p.db.Txn(true)

	err := txn.Insert("session", session)
	if err != nil {
		logrus.Info("memdb insert transaction failed: ", err)
	}

	txn.Commit()

	return key, nil
}

func (p *provider) GetSession(key string) (string, error) {

	txn := p.db.Txn(false)

	raw, err := txn.First("session", "id", key)
	if err != nil {
		logrus.Info("memdb read first transaction failed: ", err)
	}

	env := raw.(models.Session).Value

	return env, nil
}

func (p *provider) UpdateSession(key string, value string) (string, error) {

	session := models.Session{
		Key:   key,
		Value: value,
	}

	txn := p.db.Txn(true)

	err := txn.Insert("session", session)
	if err != nil {
		logrus.Info("memdb update transaction failed: ", err)
	}

	txn.Commit()

	return key, nil
}
func (p *provider) DeleteSession(key string) error {

	session := models.Session{
		Key: key,
	}

	txn := p.db.Txn(true)

	err := txn.Delete("session", session)
	if err != nil {
		logrus.Info("memdb delete transaction failed: ", err)
	}

	txn.Commit()

	return nil
}
