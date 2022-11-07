package env

import (
	"github.com/spf13/viper"
)

type Envs struct {
	Environment string
	App         struct {
		Port string
	}
	Db struct {
		Type      string
		Url       string
		Port      string
		Host      string
		Name      string
		NameSpace string
		KeySpace  string
		UserName  string
		Password  string
		Cert      string
	}
	Jwt struct {
		Secret string
	}
	Role struct {
		Roles          string
		DefaultRoles   string
		ProtectedRoles string
	}
}

func GetEnv() error {
	var envs Envs

	viper.SetConfigName("config")

	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	viper.SetConfigType("json")

	viper.ReadInConfig()

	viper.Unmarshal(&envs)

	return nil
}
