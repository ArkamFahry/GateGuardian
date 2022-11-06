package env

import (
	"os"

	"github.com/ArkamFahry/GateGuardian/server/constants"
	"github.com/ArkamFahry/GateGuardian/server/crypto"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Env struct {
	DatabaseType            string
	DatabaseURL             string
	DatabaseName            string
	DatabaseNameSpace       string
	DatabaseUsername        string
	DatabasePassword        string
	Port                    string
	EncryptionKey           string
	JwtType                 string
	JwtSecret               string
	JwtPrivateKey           string
	JwtPublicKey            string
	JwtRoleClaim            string
	ClientID                string
	AccessTokenExpiryTime   string
	CustomAccessTokenScript string
	Roles                   string
	DefaultRoles            string
	AppURL                  string
	GateGuardianURL         string
}

func GetEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Info("Failed to load .env file: ", err, "|switched to os env|")
	}

	dbType := os.Getenv(constants.DatabaseType)
	dbUrl := os.Getenv(constants.DatabaseURL)
	dbName := os.Getenv(constants.DatabaseName)
	dbNameSpace := os.Getenv(constants.DatabaseNameSpace)
	dbUserName := os.Getenv(constants.DatabaseUsername)
	dbPassword := os.Getenv(constants.DatabasePassword)
	port := os.Getenv(constants.Port)
	encryptionKey := os.Getenv(constants.EncryptionKey)
	jwtType := os.Getenv(constants.JwtType)
	jwtSecret := os.Getenv(constants.JwtSecret)
	jwtPrivateKey := os.Getenv(constants.JwtPrivateKey)
	jwtPublicKey := os.Getenv(constants.JwtPublicKey)
	jwtRoleClaim := os.Getenv(constants.JwtRoleClaim)
	clientID := os.Getenv(constants.ClientID)
	accessTokenExpiryTime := os.Getenv(constants.AccessTokenExpiryTime)
	customAccessTokenScript := os.Getenv(constants.CustomAccessTokenScript)
	roles := os.Getenv(constants.Roles)
	defaultRoles := os.Getenv(constants.DefaultRoles)
	appURL := os.Getenv(constants.AppURL)
	gateGuardianURL := os.Getenv(constants.GateGuardianURL)

	if dbType == "" {
		dbType = "sqlite"
	}

	if port == "" {
		port = "8080"
	}

	if encryptionKey == "" {
		encryptionKey = uuid.New().String()[:36-4]
	}

	if clientID == "" {
		clientID = uuid.New().String()
	}

	if jwtType == "" {
		jwtType = "RS256"
	}

	if jwtPrivateKey == "" || jwtPublicKey == "" {
		if crypto.IsRSA(jwtType) {
			_, jwtPrivateKey, jwtPublicKey, _, _ = crypto.NewRSAKey(jwtType, clientID)
		} else if crypto.IsECDSA(jwtType) {
			_, jwtPrivateKey, jwtPublicKey, _, _ = crypto.NewECDSAKey(jwtType, clientID)
		}
	}

	if jwtRoleClaim == "" {
		jwtRoleClaim = "role"
	}

	if accessTokenExpiryTime == "" {
		accessTokenExpiryTime = "15m"
	}

	if roles == "" {
		roles = "user"
		defaultRoles = "user"
	}

	if defaultRoles == "" {
		defaultRoles = "user"
	}

	if appURL == "" {
		appURL = "*"
	}

	if gateGuardianURL == "" {
		gateGuardianURL = "*"
	}

	env := Env{
		DatabaseType:            dbType,
		DatabaseURL:             dbUrl,
		DatabaseName:            dbName,
		DatabaseNameSpace:       dbNameSpace,
		DatabaseUsername:        dbUserName,
		DatabasePassword:        dbPassword,
		Port:                    port,
		EncryptionKey:           encryptionKey,
		JwtType:                 jwtType,
		JwtSecret:               jwtSecret,
		JwtPrivateKey:           jwtPrivateKey,
		JwtPublicKey:            jwtPublicKey,
		JwtRoleClaim:            jwtRoleClaim,
		AccessTokenExpiryTime:   accessTokenExpiryTime,
		CustomAccessTokenScript: customAccessTokenScript,
		ClientID:                clientID,
		Roles:                   roles,
		DefaultRoles:            defaultRoles,
		AppURL:                  appURL,
		GateGuardianURL:         gateGuardianURL,
	}

	PersistEnv(env)
}
