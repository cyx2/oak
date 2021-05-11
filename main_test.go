package main

import (
	"log"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func Test_Price_Insert(t *testing.T) {
	initialize_config()
	initialize_db()

	test_document := Price{
		Trade_id: 123456789,
		Price:    "111.11",
		Size:     "222.22",
		Bid:      "333.33",
		Ask:      "444.44",
		Volume:   "555.55",
		Time:     time.Now(),
	}

	insert_price(test_document)
}

var historic_data Historic_data

func Test_Get_Historic_Data_Single(t *testing.T) {
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

	collection.Drop(ctx)
}

func Test_Get_Historic_Data_Batch(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		// Heroku maintains config once deployed, so deferring to this config in lieu
		// of a .env file
		log.Println("TEST: Failed to pull .env config")
	}

	load_historic_data(1)
}

func Test_Websocket_Connection(t *testing.T) {
	initialize_config()

	connect_to_websocket()
}
