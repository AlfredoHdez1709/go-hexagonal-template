package mongodb

import (
	"context"
	"go-hexagonal-template/internal/infrastructure/driven/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	mongoClient   *mongo.Client
	mongoDatabase *mongo.Database
}

func NewMongoConnection(ctx context.Context, mongoUrl string, database string, appName string) *MongoRepository {
	logger.Logger.Debug("Initializing MongoDB connection...")

	clientOptions := options.Client()
	clientOptions.ApplyURI(mongoUrl)
	clientOptions.SetAppName(appName)
	clientOptions.SetCompressors([]string{"zstd", "zlib"})
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		logger.Logger.Fatalf("MongoDB connection failed: %s", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		logger.Logger.Fatalf("MongoDB ping failed: %s", err)
	}

	databaseOptions := options.Database()

	logger.Logger.Debug("MongoDB is connected")
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
		logger.Logger.Fatalw("MongoDB client is nil")
		return
	}

	err := m.mongoClient.Disconnect(ctx)
	if err != nil {
		logger.Logger.Fatalf("MongoDB disconnect failed: %s", err)
	}
	logger.Logger.Info("MongoDB disconnected")
}
