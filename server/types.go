package main

import "time"

type Coinbase_Ticker struct {
	Trade_id int       `json:"trade_id"`
	Price    string    `json:"price"`
	Size     string    `json:"size"`
	Time     time.Time `json:"time"`
	Bid      string    `json:"bid"`
	Ask      string    `json:"ask"`
	Volume   string    `json:"volume"`
}
type Coinbase_Historic_Data [][]float64
