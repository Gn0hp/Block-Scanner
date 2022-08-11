package main

import (
	"BlockScanner/internal/entities"
	"BlockScanner/internal/service"
	database "BlockScanner/internal/service/database/mysql"
	"BlockScanner/internal/service/log"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"logur.dev/logur"
)

type MigrateService struct {
	service.DefaultService
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
		migrateService.logger.Error(fmt.Sprintf("Error while migrating tables: %v", err))
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
		panic(err)
	}

	s.gormDB = gormDb
	s.logger = logger

}
