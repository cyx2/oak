package main

import (
	"log"
)

func insert_price(newPrice Price) {

	insert_result, err := collection.InsertOne(ctx, newPrice)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("DB: Inserted document with id: %s", insert_result.InsertedID)
}
