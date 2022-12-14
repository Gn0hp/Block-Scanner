package main

import (
	"block_scanner/internal/entities"
	"block_scanner/internal/services"
	database "block_scanner/internal/services/database/mysql"
	"block_scanner/internal/services/log"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"logur.dev/logur"
)

type MigrateService struct {
	services.DefaultService
	logger logur.LoggerFacade
	gormDB *gorm.DB
}

func main() {
	migrateService := MigrateService{}
	migrateService.Init()

	tables := []interface{}{
		entities.Transaction{},
	}

	err := migrateService.gormDB.AutoMigrate(tables...)
	if err != nil {
		migrateService.logger.Error(fmt.Sprintf("Sedd failed, detail: %v", err))
		return
	}
	migrateService.logger.Info("Seed complete")
}

func (m *MigrateService) Init() {
	m.DefaultService.Init()
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
	m.gormDB = gormDB
	m.logger = logger
}
