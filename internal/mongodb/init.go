package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"micro_service/config"
)

var mongoDB *mongo.Client

func InitMongo() {

	clientOptions := options.Client().ApplyURI(config.Conf.GetString("mongo.address"))

	//连接池
	clientOptions.SetMaxPoolSize(config.Conf.GetUint64("mongo.maxPoolSize"))
	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	mongoDB = client
	log.Println("Mongo is Collection!!!")

}
