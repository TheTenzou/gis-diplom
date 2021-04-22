package repository

import (
	"context"
	"log"

	"github.com/TheTenzou/diplom2.0/user-service/apperrors"
	"github.com/TheTenzou/diplom2.0/user-service/interfaces"
	"github.com/TheTenzou/diplom2.0/user-service/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoUserRepository struct {
	Users *mongo.Collection
}

func NewMongoUserRepository(usersCollection *mongo.Collection) interfaces.UserRepository {
	return &mongoUserRepository{
		Users: usersCollection,
	}
}

func (r *mongoUserRepository) FindByID(
	ctx context.Context,
	userID primitive.ObjectID,
) (model.User, error) {

	var user model.User

	err := r.Users.FindOne(ctx, model.User{ID: userID}).Decode(&user)
	if err != nil {
		return user, apperrors.NewNotFound("id", userID.Hex())
	}

	return user, nil
}

func (r *mongoUserRepository) FindByLogin(
	ctx context.Context,
	userLogin string,
) (model.User, error) {

	var user model.User

	err := r.Users.FindOne(ctx, model.User{Login: userLogin}).Decode(&user)
	if err != nil {
		return user, apperrors.NewNotFound("login", userLogin)
	}

	return user, nil
}

func (r *mongoUserRepository) FindAll(
	ctx context.Context,
	page int,
) ([]model.User, error) {
	panic("not implemented")
}

func (r *mongoUserRepository) Create(
	ctx context.Context,
	user model.User,
) (model.User, error) {
	log.Println("1")

	userID, err := r.Users.InsertOne(ctx, user)
	if err != nil {
		if err, ok := err.(mongo.WriteException); ok && err.HasErrorMessage("duplicate key error collection") {
			log.Printf(
				"Could not create a user with login: %v. Reason %v\n",
				user.Login,
				"duplicate key error collection",
			)
		}
		return model.User{}, apperrors.NewConflict("login", user.Login)
	}

	user.ID = userID.InsertedID.(primitive.ObjectID)

	return user, nil
}

func (r *mongoUserRepository) Update(
	ctx context.Context,
	user model.User,
) error {

	_, err := r.Users.UpdateByID(ctx, user.ID, bson.M{"$set": user})

	if err != nil {
		if err, ok := err.(mongo.WriteException); ok && err.HasErrorMessage("duplicate key error collection") {
			log.Printf(
				"Could not update a user with new login: %v. Reason %v\n",
				user.Login,
				"duplicate key error collection",
			)
		}
		return apperrors.NewConflict("login", user.Login)
	}

	return nil
}

func (r *mongoUserRepository) Delete(
	ctx context.Context,
	userID primitive.ObjectID,
) (model.User, error) {

	result := r.Users.FindOneAndDelete(ctx, bson.M{"_id": userID})

	var deletedUser model.User

	err := result.Decode(&deletedUser)
	if err != nil {
		log.Printf("Ubancle to delete user: %v\n", err)
		return model.User{}, apperrors.NewInternal()
	}

	return deletedUser, nil
}
