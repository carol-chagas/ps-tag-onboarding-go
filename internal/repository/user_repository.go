package repository

import (
	"context"

	"ps-tag-onboarding-go/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Save(ctx context.Context, user domain.User) error
	Update(ctx context.Context, user domain.User) error
	FindByID(ctx context.Context, id string) (domain.User, bool)
	FindByEmail(ctx context.Context, email string) (domain.User, bool)
	FindByName(ctx context.Context, firstName, lastName string) (domain.User, bool)
	FindAll(ctx context.Context) ([]domain.User, error)
	Delete(ctx context.Context, id string) error
}

type MongoUserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *MongoUserRepository {
	return &MongoUserRepository{
		collection: db.Collection("users"),
	}
}

func (r *MongoUserRepository) Save(ctx context.Context, user domain.User) error {
	_, err := r.collection.InsertOne(ctx, user)
	return err
}

func (r *MongoUserRepository) Update(ctx context.Context, user domain.User) error {
	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": user}
	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *MongoUserRepository) FindByID(ctx context.Context, id string) (domain.User, bool) {
	var user domain.User

	filter := bson.M{"_id": id}
	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return user, false
	}
	return user, true
}

func (r *MongoUserRepository) FindByEmail(ctx context.Context, email string) (domain.User, bool) {
	var user domain.User
	filter := bson.M{"email": email}
	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return user, false
	}
	return user, true
}

func (r *MongoUserRepository) FindByName(ctx context.Context, firstName, lastName string) (domain.User, bool) {
	var user domain.User
	filter := bson.M{"first_name": firstName, "last_name": lastName}
	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return user, false
	}
	return user, true
}

func (r *MongoUserRepository) FindAll(ctx context.Context) ([]domain.User, error) {
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

func (r *MongoUserRepository) Delete(ctx context.Context, id string) error {
	filter := bson.M{"_id": id}
	_, err := r.collection.DeleteOne(ctx, filter)
	return err
}
