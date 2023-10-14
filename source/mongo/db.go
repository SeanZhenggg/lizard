package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type IMongoCli interface {
	GetCollection(ctx context.Context, name string) ICollection
}

type mongoDB struct {
	DB *mongo.Database
}

func ProvideMongoDbCli() IMongoCli {
	newDB := &mongoDB{}
	newDB.DB = mongoDbConnect()

	return newDB
}

func mongoDbConnect() *mongo.Database {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// 建立连接
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("🔔🔔🔔 MONGODB CONNECT ERROR: ", err.Error(), " 🔔🔔🔔")
	}

	// 检查连接是否正常
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("🔔🔔🔔 MONGODB CONNECT ERROR: ", err.Error(), " 🔔🔔🔔")
	}

	return client.Database("lizard")
}

func (m *mongoDB) GetCollection(ctx context.Context, name string) ICollection {
	c := &Collection{
		m.DB.Collection(name),
	}

	return c
}
