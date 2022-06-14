package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "kratos-learning/api/ota/v1"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

type Hotel struct {
	Ota          string
	HotelId      string
	Arr1RoomType []RoomType
}

type RoomType struct {
	RoomTypeId   string
	RoomTypeName string
}

type RoomTypeRepo interface {
	GetHotelRoomType(ctx context.Context, p1h *Hotel) error
}

type CalendarRepo interface {
}

type OtaUsecase struct {
	rtr RoomTypeRepo
	cr  CalendarRepo
	log *log.Helper
}

func NewOtaUsecase(rtr RoomTypeRepo, cr CalendarRepo, logger log.Logger) *OtaUsecase {
	return &OtaUsecase{
		rtr: rtr,
		cr:  cr,
		log: log.NewHelper(logger),
	}
}

func (ouc *OtaUsecase) GetHotelRoomType(ctx context.Context, p1h *Hotel) error {
	err := ouc.rtr.GetHotelRoomType(ctx, p1h)

	if err != nil {

	}

	return nil
}
