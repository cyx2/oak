package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func get_historic_data(start_epoch time.Time, end_epoch time.Time) *Coinbase_Historic_Data {
	var historic_data Coinbase_Historic_Data

	start_iso := start_epoch.Format(time.RFC3339)
	end_iso := end_epoch.Format(time.RFC3339)

	cxn_string := os.Getenv("CB_URL_ROOT") + os.Getenv("CB_URL_HISTORIC") + "?start=" + start_iso + "&end=" + end_iso
	fmt.Printf("INFO: Historic data cxn string is: %s\n", cxn_string)

	resp, err := http.Get(cxn_string)
	if err != nil {
		log.Println("WARN: Coinbase API read failed")
		log.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(body, &historic_data)
	return &historic_data
}

func get_coinbase_ticker() *Coinbase_Ticker {
	var ticker_data Coinbase_Ticker

	cxn_string := os.Getenv("CB_URL_ROOT") + os.Getenv("CB_URL_TICKER")
	fmt.Printf("INFO: Ticker cxn string is: %s\n", cxn_string)

	resp, err := http.Get(cxn_string)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(body, &ticker_data)
	fmt.Printf("%+v\n", ticker_data)
	return &ticker_data
}
