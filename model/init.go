package model

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"terminal-blog-api/config"
)

type Result struct {
	db  *gorm.DB
	err error
}

func InitDatabase(cfg config.Config) Result {
	dsn := cfg.Dsn
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return Result{db, err}
}
