package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"kratos-ota/internal/biz"
)

type calendarRepo struct {
	data *Data
	log  *log.Helper
}

func NewCalendarRepo(data *Data, logger log.Logger) biz.CalendarRepo {
	return &calendarRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
