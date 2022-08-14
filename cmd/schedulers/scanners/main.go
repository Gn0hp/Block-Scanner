package main

import (
	"BlockScanner/internal/services/report/telegram"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

func main() {
	interval := time.Second * time.Duration(viper.GetInt("ticker.interval"))
	if interval == 0 {
		interval = time.Second * 5
	}
	ticker := time.NewTicker(interval)
	logrus.Infof("New ticker initiated with interval: %d. Start scanning ", interval)

	success := true
	for {
		select {
		case <-ticker.C:
			if !success {
				logrus.Info("Nothing to update. Waiting for next scan ...")
				time.Sleep(interval)
				continue
			}
			success = false
			go func() {
				db := ConnectDB()
				gormDb, err := db.DB()
				if err != nil {
					msg := fmt.Sprintf("Error while connecting to database at scan channel: %v", err)
					logrus.Error(msg)
					telegram.ReportErrorMessageTelegram(msg)
				}
				defer gormDb.Close()
				netType := viper.GetString("block_scanner.network")
				scanning(db, netType)
				success = true
			}()
		}
	}
}
