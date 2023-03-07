package database

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/yaml.v3"

	models "gin-mongo/src/models"
)

func ConnectDB() *mongo.Client {
	yamlFile, _ := os.Open("config.yaml")
	defer yamlFile.Close()

	byteValue, _ := ioutil.ReadAll(yamlFile)

	var config models.Config

	yaml.Unmarshal(byteValue, &config)
	// log.Println(config)

	client, err := mongo.NewClient(options.Client().ApplyURI(config.Url))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")
	return client
}

var Client *mongo.Client = ConnectDB()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("test").Collection(collectionName)
	return collection
}
