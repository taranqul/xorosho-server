package worker

import (
	"context"
	"fmt"
	"worker-manager-service/internal/domain/dto"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type MongoRepository struct {
	collection *mongo.Collection
	logger     *zap.Logger
	ctx        context.Context
}

func NewMongoRepository(ctx context.Context, logger *zap.Logger, mongoURI string, dbName string) (*MongoRepository, error) {

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, fmt.Errorf("mongo connect: %w", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("mongo ping: %w", err)
	}

	logger.Info("Successfully connected to MongoDB")

	coll := client.Database(dbName).Collection(dbName)

	return &MongoRepository{
		collection: coll,
		logger:     logger,
		ctx:        ctx,
	}, nil
}

func (r *MongoRepository) Create(worker *dto.WorkerRegister) error {

	_, err := r.collection.InsertOne(r.ctx, worker)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return nil
		}
		return fmt.Errorf("failed to insert worker: %w", err)
	}

	return nil
}

func (r *MongoRepository) Read(name string) (*dto.WorkerRegister, error) {
	var result *dto.WorkerRegister

	filter := bson.M{"_id": name}
	err := r.collection.FindOne(r.ctx, filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		} else {
			return nil, err
		}
	}

	return result, nil

}
