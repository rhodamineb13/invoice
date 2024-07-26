package infrastructure

import (
	"github.com/spf13/viper"
)

type configEnv struct {
	Host          string `mapstructure:"DB_HOST"`
	Port          string `mapstructure:"DB_PORT"`
	User          string `mapstructure:"DB_USER"`
	Password      string `mapstructure:"DB_PASSWORD"`
	DBName        string `mapstructure:"DB_NAME"`
}

var config = NewConfig()

func NewConfig() *configEnv {
	var conf *configEnv

	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&conf); err != nil {
		panic(err)
	}

	return conf
}
