package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
	"log"
	"micro_service/config"
	"sync"
	"time"
)

type M = bson.M

var mongoDB *mongo.Client

var once sync.Once

func InitMongo() {

	once.Do(func() {
		clientOptions := options.Client().ApplyURI(config.Conf.GetString("mongo.address"))
		// Connect to MongoDB
		client, err := mongo.Connect(context.Background(), clientOptions)
		if err != nil {
			log.Fatal(err)
		}
		err = client.Ping(context.Background(), nil)
		if err != nil {
			log.Fatal(err)
		}
		mongoDB = client
		log.Println("Mongo is Collection!!!")
	})

}

type mongoCollection struct {
	Timeout    time.Duration
	Collection *mongo.Collection
}

func getCollection(database string, collection string, opts ...*options.CollectionOptions) *mongo.Collection {
	return mongoDB.Database(database).Collection(collection, opts...)
}

func MongoDBCurd(database string, collection string, opts ...*options.CollectionOptions) *mongoCollection {
	dbCollection := getCollection(database, collection, opts...)
	return &mongoCollection{Timeout: 5, Collection: dbCollection}
}

func (m *mongoCollection) FindOne(filter interface{}, result interface{}, opts ...*options.FindOneOptions) (err error) {
	err = m.Collection.FindOne(context.Background(), filter, opts...).Decode(result)
	return
}
func (m *mongoCollection) Find(filter interface{}, result interface{}, opts ...*options.FindOptions) (err error) {
	find, err := m.Collection.Find(context.Background(), filter, opts...)
	if err != nil {
		return err
	}
	return find.All(context.Background(), result)
}
func (m *mongoCollection) Insert(data []interface{}, opts ...*options.InsertManyOptions) (err error) {
	_, err = m.Collection.InsertMany(context.Background(), data, opts...)
	return
}
func (m *mongoCollection) InsertOne(data interface{}, opts ...*options.InsertOneOptions) (err error) {
	_, err = m.Collection.InsertOne(context.Background(), data, opts...)
	return
}
func (m *mongoCollection) Count(filter interface{}, opts ...*options.CountOptions) (count int64, err error) {
	return m.Collection.CountDocuments(context.Background(), filter, opts...)
}
