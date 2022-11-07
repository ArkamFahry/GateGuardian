package env

import "github.com/spf13/viper"

func GetEnv() {
	viper.SetConfigName("config")

	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	viper.SetConfigType("json")
}
