package main

import (
	"log"
)

func insertPrice(newPrice Price) {

	insertResult, err := collection.InsertOne(ctx, newPrice)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("DB: Inserted document with id: %s", insertResult.InsertedID)
}
