package data

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestNewDB(t *testing.T) {
	db, _ := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/kratos-learning?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	roomTypeRepo := &roomTypeRepo{
		data: &Data{db: db},
		log:  nil,
	}
	roomTypeRepo.TestCreateRoomTypeMock()
}
