package repository

import (
	"github.com/rfaelmarini/call-cat-api/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ResponseRepository interface {
	Save(response entity.Response)
	Find(url string) entity.Response
}

type database struct {
	connection *gorm.DB
}

func NewResponseRepository() ResponseRepository {
	dsn := "root@tcp(127.0.0.1:3306)/callcatapidb"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	db.AutoMigrate(&entity.Response{})
	return &database{
		connection: db,
	}
}

func (db *database) Save(response entity.Response) {
	db.connection.Create(&response)
}

func (db *database) Find(url string) entity.Response {
	response := entity.Response{}
	db.connection.Where("requested_url = ?", url).First(&response)
	return response
}
