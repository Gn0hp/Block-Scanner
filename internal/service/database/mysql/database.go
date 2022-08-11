package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ClientInterface interface {
	Ping() error
	Connect(connectString string) error
}

func NewConnector(config Config) (*gorm.DB, error) {
	config.Params["parseTime"] = "true"
	config.Params["rejectReadOnly"] = "true"

	db, err := gorm.Open(mysql.Open(config.GetDriverSourceName()), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Connect database failed, detail: %v", err))
	}
	return db, err
}
