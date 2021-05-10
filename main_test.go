package main

import (
	"log"
	"testing"
	"time"

	"github.com/joho/godotenv"
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

var historic_data Historic_data

func TestGetHistoricDataSingle(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		// Heroku maintains config once deployed, so deferring to this config in lieu
		// of a .env file
		log.Println("TEST: Failed to pull .env config")
	}

	start_epoch := time.Unix(int64(1620457200), 0)
	end_epoch := time.Unix(int64(1620457200), 0)

	historic_data_ptr := get_historic_data(start_epoch, end_epoch)
	historic_data = *historic_data_ptr

	if historic_data[0][1] != 57805.87 {
		t.Errorf("Returned value does not match expected value, got %v want %v", historic_data[0][1], 57805.87)
	}
}

func TestGetHistoricDataBatch(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		// Heroku maintains config once deployed, so deferring to this config in lieu
		// of a .env file
		log.Println("TEST: Failed to pull .env config")
	}

	load_historic_data(5)
}
