package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DB_NAME     = "go-do"
	MONGODB_URI = "mongodb://localhost:27017"
)

type MongoConnection struct {
	client *mongo.Client
	ctx    context.Context
}

func New() *MongoConnection {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MONGODB_URI))
	if err != nil {
		panic(err)
	}

	return &MongoConnection{
		ctx:    context.TODO(),
		client: client,
	}
}

func (m *MongoConnection) Connection(collection string) *mongo.Collection {
	return m.client.Database(DB_NAME).Collection(collection)
}

func (m *MongoConnection) Disconnect() {
	if m.client != nil {
		if err := m.client.Disconnect(m.ctx); err != nil {
			panic(err)
		}
	}
}
