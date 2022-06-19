//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"kratos-ota/internal/biz"
	"kratos-ota/internal/conf"
	"kratos-ota/internal/data"
	"kratos-ota/internal/service"
	"kratos-ota/internal/task"
)

func wireCron(*conf.Data, *conf.Auth, log.Logger) (*task.OtaCron, func(), error) {
	panic(wire.Build(data.ProviderSet, biz.ProviderSet, service.ProviderSet, newCron))
}
