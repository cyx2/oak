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

type Coinbase_Websocket_Ticket_Payload struct {
	Type      string    `json:"type"`
	Sequence  int64     `json:"sequence"`
	ProductID string    `json:"product_id"`
	Price     string    `json:"price"`
	Open24H   string    `json:"open_24h"`
	Volume24H string    `json:"volume_24h"`
	Low24H    string    `json:"low_24h"`
	High24H   string    `json:"high_24h"`
	Volume30D string    `json:"volume_30d"`
	BestBid   string    `json:"best_bid"`
	BestAsk   string    `json:"best_ask"`
	Side      string    `json:"side"`
	Time      time.Time `json:"time"`
	TradeID   int       `json:"trade_id"`
	LastSize  string    `json:"last_size"`
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
