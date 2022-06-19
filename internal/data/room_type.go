package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm/clause"
	"kratos-ota/internal/biz"
	"math/rand"
	"time"
)

type roomTypeRepo struct {
	data *Data
	log  *log.Helper
}

func NewRoomTypeRepo(data *Data, logger log.Logger) biz.RoomTypeRepo {
	return &roomTypeRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

type RoomType struct {
	Ota          string
	HotelId      string
	RoomTypeId   string
	RoomTypeName string
}

func (RoomType) TableName() string {
	return "room_type"
}

func (r *roomTypeRepo) GetHotelRoomType(ctx context.Context, p1h *biz.Hotel) error {
	p1rt := &RoomType{
		Ota:        p1h.Ota,
		HotelId:    p1h.HotelId,
		RoomTypeId: p1h.Arr1RoomType[0].RoomTypeId,
	}
	r.data.db.Where(p1rt).First(p1rt)
	p1h.Arr1RoomType[0].RoomTypeName = p1rt.RoomTypeName
	return nil
}

func (r *roomTypeRepo) TestCreateRoomTypeMock() {
	rand.Seed(time.Now().Unix())
	arr1ota := []string{"agoda", "expedia", "booking"}
	arr1letters := []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	for i := 0; i < 10; i++ {
		arr1RoomType := make([]RoomType, 10000)
		for j := 0; j < 10000; j++ {
			arr1RoomType[j].Ota = arr1ota[rand.Intn(3)]
			arr1RoomType[j].HotelId = fmt.Sprintf("%d", rand.Intn(10000)+10000)
			arr1RoomType[j].RoomTypeId = fmt.Sprintf("%d", rand.Intn(1000000)+1000000)

			b := make([]rune, 20)
			for i := range b {
				b[i] = arr1letters[rand.Intn(62)]
			}
			arr1RoomType[j].RoomTypeName = string(b)
		}
		tx := r.data.db.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(arr1RoomType, len(arr1RoomType))
		if tx.Error != nil {
			fmt.Println(tx.Error)
		}
	}
}
