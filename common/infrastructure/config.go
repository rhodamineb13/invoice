package infrastructure

import (
	"github.com/spf13/viper"
)

type ConfigEnv struct {
	Host          string `mapstructure:"DB_HOST"`
	Port          string `mapstructure:"DB_PORT"`
	User          string `mapstructure:"DB_USER"`
	Password      string `mapstructure:"DB_PASSWORD"`
	DBname        string `mapstructure:"DB_NAME"`
	Issuer        string `mapstructure:"ISSUER"`
	LibSecretKey  string `mapstructure:"SECRET_KEY"`
	Duration      int    `mapstructure:"EXPIRY"`
	AdminEmail    string `mapstructure:"ADMIN_EMAIL"`
	AdminPassword string `mapstructure:"ADMIN_PASSWORD"`
}

var config = NewConfig()

func NewConfig() *ConfigEnv {
	var conf *ConfigEnv

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
