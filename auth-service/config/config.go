package config

import "github.com/spf13/viper"

type ApiInfo struct {
	Port     string `mapstructure:"port"`
	Host     string `mapstructure:"host"`
	Version  string `mapstructure:"version"`
	LogLevel string `mapstructure:"log_level"`
}

type Schema struct {
	ApiInfo     ApiInfo `mapstructure:"api_info"`
	PostgresDSN string  `mapstructure:"pg_dsn"`
}

var Config *Schema

func LoadConfig() {
	viper.SetConfigFile("config/config.yml")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	err := viper.Unmarshal(&Config)
	if err != nil {
		panic(err)
	}
}
