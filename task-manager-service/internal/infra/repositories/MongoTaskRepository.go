package repositories

import (
	"context"
	"fmt"
	"task-manager-service/internal/domain"
	"task-manager-service/internal/infra/config"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type MongoTaskRepository struct {
	collection *mongo.Collection
	logger     *zap.Logger
	ctx        context.Context
}

func NewMongoTaskRepository(cfg *config.Config, ctx context.Context, logger *zap.Logger) (*MongoTaskRepository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		logger.Sugar().Fatalf("Connect failed: %v", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		logger.Sugar().Fatalf("Ping failed: %v", err)
	}

	logger.Info("Succesfully connected")

	db := client.Database(cfg.MongoDB)
	coll := db.Collection("tasks")

	return &MongoTaskRepository{
		collection: coll,
		logger:     logger,
		ctx:        ctx,
	}, nil
}

func (r *MongoTaskRepository) CreateTask(task domain.TaskInRepository) (uuid.UUID, error) {
	res, err := r.collection.InsertOne(r.ctx, task)

	if err != nil {
		panic(err)
	}

	return uuid.Parse(fmt.Sprint(res.InsertedID))
}
