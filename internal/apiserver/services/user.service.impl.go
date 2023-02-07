package services

import (
	"context"
	"github.com/shakh9006/golang-gin-jwt-auth/internal/apiserver/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
)

type UserServiceImpl struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewUserService(collection *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceImpl{
		collection: collection,
		ctx:        ctx,
	}
}

func (uc *UserServiceImpl) FindUserById(id string) (*models.DBResponse, error) {
	oid, _ := primitive.ObjectIDFromHex(id)
	var user *models.DBResponse

	query := bson.M{"_id": oid}
	err := uc.collection.FindOne(uc.ctx, query).Decode(&user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *UserServiceImpl) FindUserByEmail(email string) (*models.DBResponse, error) {
	var user *models.DBResponse

	query := bson.M{"email": strings.ToLower(email)}
	err := uc.collection.FindOne(uc.ctx, query).Decode(&user)

	if err != nil {
		return nil, err
	}

	return user, nil
}
