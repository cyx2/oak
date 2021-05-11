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

type Coinbase_Websocket_Subscribe struct {
	Type        string                                 `json:"type"`
	Product_ids []string                               `json:"product_ids"`
	Channels    []Coinbase_Websocket_Channel_Subscribe `json:"channels"`
}

type Coinbase_Websocket_Channel_Subscribe struct {
	Name        string   `json:"name"`
	Product_ids []string `json:"product_ids"`
}
