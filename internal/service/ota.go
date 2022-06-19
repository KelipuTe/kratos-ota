package service

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	v1 "kratos-ota/api/ota/v1"
	"kratos-ota/internal/biz"
)

type OtaService struct {
	v1.UnimplementedOtaServer
	ouc *biz.OtaUsecase
}

// NewOtaService
// api 层 --> service 层 --> biz 层
// service 层应该把 api 层的数据结构转换成 biz 层的数据结构
func NewOtaService(ouc *biz.OtaUsecase) *OtaService {
	return &OtaService{ouc: ouc}
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
	err := s.ouc.GetHotelRoomType(ctx, p1h)

	if err != nil {
		return nil, errors.New(500, err.Error(), err.Error())
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

func (s *OtaService) CalendarJob()  {
	fmt.Println("CalendarJob")
}
