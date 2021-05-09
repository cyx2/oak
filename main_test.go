package main

import (
	"testing"
	"time"
)

// func TestPriceInsert(t *testing.T) {
// 	initializeConfig()
// 	initializeDb()

// 	testDocument := Price{
// 		Trade_id: 123456789,
// 		Price:    "111.11",
// 		Size:     "222.22",
// 		Bid:      "333.33",
// 		Ask:      "444.44",
// 		Volume:   "555.55",
// 		Time:     time.Now(),
// 	}

// 	insertPrice(testDocument)
// }

// func TestTimeTemporary(t *testing.T) {
// 	epoch := int64(1620457200)
// 	epoch_time := time.Unix(epoch, 0)
// 	iso_time := epoch_time.Format(time.RFC3339)

// 	fmt.Printf("Calculated time to be %s", iso_time)
// }

func TestGetHistoricData(t *testing.T) {
	main()

	start_epoch := time.Unix(int64(1620457200), 0)
	end_epoch := time.Unix(int64(1620457200), 0)

	get_historic_data(start_epoch, end_epoch)
}
