package env

import (
	"github.com/spf13/viper"
)

type Envs struct {
	PORT                     string
	DB_TYPE                  string
	DB_URL                   string
	JWT_SECRET               string
	JWT_TYPE                 string
	JWT_PUBLIC_KEY           string
	JWT_PRIVATE_KEY          string
	ACCESS_TOKEN_EXPIRY_TIME string
	CLIENT_ID                string
	ALLOWED_ROLES            string
	DEFAULT_ROLES            string
	DEFAULT_ROLE             string
}

func GetEnv() error {
	var envs Envs

	viper.SetConfigName("app")

	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	viper.SetConfigType("env")

	viper.ReadInConfig()

	viper.Unmarshal(&envs)

	PersistEnv(envs)

	return nil
}
