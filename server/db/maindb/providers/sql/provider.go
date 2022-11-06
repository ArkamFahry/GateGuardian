package sql

import (
	"github.com/ArkamFahry/GateGuardian/server/constants"
	"github.com/ArkamFahry/GateGuardian/server/db/maindb/models"
	"github.com/ArkamFahry/GateGuardian/server/env"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type provider struct {
	db *gorm.DB
}

func NewSqlProvider() (*provider, error) {
	var sqlDB *gorm.DB
	var err error
	dbType, _ := env.GetEnvByKey(constants.DatabaseType)
	dbURL, _ := env.GetEnvByKey(constants.DatabaseURL)

	switch dbType {
	case constants.DbTypePostgres, constants.DbTypeYugabyte, constants.DbTypeCockroachDB:
		sqlDB, err = gorm.Open(postgres.Open(dbURL))
	case constants.DbTypeSqlite:
		sqlDB, err = gorm.Open(sqlite.Open(dbURL))
	case constants.DbTypeMysql, constants.DbTypeMariaDB, constants.DbTypePlanetScaleDB:
		sqlDB, err = gorm.Open(mysql.Open(dbURL))
	case constants.DbTypeSqlserver:
		sqlDB, err = gorm.Open(sqlserver.Open(dbURL))
	}

	if err != nil {
		return nil, err
	}

	err = sqlDB.AutoMigrate(&models.User{}, &models.Session{})
	if err != nil {
		return nil, err
	}

	return &provider{
		db: sqlDB,
	}, err
}
