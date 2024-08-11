package repositories

import (
	"errors"

	domain "github.com/Johna210/A2SV-Backend-Track/Track7_clean_architecture/Domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

type taskRepository struct {
	database   mongo.Database
	collection string
}

func NewTaskRepository(db mongo.Database, collection string) domain.TaskRepository {
	return &taskRepository{
		database:   db,
		collection: collection,
	}
}

func (tr *taskRepository) CreateTask(c context.Context, task *domain.Task) error {
	collection := tr.database.Collection(tr.collection)

	_, err := collection.InsertOne(c, task)
	if err != nil {
		return err
	}

	return nil
}

func (tr *taskRepository) Fetch(c context.Context) ([]domain.Task, error) {
	collection := tr.database.Collection(tr.collection)

	cursor, err := collection.Find(c, nil)
	if err != nil {
		return nil, err
	}

	var tasks []domain.Task
	err = cursor.All(c, &tasks)
	if err != nil {
		return []domain.Task{}, err
	}

	return tasks, err
}

func (tr *taskRepository) GetByID(c context.Context, id string) (domain.Task, error) {
	collection := tr.database.Collection(tr.collection)

	var task domain.Task
	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return task, err
	}
	err = collection.FindOne(c, idHex).Decode(&task)
	return task, err
}

func (tr *taskRepository) UpdateTask(c context.Context, task *domain.TaskUpdate, id string) (domain.Task, error) {
	collection := tr.database.Collection(tr.collection)

	updateFields := make(bson.M)
	if task.Title != nil {
		updateFields["title"] = task.Title
	}

	if task.Status != nil {
		updateFields["status"] = task.Status
	}

	if task.Description != nil {
		updateFields["description"] = task.Description
	}

	var taskResponse domain.Task
	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return taskResponse, err
	}

	result, err := collection.UpdateOne(
		c,
		bson.M{"_id": idHex},
		bson.M{"$set": updateFields},
	)

	if err != nil {
		return taskResponse, err
	}

	if result.MatchedCount == 0 {
		return domain.Task{}, errors.New("no task found with the given ID")
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&taskResponse)
	if err != nil {
		return domain.Task{}, err
	}

	return taskResponse, nil
}

func (tr *taskRepository) DeleteTask(c context.Context, id string) error {
	collection := tr.database.Collection(tr.collection)

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	response, err := collection.DeleteOne(c, bson.M{"_id": idHex})
	if err != nil {
		return err
	}

	if response.DeletedCount == 0 {
		return errors.New("unable to delete Task")
	}

	return nil

}
