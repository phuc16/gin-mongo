package mongoDb

import (
	"context"

	database "gin-mongo/configuration"
	models "gin-mongo/src/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection = database.GetCollection(database.Client, "users")

func GetAllUsers(ctx context.Context, fromDate string, toDate string) ([]models.User, error) {
	res, err := collection.Find(ctx, bson.M{"status": "active", "created_at": bson.M{"$gte": fromDate, "$lt": toDate}})

	if err != nil {
		return nil, err
	}

	defer res.Close(ctx)

	users := []models.User{}
	for res.Next(ctx) {
		var user models.User
		if err = res.Decode(&user); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, err
}

func GetUserById(ctx context.Context, filter bson.M) (models.User, error) {
	var res models.User
	err := collection.FindOne(ctx, filter).Decode(&res)

	if err != nil {
		return models.User{}, err
	}

	return res, nil
}

func GetUserByName(ctx context.Context, filter bson.M) (models.User, error) {
	var res models.User
	err := collection.FindOne(ctx, filter).Decode(&res)

	if err != nil {
		return models.User{}, err
	}

	return res, nil
}

func UpdateUserById(ctx context.Context, filter bson.M, user models.User) (int64, error) {
	res, err := collection.UpdateOne(ctx, filter, bson.M{"$set": user})

	if err != nil {
		return -1, err
	}

	return res.MatchedCount, nil
}

func DeleteUserById(ctx context.Context, filter bson.M, set bson.M) (int64, error) {

	res, err := collection.UpdateOne(ctx, filter, bson.M{"$set": set})

	if err != nil {
		return -1, err
	}

	return res.MatchedCount, nil
}

func CreateUserNew(ctx context.Context, user models.User) error {

	_, err := collection.InsertOne(ctx, user)

	if err != nil {
		return err
	}

	return nil
}

func UserLogin(ctx context.Context, filter bson.M) (int64, error) {
	res, err := collection.UpdateOne(ctx, filter, bson.M{"$set": bson.M{"is_logged": true}})

	if err != nil {
		return -1, err
	}

	return res.MatchedCount, nil
}

func UserLogout(ctx context.Context, filter bson.M) (int64, error) {
	res, err := collection.UpdateOne(ctx, filter, bson.M{"$set": bson.M{"is_logged": false}})

	if err != nil {
		return -1, err
	}

	return res.MatchedCount, nil
}
