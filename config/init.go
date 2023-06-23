package config

import "github.com/spf13/viper"

type Config struct {
	dsn string // the dsn url for database
}

func InitConfig() Config {
	config := Config{
		dsn: viper.GetString("db.dsn"),
	}
	return config
}
