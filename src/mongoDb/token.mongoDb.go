package mongoDb

import (
	"context"
	model "gin-mongo/src/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// var userCollection *mongo.Collection = database.GetCollection(database.Client, "tokens")

func CreateNewToken(ctx context.Context, objId primitive.ObjectID, token string) (int64, error) {
	filter := bson.M{
		"_id": objId,
	}

	update := bson.M{
		"$set": bson.M{
			"token": token,
		},
	}
	res, err := userCollection.UpdateOne(ctx, filter, update)

	if err != nil {
		return -1, err
	}

	return res.MatchedCount, nil
}

func GetToken(ctx context.Context, token string) error {
	filter := bson.M{
		"token": token,
	}

	var res model.User
	err := userCollection.FindOne(ctx, filter).Decode(&res)

	log.Println(err)

	if err != nil {
		return err
	}

	return nil

}

func DeleteToken(ctx context.Context, token string) (int64, error) {
	filter := bson.M{
		"token": token,
	}

	update := bson.M{
		"$set": bson.M{
			"token":     "",
			"is_logger": false,
		},
	}

	res, err := userCollection.UpdateOne(ctx, filter, update)

	if err != nil {
		return -1, err
	}

	return res.MatchedCount, nil
}
