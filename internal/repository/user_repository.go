package repository

import (
	"context"

	"ps-tag-onboarding-go/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository interface {
	Save(user domain.User) error
	Update(user domain.User) error
	FindByID(id string) (domain.User, bool)
	FindByName(firstName, lastName string) (domain.User, bool)
	FindAll() ([]domain.User, error)
	Delete(id string) error
}

type MongoUserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *MongoUserRepository {
	return &MongoUserRepository{
		collection: db.Collection("users"),
	}
}

func (r *MongoUserRepository) Save(user domain.User) error {
	ctx := context.TODO()

	if user.ID == "" {
		user.ID = primitive.NewObjectID().Hex()
		_, err := r.collection.InsertOne(ctx, user)
		return err
	}

	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": user}
	opts := options.Update().SetUpsert(true)
	_, err := r.collection.UpdateOne(ctx, filter, update, opts)
	return err
}

func (r *MongoUserRepository) Update(user domain.User) error {
	ctx := context.TODO()

	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": user}
	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *MongoUserRepository) FindByID(id string) (domain.User, bool) {
	ctx := context.TODO()
	var user domain.User

	filter := bson.M{"_id": id}
	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return user, false
	}
	return user, true
}

func (r *MongoUserRepository) FindByName(firstName, lastName string) (domain.User, bool) {
	ctx := context.TODO()

	var user domain.User
	filter := bson.M{"first_name": firstName, "last_name": lastName}
	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return user, false
	}
	return user, true
}

func (r *MongoUserRepository) FindAll() ([]domain.User, error) {
	ctx := context.TODO()
	var users []domain.User

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user domain.User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *MongoUserRepository) Delete(id string) error {
	ctx := context.TODO()

	filter := bson.M{"_id": id}
	_, err := r.collection.DeleteOne(ctx, filter)
	return err
}
