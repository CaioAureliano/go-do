package repository

import (
	"github.com/CaioAureliano/go-do/internal/todo/dto"
	"github.com/CaioAureliano/go-do/internal/todo/model"
	"github.com/CaioAureliano/go-do/pkg/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoRepository interface {
	GetById(id primitive.ObjectID) (*model.Todo, error)
	Create(todo *model.Todo) (*model.Todo, error)
	Find(filter primitive.M) (*dto.FindResponse, error)
	Update(todo *model.Todo) (*model.Todo, error)
	DeleteById(id primitive.ObjectID) error
}

type todoRepository struct {
	db *database.MongoConnection
}

const (
	COLLECTION_NAME = "todos"
)

func New() TodoRepository {
	return todoRepository{
		db: database.New(),
	}
}

func (r todoRepository) GetById(id primitive.ObjectID) (*model.Todo, error) {
	client := r.db.Connection(COLLECTION_NAME)
	defer r.db.Disconnect()

	var todo *model.Todo
	if err := client.FindOne(r.db.Ctx, bson.M{"_id": id}).Decode(&todo); err != nil {
		return nil, err
	}

	return todo, nil
}

func (r todoRepository) Create(todo *model.Todo) (*model.Todo, error) {
	client := r.db.Connection(COLLECTION_NAME)
	defer r.db.Disconnect()

	res, err := client.InsertOne(r.db.Ctx, todo)
	if err != nil {
		return nil, err
	}

	var created *model.Todo
	if err := client.FindOne(r.db.Ctx, bson.M{"_id": res.InsertedID}).Decode(&created); err != nil {
		return nil, err
	}

	return created, nil
}

func (r todoRepository) Find(filter primitive.M) (*dto.FindResponse, error) {
	client := r.db.Connection(COLLECTION_NAME)
	defer r.db.Disconnect()

	cursor, err := client.Find(r.db.Ctx, filter)
	if err != nil {
		return nil, err
	}

	var todos []*model.Todo
	if err := cursor.All(r.db.Ctx, &todos); err != nil {
		return nil, err
	}

	return &dto.FindResponse{
		Todos: todos,
		Count: len(todos),
	}, nil
}

func (r todoRepository) Update(todo *model.Todo) (*model.Todo, error) {
	client := r.db.Connection(COLLECTION_NAME)
	defer r.db.Disconnect()

	_, err := client.UpdateByID(r.db.Ctx, todo.ID, bson.M{"$set": todo})
	if err != nil {
		return nil, err
	}

	var updated *model.Todo
	if err := client.FindOne(r.db.Ctx, bson.M{"_id": todo.ID}).Decode(&updated); err != nil {
		return nil, err
	}

	return updated, nil
}

func (r todoRepository) DeleteById(id primitive.ObjectID) error {
	client := r.db.Connection(COLLECTION_NAME)
	defer r.db.Disconnect()

	if _, err := client.DeleteOne(r.db.Ctx, bson.M{"_id": id}); err != nil {
		return err
	}

	return nil
}
