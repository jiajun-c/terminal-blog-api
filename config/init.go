package config

import "github.com/spf13/viper"

type Config struct {
	Dsn string // the dsn url for database
}

func InitConfig() Config {
	config := Config{
		Dsn: viper.GetString("db.dsn"),
	}
	return config
}
