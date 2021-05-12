package main

import "time"

type Price struct {
	Trade_id int
	Price    string
	Size     string
	Bid      string
	Ask      string
	Volume   string
	Time     time.Time
}

type Coinbase_Historic_Data [][]float64

type Coinbase_Ticker struct {
	TradeID int       `json:"trade_id"`
	Price   string    `json:"price"`
	Size    string    `json:"size"`
	Time    time.Time `json:"time"`
	Bid     string    `json:"bid"`
	Ask     string    `json:"ask"`
	Volume  string    `json:"volume"`
}
