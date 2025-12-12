package repositories

import (
	"context"
	"fmt"
	"task-manager-service/internal/domain"
	"task-manager-service/internal/infra/config"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
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

func (r *MongoTaskRepository) Create(task *domain.TaskInRepository) (uuid.UUID, error) {

	res, err := r.collection.InsertOne(r.ctx, task)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to insert task: %w", err)
	}

	id, err := uuid.Parse(fmt.Sprint(res.InsertedID))
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid UUID: %w", err)
	}

	return id, nil
}

func (r *MongoTaskRepository) GetStatus(id string) (*string, error) {
	var result domain.TaskInRepository

	filter := bson.M{"_id": id}
	err := r.collection.FindOne(r.ctx, filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		} else {
			return nil, err
		}
	}

	return &result.Status, nil

}

func (r *MongoTaskRepository) Get(id string) (*domain.TaskInRepository, error) {
	var result domain.TaskInRepository

	filter := bson.M{"_id": id}
	err := r.collection.FindOne(r.ctx, filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		} else {
			return nil, err
		}
	}

	return &result, nil

}

func (r *MongoTaskRepository) Put(task domain.TaskInRepository) error {
	filter := bson.M{"_id": task.Id}
	update := bson.M{
		"$set": bson.M{
			"status":  task.Status,
			"type":    task.Type,
			"objects": task.Objects,
			"payload": task.Payload,
		},
	}

	res, err := r.collection.UpdateOne(r.ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update task: %w", err)
	}

	if res.MatchedCount == 0 {
		return fmt.Errorf("task with id %s not found", task.Id)
	}

	return nil
}
