package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"kratos-ota/internal/conf"
	"runtime"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewDB, NewRDB, NewData, NewRoomTypeRepo, NewCalendarRepo)

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

func NewRDB(c *conf.Data, logger log.Logger) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     c.Redis.Addr,
		Password: c.Redis.Pass,
		DB:       int(c.Redis.Db),
		PoolSize: runtime.NumCPU() * 4,
	})
}

type Data struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewData(c *conf.Data, logger log.Logger, db *gorm.DB) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{db: db}, cleanup, nil
}
