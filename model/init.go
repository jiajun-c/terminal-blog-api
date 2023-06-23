package model

import "terminal-blog-api/config"

type Result struct {
	cfg config.Config
}

func InitDatabase(cfg config.Config) Result {
	return Result{}
}
