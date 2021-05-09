package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db         string
	collection *mongo.Collection
	ctx        context.Context
)

type price struct {
	trade_id int
	price    float32
	size     float32
	bid      float32
	ask      float32
	volume   float32
	time     time.Time
}

func initializeConfig() {
	// Utilize godotenv to load config from .env file on local machine
	err := godotenv.Load()
	if err != nil {
		// Heroku maintains config once deployed, so deferring to this config in lieu
		// of a .env file
		log.Println("INIT: No .env file found in repo, deferring to system config")
	}
	log.Printf("INIT: Loaded the %s config set", os.Getenv("ENV"))
}

func initializeDb() {
	// Determine the Atlas DB based on the ENV config variable
	switch os.Getenv("ENV") {
	case "PROD":
		db = "oak_prod"
	default:
		db = "oak_dev"
	}

	log.Printf("INIT: Writing to the %s db", db)

	// MongoDB Atlas connection params and string computed based on the environment
	cxnParams := "/?retryWrites=true&w=majority"
	dbCxnString := "mongodb+srv://" + os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@" + os.Getenv("DB_URL") + "/" + db + cxnParams

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbCxnString))
	if err != nil {
		log.Fatal(err)
	}

	// Set the global collection variable for all dbs
	collection = client.Database(db).Collection("prices")
	log.Printf("INFO: Connected to the Atlas cluster %s", os.Getenv("DB_URL"))
}

func main() {
	initializeConfig()
	initializeDb()
}
