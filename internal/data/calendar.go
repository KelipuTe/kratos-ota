package data

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-ota/internal/biz"
	"strconv"
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

const (
	calendarSyncStatusTodo int = 0
)

// ## gorm ##

type CalendarQueue struct {
	FlowId     string
	RoomTypeId int
	SyncType   int
	CreateAt   time.Time
	Status     int
	Msg        string
	UpdateAt   time.Time
}

func (CalendarQueue) TableName() string {
	return "calendar_queue"
}

type CalendarDetail struct {
	FlowId     string
	RoomTypeId int
	Date       time.Time
	Num        int
	Price      int
}

func (CalendarDetail) TableName() string {
	return "calendar_detail"
}

func (c *calendarRepo) SaveCalendarDetail(ctx context.Context, p1cq *biz.CalendarQueue) error {
	arr1Len := len(p1cq.Arr1P1CalendarDetail)
	arr1p1c := make([]*CalendarDetail, arr1Len)
	for i := 0; i < arr1Len; i++ {
		arr1p1c[i] = &CalendarDetail{
			FlowId:     p1cq.FlowId,
			RoomTypeId: p1cq.RoomTypeId,
			Date:       p1cq.Arr1P1CalendarDetail[i].Date,
			Num:        p1cq.Arr1P1CalendarDetail[i].Num,
			Price:      p1cq.Arr1P1CalendarDetail[i].Price,
		}
	}
	c.data.db.CreateInBatches(arr1p1c, arr1Len)
	return nil
}

func (c *calendarRepo) SaveCalendarQueue(ctx context.Context, p1cq *biz.CalendarQueue) error {
	p1c := &CalendarQueue{
		RoomTypeId: p1cq.RoomTypeId,
		SyncType:   p1cq.SyncType,
		CreateAt:   p1cq.CreateAt,
		FlowId:     p1cq.FlowId,
		Status:     0,
		Msg:        "",
		UpdateAt:   p1cq.CreateAt,
	}
	c.data.db.Create(p1c)
	return nil
}

func (c *calendarRepo) ListCalendarQueueByRoomTypeId(ctx context.Context, id int) ([]*biz.CalendarQueue, error) {
	arr1p1c := make([]*biz.CalendarQueue, 1)

	c.data.db.Table("calendar_queue").
		Where("room_type_id = ?", id).
		Where("status in (?)", []int{calendarSyncStatusTodo}).
		Find(&arr1p1c)

	return arr1p1c, nil
}

// ## redis ##

const (
	keyCalendarQueue = "calendar_queue"
)

func (c *calendarRepo) RPUSHCalendarQueue(ctx context.Context, id int) error {
	val, err := c.data.rdb.RPush(ctx, keyCalendarQueue, id).Result()
	spew.Dump("RPUSHCalendarQueue", val, err)
	return nil
}

func (c *calendarRepo) LPOPCalendarQueue(ctx context.Context) (int, error) {
	val, err := c.data.rdb.LPop(ctx, keyCalendarQueue).Result()
	if nil != err {
		return -1, err
	}
	spew.Dump("LPopCalendarQueue", val, err)

	valInt, err := strconv.Atoi(val)
	if nil != err {
		return -1, err
	}

	return valInt, nil
}
