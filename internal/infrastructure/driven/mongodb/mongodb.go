package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type MongoRepository struct {
	mongoClient   *mongo.Client
	mongoDatabase *mongo.Database
}

func NewMongoConnection(ctx context.Context, mongoUrl string, database string, appName string) *MongoRepository {
	log.Print("Initializing MongoDB connection...")

	clientOptions := options.Client()
	clientOptions.ApplyURI(mongoUrl)
	clientOptions.SetAppName(appName)
	clientOptions.SetCompressors([]string{"zstd", "zlib"})
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("MongoDB connection failed: %s", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Printf("MongoDB ping failed: %s", err)
	}

	databaseOptions := options.Database()

	log.Print("MongoDB is connected")
	return &MongoRepository{
		mongoClient:   client,
		mongoDatabase: client.Database(database, databaseOptions),
	}
}

func (m *MongoRepository) GetDatabase() *mongo.Database {
	return m.mongoDatabase
}

func (m *MongoRepository) DisconnectMongoDB(ctx context.Context) {
	if m.mongoClient == nil {
		log.Fatal("MongoDB client is nil")
		return
	}

	err := m.mongoClient.Disconnect(ctx)
	if err != nil {
		log.Fatalf("MongoDB disconnect failed: %s", err)
	}
	log.Print("MongoDB disconnected")
}
