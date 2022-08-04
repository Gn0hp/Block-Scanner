package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"time"
)

func connectDB() *gorm.DB {
	service := BlockScanner{
		Context: context.Background(),
	}
	service.Init()
	return service.GormDB
}
func main() {
	duration := time.Second * time.Duration(viper.GetInt("scanner_ticker.ticker"))
	if duration == 0 {
		duration = time.Second * 5
	}
	ticker := time.NewTicker(duration)
	logrus.Info("Starting scanner")
	success := true
	for {
		select {
		case <-ticker.C:
			if !success {
				logrus.Info("Nothing new to update. Waiting for the next scan")
				time.Sleep(duration)
				continue
			}
			success = false
			go func() {
				db := connectDB()
				gormDB, err := db.DB()
				if err != nil {
					logrus.Errorf("Error with DB in scanner channel, detail: %v", err)
				}
				defer gormDB.Close()
				scanning("testnet")
				success = true
			}()
		}
	}
}
