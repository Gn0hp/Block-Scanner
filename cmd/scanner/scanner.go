package main

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
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
	Sort string //asc or desc
}

func scanning(network_type string) {
	netURL := viper.GetString(fmt.Sprintf("bsc_scan_api.url_%s", network_type))

	logrus.Info("url request is: ", filterURL(netURL))

	resp, err := http.Get(netURL)
	if err != nil {
		logrus.Errorf("Error retrieving data: %v", err)
	}
	logrus.Infof("data is: %v", resp)
	//req := filterURL(netURL)
}

/* https://api.bscscan.com/api
   ?module=logs
   &action=getLogs
   &fromBlock=4993830
   &toBlock=4993832
   &address=0xe561479bebee0e606c19bb1973fc4761613e3c42
   &topic0=0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef
   &apikey=YourApiKeyToken*/
func filterURL(url string) string {
	var query = ""
	var params ScannerParams
	queryCfgBytes, _ := json.Marshal(viper.GetStringMap("api_bsc_params"))
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
