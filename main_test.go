package main

import (
	"testing"
	"time"
)

func TestPriceInsert(t *testing.T) {
	initializeConfig()
	initializeDb()

	testDocument := Price{
		Trade_id: 123456789,
		Price:    "111.11",
		Size:     "222.22",
		Bid:      "333.33",
		Ask:      "444.44",
		Volume:   "555.55",
		Time:     time.Now(),
	}

	insertPrice(testDocument)
}
