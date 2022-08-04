package main

import (
	"block_scanner/internal/services"
	database "block_scanner/internal/services/database/mysql"
	"block_scanner/internal/services/log"
	"context"
	"encoding/json"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"logur.dev/logur"
)

type BlockScanner struct {
	services.DefaultService
	Context context.Context
	Logger  logur.LoggerFacade
	GormDB  *gorm.DB
}

func (bs *BlockScanner) Init() {
	bs.DefaultService.Init()
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
	gormDB, err := database.NewConnector(dbCf)
	if err != nil {
		panic(err)
	}
	bs.Logger = logger
	bs.GormDB = gormDB
}
