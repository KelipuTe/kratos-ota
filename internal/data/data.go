package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"kratos-learning/internal/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewDB, NewData, NewRoomTypeRepo, NewCalendarRepo)

func NewDB(c *conf.Data, logger log.Logger) *gorm.DB {
	if "mysql" == c.Database.Driver {
		dsn := c.Database.Source
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		return db
	}
	panic("wrong database driver")
}

type Data struct {
	db *gorm.DB
}

func NewData(c *conf.Data, logger log.Logger, db *gorm.DB) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{db: db}, cleanup, nil
}
