package main

import (
	"fmt"
	"time"
)

func load_historic_data(minutes int) {
	later_bound := time.Now().Unix()

	// 60 * minutes, because Coinbase API returns max 300 length array
	// and Unix Epoch time is in seconds
	early_bound := later_bound - (60 * int64(minutes))

	later_bound_iso := time.Unix(later_bound, 0)
	early_bound_iso := time.Unix(early_bound, 0)

	historic_data_ptr := get_historic_data(early_bound_iso, later_bound_iso)
	historic_data := *historic_data_ptr

	fmt.Printf("INFO: Number of historic records is %v\n", len(historic_data))

	for i := range historic_data {
		// price_string := strconv.FormatFloat(historic_data[i][4], 'E', -1, 64)

		price_string := fmt.Sprintf("%v", historic_data[i][4])

		// fmt.Printf("Price in this object is %v\n", price_string)

		load_price := Coinbase_Ticker{
			Trade_id: 0,
			Price:    price_string,
			Size:     "",
			Time:     time.Unix(int64(historic_data[i][0]), 0),
			Bid:      "",
			Ask:      "",
			Volume:   "",
		}

		fmt.Printf("Price in the data structure is %s\n", load_price.Price)

		insert_price(load_price)
	}
}
