package biz

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/sync/errgroup"
	v1 "kratos-ota/api/ota/v1"
	"kratos-ota/internal/conf"
	"kratos-ota/internal/pkg/middleware/auth"
	"time"
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
	Id           int
	RoomTypeId   string
	RoomTypeName string
}

type CalendarQueue struct {
	FlowId               string
	RoomTypeId           int
	SyncType             int
	CreateAt             time.Time
	Arr1P1CalendarDetail []*CalendarDetail
}

type CalendarDetail struct {
	Date  time.Time
	Num   int
	Price int
}

type RoomTypeRepo interface {
	GetHotelRoomType(ctx context.Context, p1h *Hotel) error
}

type CalendarRepo interface {
	SaveCalendarDetail(ctx context.Context, p1cq *CalendarQueue) error
	SaveCalendarQueue(ctx context.Context, p1cq *CalendarQueue) error
	RPUSHCalendarQueue(ctx context.Context, id int) error
	LPOPCalendarQueue(ctx context.Context) (int, error)
	ListCalendarQueueByRoomTypeId(ctx context.Context, id int) ([]*CalendarQueue, error)
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
	if err := ouc.rtr.GetHotelRoomType(ctx, p1h); err != nil {
		return err
	}

	return nil
}

func (ouc *OtaUsecase) PushCalendar(ctx context.Context, p1cq *CalendarQueue) error {
	if err := ouc.cr.SaveCalendarDetail(ctx, p1cq); err != nil {
	}

	if err := ouc.cr.SaveCalendarQueue(ctx, p1cq); err != nil {
	}

	if err := ouc.cr.RPUSHCalendarQueue(ctx, p1cq.RoomTypeId); err != nil {
	}

	return nil
}

func (ouc *OtaUsecase) ExecCalendar(ctx context.Context) bool {
	id, err := ouc.cr.LPOPCalendarQueue(ctx)
	if nil != err {
		return false
	}

	arr1p1c, _ := ouc.cr.ListCalendarQueueByRoomTypeId(ctx, id)
	spew.Dump(arr1p1c)

	eg := &errgroup.Group{}

	for i := 0; i < len(arr1p1c); i++ {
		eg.Go(func() error {
			spew.Dump(arr1p1c[i])
			return nil
		})
	}

	if err := eg.Wait(); nil != err {
	}

	return false
}
