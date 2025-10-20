package database

import (
	"fmt"
	"log"
	"os"

	// "github.com/MemoMeto35/Full-Stack-Movie-Streaming-App-in-GO/Server/MagiStreamMoviesServer/database"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func DBInstance() *mongo.Client{
	err:= godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: ubable to load .env file")
	}
	MongoDb:= os.Getenv("MONGODB_URI")
	if( MongoDb == ""){
		log.Fatal("MONGODB_URI not found in environment variables")
	}
	fmt.Println("Connecting to MongoDB at", MongoDb)
	clientOptions:= options.Client().ApplyURI(MongoDb)
	client, err:= mongo.Connect(nil, clientOptions)
	if err != nil {
		return nil
	}
	return client
}
var Client *mongo.Client = DBInstance()
func OpenCollection(collectionName string) *mongo.Collection{
	err :=  godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: ubable to load .env file")
	}
	databaseName:= os.Getenv("DATABASE_NAME")
	fmt.Println("Using database:", databaseName)
	collection:= Client.Database(databaseName).Collection(collectionName)
	if collection == nil {
		return nil
	}
	return collection
}