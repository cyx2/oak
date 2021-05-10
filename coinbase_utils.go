package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Historic_data [][]float32

func get_historic_data(start time.Time, end time.Time) *Historic_data {
	// Construct the HTTP request conformant to Coinbase API
	// 1) URL, 2) Params

	start_iso := start.Format(time.RFC3339)
	end_iso := end.Format(time.RFC3339)

	cxn_string := os.Getenv("CB_URL_ROOT") + os.Getenv("CB_URL_HISTORIC") + "?start=" + start_iso + "&end=" + end_iso

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
