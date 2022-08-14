package main

import (
	"BlockScanner/internal/entities"
	"BlockScanner/internal/services/report/telegram"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"time"
)

type ScannerParams struct {
	Module    string
	Action    string
	BlockNo   uint //wait for testing to know whether it's type is uint32 or uint64
	ApiKey    string
	TxHash    string
	TimeStamp uint64 //the integer representing the Unix timestamp in seconds.
	Closest   string //the closest available block to the provided timestamp, either before or after
	//StartDate time.Time //format yyyy-MM-dd
	//EndDate   time.Time		// not use because of err while parsing
	FromBlock uint
	ToBlock   uint
	Address   string
	Sort      string //asc or desc
}

func scanning(db *gorm.DB, network string) {
	logrus.Info("Scanning ...")

	var transactionsArr []*entities.Transaction

	//TODO : scan and parse transactions from bscscsan.com
	var rpcUrl string
	if network == "mainnet" {
		rpcUrl = viper.GetString("api.mainnetUrl")
	} else {
		rpcUrl = viper.GetString("api.testnetUrl")
	}

	req := filterParams(rpcUrl)
	logrus.Info("Request: ", req)

	res, err := http.Get(req)
	if err != nil {
		logrus.Errorf("Error while getting response from bscscan.com: %v", err)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	logrus.Info("Response: ", string(body))
	return
	result := db.CreateInBatches(transactionsArr, 20)
	if result.Error != nil {
		msg := fmt.Sprintf("Error while save to db in batches: %v", result.Error)
		logrus.Error(msg)
		telegram.ReportErrorMessageTelegram(msg)
	}
	if result.RowsAffected > 0 {
		logrus.Infof("%d rows affected", result.RowsAffected)
	}
	logrus.Info("Scanning finished at: ", time.Now())

}

/* https://api.bscscan.com/api
   ?module=logs
   &action=getLogs
   &fromBlock=
   &toBlock=
   &address=
   &topic0=
   &apikey=YourApiKeyToken*/
func filterParams(url string) string {
	var query = ""
	var params ScannerParams
	queryCfgBytes, _ := json.Marshal(viper.GetStringMap("scanner_params"))
	err := json.Unmarshal(queryCfgBytes, &params)
	if err != nil {
		logrus.Errorf("Error while unmarshalling api_bsc_params: %v", err)
	}
	if params.Module != "" {
		query = fmt.Sprintf("module=%s", params.Module)
	}
	if params.Action != "" {
		query = fmt.Sprintf("%s&action=%s", query, params.Action)
	}
	if params.BlockNo != 0 {
		query = fmt.Sprintf("%s&blockno=%d", query, params.BlockNo)
	}
	if params.TxHash != "" {
		query = fmt.Sprintf("%s&txhash=%s", query, params.TxHash)
	}
	if params.TimeStamp != 0 {
		params.TimeStamp = uint64(time.Now().Unix())
		query = fmt.Sprintf("%s&timestamp=%d", query, params.TimeStamp)
	}
	if params.Closest != "" {
		query = fmt.Sprintf("%s&closest=%s", query, params.Closest)
	}
	if params.FromBlock != 0 {
		query = fmt.Sprintf("%s&fromBlock=%d", query, params.FromBlock)
	}
	if params.ToBlock != 0 {
		query = fmt.Sprintf("%s&toBlock=%d", query, params.ToBlock)
	}
	//if params.StartDate != nil {
	//	query = fmt.Sprintf("module=%s", params.Module)
	//}
	//if params.EndDate != "" {
	//	query = fmt.Sprintf("module=%s", params.Module)
	//}
	if params.Sort != "" {
		query = fmt.Sprintf("&sort=%s", params.Sort)
	}
	query = query + "&apikey=" + params.ApiKey
	return fmt.Sprintf("%s?%s", url, query)
}
