package mysql

import (
	"BlockScanner/internal/services/report/telegram"
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
		msg := fmt.Sprintf("Connect database failed, detail: %v", err)
		telegram.ReportErrorMessageTelegram(msg)
		panic(err)
	}
	return db, err
}
