package main

import (
	"BlockScanner/internal/services"
	database "BlockScanner/internal/services/database/mysql"
	"BlockScanner/internal/services/log"
	"BlockScanner/internal/services/report/telegram"
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"logur.dev/logur"
)

type BlockScanner struct {
	services.DefaultService
	Context context.Context
	logger  logur.LoggerFacade
	gormDB  *gorm.DB
}

func (s *BlockScanner) Init() {
	s.DefaultService.Init()
	var (
		logCf = log.Config{}
		dbCf  = database.Config{}
	)
	cfBytes, _ := json.Marshal(viper.GetStringMap("log"))
	json.Unmarshal(cfBytes, &logCf)
	cfBytes, _ = json.Marshal(viper.GetStringMap("database"))
	json.Unmarshal(cfBytes, &dbCf)

	logger := log.NewLogger(logCf)
	log.SetStandaloneLogger(logger)

	if dbCf.Params == nil {
		dbCf.Params = make(map[string]string)
	}
	gormDb, err := database.NewConnector(dbCf)
	if err != nil {
		msg := fmt.Sprintf("Error while connecting to database: %v", err)
		logrus.Error(msg)
		telegram.ReportErrorMessageTelegram(msg)
		panic(err)
	}

	s.logger = logger
	s.gormDB = gormDb
}

func ConnectDB() *gorm.DB {
	service := BlockScanner{
		Context: context.Background(),
	}
	service.Init()
	return service.gormDB
}
