package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ClientInterface interface {
	Ping() error
	Connect(connectingString string) error
}

func NewConnector(config Config) (*gorm.DB, error) {
	//Set mandatory params
	config.Params["parseTime"] = "true"
	config.Params["rejectReadOnly"] = "true"

	db, err := gorm.Open(mysql.Open(config.GetDriverSourceName()), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("connect database failed, detail: %v", err))
	}
	return db, nil
}
