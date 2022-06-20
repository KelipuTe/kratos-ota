package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-ota/internal/biz"
	"time"
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

type CalendarQueue struct {
	Ota        string
	HotelId    string
	RoomTypeId string
	SyncType   int
	CreateTime time.Time
	FlowId     string
	Status     int
	Msg        string
	UpdateTime time.Time
}

func (CalendarQueue) TableName() string {
	return "calendar_queue"
}

type CalendarDetail struct {
	FlowId     string
	Ota        string
	HotelId    string
	RoomTypeId string
	Date       time.Time
	Num        int
	Price      int
}

func (CalendarDetail) TableName() string {
	return "calendar_detail"
}

func (c *calendarRepo) SaveCalendarDetail(ctx context.Context, p1cq *biz.CalendarQueue) error {
	arr1Len := len(p1cq.Arr1CalendarDetail)
	arr1p1c := make([]*CalendarDetail, arr1Len)
	for i := 0; i < arr1Len; i++ {
		arr1p1c[i] = &CalendarDetail{
			FlowId:     p1cq.FlowId,
			Ota:        p1cq.Ota,
			HotelId:    p1cq.HotelId,
			RoomTypeId: p1cq.RoomTypeId,
			Date:       p1cq.Arr1CalendarDetail[i].Date,
			Num:        p1cq.Arr1CalendarDetail[i].Num,
			Price:      p1cq.Arr1CalendarDetail[i].Price,
		}
	}
	c.data.db.CreateInBatches(arr1p1c, arr1Len)
	return nil
}

func (c *calendarRepo) SaveCalendarQueue(ctx context.Context, p1cq *biz.CalendarQueue) error {
	p1c := &CalendarQueue{
		Ota:        p1cq.Ota,
		HotelId:    p1cq.HotelId,
		RoomTypeId: p1cq.RoomTypeId,
		SyncType:   p1cq.SyncType,
		CreateTime: p1cq.CreateTime,
		FlowId:     p1cq.FlowId,
		Status:     0,
		Msg:        "",
		UpdateTime: p1cq.CreateTime,
	}
	c.data.db.Create(p1c)
	return nil
}
