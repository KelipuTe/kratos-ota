package service

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/metadata"
	v1 "kratos-ota/api/ota/v1"
	"kratos-ota/internal/biz"
	"kratos-ota/internal/conf"
	"time"
)

type OtaService struct {
	v1.UnimplementedOtaServer
	ouc *biz.OtaUsecase

	calendarJobStatus bool
}

// NewOtaService
// api 层 --> service 层 --> biz 层
// service 层应该把 api 层的数据结构转换成 biz 层的数据结构
func NewOtaService(ouc *biz.OtaUsecase) *OtaService {
	return &OtaService{
		ouc:               ouc,
		calendarJobStatus: false,
	}
}

func (s *OtaService) GetToken(ctx context.Context, in *v1.GetTokenRequest) (*v1.GetTokenReply, error) {
	p1t := &biz.Token{
		ClientId:     in.ClientId,
		ClientSecret: in.ClientSecret,
	}

	err := s.ouc.GetToken(ctx, p1t)
	if err != nil {
		return nil, errors.New(500, err.Error(), err.Error())
	}

	return &v1.GetTokenReply{
		ClientId: p1t.ClientId,
		Token:    p1t.Token,
	}, nil
}

func (s *OtaService) GetHotelRoomType(ctx context.Context, in *v1.GetHotelRoomTypeRequest) (*v1.GetHotelRoomTypeReply, error) {
	p1h := &biz.Hotel{
		Ota:     in.Ota,
		HotelId: in.HotelId,
		Arr1RoomType: []biz.RoomType{
			{RoomTypeId: in.RoomTypeId},
		},
	}

	if err := s.ouc.GetHotelRoomType(ctx, p1h); nil != err {
		return nil, err
	}

	return &v1.GetHotelRoomTypeReply{
		Ota:     in.Ota,
		HotelId: in.HotelId,
		RoomType: &v1.RoomType{
			RoomTypeId:   p1h.Arr1RoomType[0].RoomTypeId,
			RoomTypeName: p1h.Arr1RoomType[0].RoomTypeName,
		},
	}, nil
}

func (s *OtaService) PushCalendar(ctx context.Context, in *v1.PushCalendarRequest) (*v1.PushCalendarReply, error) {
	var flowId string
	if md, ok := metadata.FromServerContext(ctx); ok {
		flowId = md.Get("x-md-global-flowid")
	}

	p1h := &biz.Hotel{
		Ota:     in.Ota,
		HotelId: in.HotelId,
		Arr1RoomType: []biz.RoomType{
			{RoomTypeId: in.RoomTypeId},
		},
	}

	if err := s.ouc.GetHotelRoomType(ctx, p1h); nil != err {
		return nil, err
	}

	t1CreateAt, _ := time.ParseInLocation(conf.FORMAT_YMDHIS, in.CreateAt, time.Local)
	p1c := &biz.CalendarQueue{
		FlowId:               flowId,
		RoomTypeId:           p1h.Arr1RoomType[0].Id,
		SyncType:             int(in.SyncType),
		CreateAt:             t1CreateAt,
		Arr1P1CalendarDetail: make([]*biz.CalendarDetail, len(in.CalendarList)),
	}
	for i := 0; i < len(in.CalendarList); i++ {
		t1Date, _ := time.ParseInLocation(conf.FORMAT_YMD, in.CalendarList[i].Date, time.Local)
		p1c.Arr1P1CalendarDetail[i] = &biz.CalendarDetail{
			Date:  t1Date,
			Num:   int(in.CalendarList[i].Num),
			Price: int(in.CalendarList[i].Price),
		}
	}

	s.ouc.PushCalendar(ctx, p1c)

	return &v1.PushCalendarReply{
		Ota:        in.Ota,
		HotelId:    in.HotelId,
		RoomTypeId: in.RoomTypeId,
		FlowId:     flowId,
	}, nil
}

func (s *OtaService) CalendarJob() {
	fmt.Println("CalendarJob")

	if s.calendarJobStatus {
		return
	}

	go func() {
		defer func() {
			recover()
			s.calendarJobStatus = false
		}()

		s.calendarJobStatus = true

		for {
			// 无限循环
			ctx := context.Background()
			hasQueue := s.ouc.ExecCalendar(ctx)
			if !hasQueue {
				// 如果没有队列就延时 1 秒
				time.Sleep(time.Second)
			}
		}
	}()
}
