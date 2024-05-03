package repositories

import (
	"context"

	"github.com/kylerequez/go-crud-api/src/common"
	"github.com/kylerequez/go-crud-api/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	databaseName   string
	collectionName string
}

func NewUserRepository(databaseName string, collectionName string) *UserRepository {
	return &UserRepository{
		databaseName:   databaseName,
		collectionName: collectionName,
	}
}

func (ur *UserRepository) GetAllUsers() ([]primitive.M, error) {
	if err := common.ConnectDB(); err != nil {
		return nil, err
	}
	defer common.CloseDB()

	DB := common.GetDB(ur.databaseName)
	ctx := context.TODO()

	cursor, err := DB.Collection(ur.collectionName).Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []bson.M
	if err = cursor.All(context.TODO(), &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (ur *UserRepository) GetUserById(oid primitive.ObjectID) (*models.User, error) {
	if err := common.ConnectDB(); err != nil {
		return nil, err
	}
	defer common.CloseDB()

	DB := common.GetDB(ur.databaseName)
	ctx := context.TODO()

	filter := bson.D{{"_id", oid}}
	var user models.User
	err := DB.Collection(ur.collectionName).FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) InsertUser(newUser models.User) (*mongo.InsertOneResult, error) {
	if err := common.ConnectDB(); err != nil {
		return nil, err
	}
	defer common.CloseDB()

	DB := common.GetDB(ur.databaseName)
	ctx := context.TODO()

	result, err := DB.Collection(ur.collectionName).InsertOne(ctx, newUser)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (ur *UserRepository) UpdateUser(oid primitive.ObjectID, user models.User) (*models.User, error) {
	if err := common.ConnectDB(); err != nil {
		return nil, err
	}
	defer common.CloseDB()

	DB := common.GetDB(ur.databaseName)
	ctx := context.TODO()

	filter := bson.D{{"_id", oid}}
	update := bson.D{{"$set", bson.D{{"username", user.Username}, {"password", user.Password}}}}
	var updatedUser models.User
	err := DB.Collection(ur.collectionName).FindOneAndUpdate(ctx, filter, update).Decode(&updatedUser)
	if err != nil {
		return nil, err
	}

	return &updatedUser, nil
}

func (ur *UserRepository) DeleteUserById(oid primitive.ObjectID) (*mongo.DeleteResult, error) {
	if err := common.ConnectDB(); err != nil {
		return nil, err
	}
	defer common.CloseDB()

	DB := common.GetDB(ur.databaseName)
	ctx := context.TODO()

	filter := bson.D{{"_id", oid}}
	result, err := DB.Collection(ur.collectionName).DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}
