package entities

import "time"

type Transaction struct {
	DefaultModel
	TransactionHash string    `json:"transaction_hash"`
	Status          string    `json:"status"`
	BlockNum        uint32    `json:"block_num"`
	TimeStamp       time.Time `json:"time_stamp"`
	From            string    `json:"from"`
	To              string    `json:"to"`
	Value           float64   `json:"value"`
	TransactionFee  float64   `json:"transaction_fee"`
}
