package database

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DB_NAME     = "go-do"
	MONGODB_URI = os.Getenv("MONGO_URI")
)

type MongoConnection struct {
	Ctx context.Context

	client *mongo.Client
}

func New() *MongoConnection {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MONGODB_URI))
	if err != nil {
		panic(err)
	}

	return &MongoConnection{
		Ctx:    context.TODO(),
		client: client,
	}
}

func (m *MongoConnection) Connection(collection string) *mongo.Collection {
	return m.client.Database(DB_NAME).Collection(collection)
}

func (m *MongoConnection) Disconnect() {
	if m.client != nil {
		if err := m.client.Disconnect(m.Ctx); err != nil {
			panic(err)
		}
	}
}
