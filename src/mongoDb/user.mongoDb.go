package mongoDb

import (
	"context"
	"log"
	"time"

	database "gin-mongo/configuration"
	models "gin-mongo/src/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection = database.GetCollection(database.Client, "users")

func CreateUser(user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, bson.M{"name": user.Name, "age": user.Age, "status": "active"})

	if err != nil {
		return err
	}

	return nil
}

func GetAllUsers() ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := collection.Find(ctx, bson.M{"status": "active"})

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer res.Close(ctx)

	users := []models.User{}
	for res.Next(ctx) {
		var user models.User
		if err = res.Decode(&user); err != nil {
			log.Fatal(err)
		}

		users = append(users, user)
	}

	return users, err
}

func GetUserById(objId primitive.ObjectID) (bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var res bson.M
	err := collection.FindOne(ctx, bson.M{"_id": objId, "status": "active"}).Decode(&res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func UpdateUser(objId primitive.ObjectID, user models.User) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := collection.UpdateOne(ctx, bson.M{"_id": objId, "status": "active"}, bson.M{"$set": user})

	if err != nil {
		return -1, err
	}

	return res.MatchedCount, nil
}

func DeleteUserById(objId primitive.ObjectID) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := collection.UpdateOne(ctx, bson.M{"_id": objId, "status": "active"}, bson.M{"$set": bson.M{"status": "deleted"}})

	if err != nil {
		return -1, err
	}

	return res.MatchedCount, nil
}
