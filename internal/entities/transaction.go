package entities

type Transaction struct {
	DefaultModel
	TransactionHash string  `json:"transaction_hash"`
	Status          string  `json:"status"`
	BlockNum        uint32  `json:"block_num"`
	TimeStamp       uint64  `json:"timestamp"`
	From            string  `json:"from"`
	To              string  `json:"to"`
	Value           float64 `json:"value"`
	TransactionFee  float64 `json:"transaction_fee"`
}
