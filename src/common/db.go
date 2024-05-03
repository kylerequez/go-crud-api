package common

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func ConnectDB() error {
	log.Println(":::-::: Connecting to DB...")
	var dbHost, dbPort, dbString = os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_STRING")

	var dbURI string = fmt.Sprintf(dbString, dbHost, dbPort)

	clientOpts := options.Client().ApplyURI(dbURI)
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		return errors.New(err.Error())
	}

	DB = client
	log.Println(":::-::: Successfully Connected to DB")
	return nil
}

func CloseDB() error {
	log.Println(":::-::: Closed DB")
	return DB.Disconnect(context.TODO())
}

func GetDB(database string) *mongo.Database {
	return DB.Database(database)
}
