package main

import (
	"BlockScanner/internal/entities"
	"BlockScanner/internal/services"
	database "BlockScanner/internal/services/database/mysql"
	"BlockScanner/internal/services/log"
	"BlockScanner/internal/services/report/telegram"
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
		msg := fmt.Sprintf("Error while migrating tables: %v", err)
		migrateService.logger.Error(msg)
		telegram.ReportErrorMessageTelegram(msg)
		return
	}
	migrateService.logger.Info("Seed completed")
}

func (s *MigrateService) Init() {
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
		telegram.ReportErrorMessageTelegram(fmt.Sprintf("Error while create new connector to database: %v", err))
		panic(err)
	}

	s.gormDB = gormDb
	s.logger = logger

}
