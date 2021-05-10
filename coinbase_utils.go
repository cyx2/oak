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

type Historic_data [][]float64

func get_historic_data(start_epoch time.Time, end_epoch time.Time) *Historic_data {
	// Construct the HTTP request conformant to Coinbase API
	// 1) URL, 2) Params

	start_iso := start_epoch.Format(time.RFC3339)
	end_iso := end_epoch.Format(time.RFC3339)

	cxn_string := os.Getenv("CB_URL_ROOT") + os.Getenv("CB_URL_HISTORIC") + "?start=" + start_iso + "&end=" + end_iso

	fmt.Printf("INFO: Cxn string is: %s\n", cxn_string)

	resp, err := http.Get(cxn_string)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var historic_data Historic_data
	json.Unmarshal(body, &historic_data)
	return &historic_data
}
