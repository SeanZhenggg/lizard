package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"lizard/source/config"
	"log"
)

type IMongoCli interface {
	GetCollection(ctx context.Context, name string) ICollection
}

type mongoDB struct {
	DB       *mongo.Database
	DbConfig config.DbConfig
}

func ProvideMongoDbCli(config config.IConfigEnv) IMongoCli {
	dbCfg := config.GetDbConfig()
	log.Printf("dbCfg : %v", dbCfg)
	newDB := &mongoDB{
		DbConfig: dbCfg,
		DB:       dbConnect(dbCfg),
	}

	return newDB
}

func dbConnect(dbConfig config.DbConfig) *mongo.Database {
	credential := options.Credential{
		Username: "root",
		Password: "mongopw",
	}
	log.Printf("Connecting to %s", fmt.Sprintf(
		"mongodb://%s:%s",
		dbConfig.Host,
		dbConfig.Port,
	))
	clientOptions := options.Client().ApplyURI(fmt.Sprintf(
		"mongodb://%s:%s",
		dbConfig.Host,
		dbConfig.Port,
	)).
		SetAuth(credential).
		SetMaxPoolSize(20).
		SetWriteConcern(writeconcern.New(writeconcern.WMajority()))

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

	return client.Database(dbConfig.DbName)
}

func (m *mongoDB) GetCollection(ctx context.Context, name string) ICollection {
	c := &Collection{
		m.DB.Collection(name),
	}

	return c
}
