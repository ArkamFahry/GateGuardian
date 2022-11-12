package sql

import (
	"github.com/ArkamFahry/GateGuardian/server/constants"
	"github.com/ArkamFahry/GateGuardian/server/db/models"
	"github.com/ArkamFahry/GateGuardian/server/memorystore/envstore"
	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type provider struct {
	db *gorm.DB
}

func NewProvider() (*provider, error) {
	var sqlDB *gorm.DB
	var err error

	dbType, _ := envstore.Provider.GetEnv(constants.DbType)
	dbUrl, _ := envstore.Provider.GetEnv(constants.DbUrl)

	switch dbType {
	case constants.DbTypePostgresql:
		sqlDB, err = gorm.Open(postgres.Open(dbUrl))
	case constants.DbTypeSqlite:
		sqlDB, err = gorm.Open(sqlite.Open(dbUrl + "?_pragma=busy_timeout(5000)&_pragma=journal_mode(WAL)"))
	case constants.DbTypeMysql:
		sqlDB, err = gorm.Open(mysql.Open(dbUrl))
	case constants.DbTypeSqlserver:
		sqlDB, err = gorm.Open(sqlserver.Open(dbUrl))
	}

	err = sqlDB.AutoMigrate(&models.User{}, &models.Session{}, &models.State{})
	if err != nil {
		return nil, err
	}

	return &provider{
		db: sqlDB,
	}, nil
}
