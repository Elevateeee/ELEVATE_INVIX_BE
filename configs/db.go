package configs

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



var DB *mongo.Client


func ConnectDB()  {
	urlDb:= os.Getenv("MONGO_URI")
	if urlDb == "" {
		log.Fatal("URL Mongo belom di set di .env")
	}

	clientOptions := options.Client().ApplyURI(urlDb)
	contx, cencel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cencel()

	client, err := mongo.Connect(contx, clientOptions)
	if err != nil {
		log.Fatal("Gagal connect ke MongoDB:", err)
	}

	err = client.Ping(contx, nil)
	if err != nil {
		log.Fatal("Gagal ping MongoDB:", err)
	}
	log.Println("Berhasil connect ke MongoDB")
	DB = client
}

func GetCollection(collectionName string) *mongo.Collection {
	databaseName:= "Invix"
	return DB.Database(databaseName).Collection(collectionName)
}
