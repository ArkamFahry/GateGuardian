package env

import (
	"github.com/spf13/viper"
)

type Envs struct {
	PORT    string
	DB_TYPE string
	DB_URL  string
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
