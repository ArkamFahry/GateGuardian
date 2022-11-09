package sql

import (
	"github.com/ArkamFahry/GateGuardian/server/constants"
	"github.com/ArkamFahry/GateGuardian/server/db/models"
	"github.com/ArkamFahry/GateGuardian/server/memorystore/envstore"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type provider struct {
	db *gorm.DB
}

func NewProvider() (*provider, error) {
	var sqlDb *gorm.DB
	var err error

	dbUrl, _ := envstore.Provider.GetEnv(constants.DB_URL)

	sqlDb, err = gorm.Open(sqlite.Open(dbUrl))

	if err != nil {
		return nil, err
	}

	err = sqlDb.AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}

	return &provider{
		db: sqlDb,
	}, nil
}
