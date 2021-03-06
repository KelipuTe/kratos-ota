// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2/log"
	"kratos-ota/internal/biz"
	"kratos-ota/internal/conf"
	"kratos-ota/internal/data"
	"kratos-ota/internal/service"
	"kratos-ota/internal/task"
)

// Injectors from wire.go:

func wireCron(confData *conf.Data, auth *conf.Auth, logger log.Logger) (*task.OtaCron, func(), error) {
	db := data.NewDB(confData, logger)
	dataData, cleanup, err := data.NewData(confData, logger, db)
	if err != nil {
		return nil, nil, err
	}
	roomTypeRepo := data.NewRoomTypeRepo(dataData, logger)
	calendarRepo := data.NewCalendarRepo(dataData, logger)
	otaUsecase := biz.NewOtaUsecase(roomTypeRepo, calendarRepo, auth, logger)
	otaService := service.NewOtaService(otaUsecase)
	otaCron := newCron(logger, otaService)
	return otaCron, func() {
		cleanup()
	}, nil
}
