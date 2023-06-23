//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"terminal-blog-api/config"
	"terminal-blog-api/model"
)

func InitBlog() model.Result {
	wire.Build(model.InitDatabase, config.InitConfig)
	return model.Result{}
}
