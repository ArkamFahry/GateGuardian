package sql

import (
	"gategaurdian/server/constants"
	"gategaurdian/server/database/maindb/models"
	"gategaurdian/server/database/memorydb"
	"time"

	"github.com/glebarez/sqlite"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type provider struct {
	db *gorm.DB
}

// This NewProvider represents the gorm sql database provider
func NewProvider() (*provider, error) {
	var sqlDb *gorm.DB
	var err error

	// Custom logger setup for gorm
	customLogger := logger.New(
		logrus.StandardLogger(),
		logger.Config{
			SlowThreshold:             time.Second,  // Slow SQL threshold
			LogLevel:                  logger.Error, // Log level
			IgnoreRecordNotFoundError: true,         // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,        // Disable color
		},
	)

	ormConfig := &gorm.Config{
		Logger: customLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: models.Prefix,
		},
		AllowGlobalUpdate: true,
	}

	dbType := memorydb.RequiredEnvStoreObj.GetRequiredEnv().DatabaseType
	dbUrl := memorydb.RequiredEnvStoreObj.GetRequiredEnv().DatabaseUrl

	// Depending on the sql db type specific database driver is initialized
	switch dbType {
	case constants.DbTypeSqlite:
		sqlDb, err = gorm.Open(sqlite.Open(dbUrl+"?_pragma=busy_timeout(5000)&_pragma=journal_mode(WAL)"), ormConfig)
	}

	if err != nil {
		return nil, err
	}

	// Runs an automatic migration on sql database create the tables required by gate_guardian to operate
	err = sqlDb.AutoMigrate(&models.User{}, &models.VerificationRequest{}, &models.Session{}, &models.Env{})
	if err != nil {
		return nil, err
	}

	return &provider{
		db: sqlDb,
	}, nil
}
