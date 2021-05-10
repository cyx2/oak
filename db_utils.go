package main

import (
	"log"
)

func insert_price(newPrice Price) {

	insertResult, err := collection.InsertOne(ctx, newPrice)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("DB: Inserted document with id: %s", insertResult.InsertedID)
}
