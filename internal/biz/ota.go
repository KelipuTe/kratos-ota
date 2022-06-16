package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "kratos-learning/api/ota/v1"
	"kratos-learning/internal/conf"
	"kratos-learning/internal/pkg/middleware/auth"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

type Token struct {
	ClientId     string
	ClientSecret string
	Token        string
}

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
	rtr   RoomTypeRepo
	cr    CalendarRepo
	cauth *conf.Auth
	log   *log.Helper
}

func NewOtaUsecase(rtr RoomTypeRepo, cr CalendarRepo, cauth *conf.Auth, logger log.Logger) *OtaUsecase {
	return &OtaUsecase{
		rtr:   rtr,
		cr:    cr,
		cauth: cauth,
		log:   log.NewHelper(logger),
	}
}

func (ouc *OtaUsecase) GetToken(ctx context.Context, p1t *Token) error {
	p1t.Token = auth.MakeJWT(ouc.cauth.Jwt.SecretKey, p1t.ClientId)

	return nil
}

func (ouc *OtaUsecase) GetHotelRoomType(ctx context.Context, p1h *Hotel) error {
	err := ouc.rtr.GetHotelRoomType(ctx, p1h)
	if err != nil {
	}

	return nil
}
