package repo

import (
	"context"
	"fmt"
	"service_2/internal/models"
	"service_2/internal/usecase"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repo struct {
	*mongo.Client
}

func NewRepo(db *mongo.Client) usecase.CreateAndRead {
	return &Repo{db}
}

func (r *Repo) CreateUser(u models.User) error {
	collection := r.Client.Database("test").Collection("users")

	InsertRes, err := collection.InsertOne(context.TODO(), u)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return err
	}
	fmt.Println("Inserted a single document: ", InsertRes.InsertedID)

	return nil
}

func (r *Repo) ReadUser(u models.User) (models.User, error) {
	var result models.User
	collection := r.Client.Database("test").Collection("users")

	filter := bson.M{"email": u.Email}

	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return models.User{}, err
	}
	fmt.Printf("Found a single document: %+v\n", result)

	return result, nil
}
