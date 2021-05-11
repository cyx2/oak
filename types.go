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
