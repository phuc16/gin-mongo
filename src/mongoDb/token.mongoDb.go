package mongoDb

import (
	"context"
	"log"
	"time"

	database "gin-mongo/configuration"
	models "gin-mongo/src/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var tokenCollection *mongo.Collection = database.GetCollection(database.Client, "tokens")

func CreateNewToken(ctx context.Context, token models.Token) error {
	_, err := tokenCollection.InsertOne(ctx, token)

	if err != nil {
		return err
	}

	return nil
}

func GetToken(ctx context.Context, token string) error {
	filter := bson.M{
		"access_token": token,
		"disabled":     false,
		"expired_at": bson.M{
			"$gt": time.Now().Format(timeFormat),
		},
	}

	var res models.Token
	err := tokenCollection.FindOne(ctx, filter).Decode(&res)

	log.Println(err)

	if err != nil {
		return err
	}

	return nil

}

func DeleteToken(ctx context.Context, token string) (int64, error) {
	filter := bson.M{
		"access_token": token,
	}

	update := bson.M{
		"$set": bson.M{
			"disabled": true,
		},
	}

	res, err := tokenCollection.UpdateOne(ctx, filter, update)

	if err != nil {
		return -1, err
	}

	return res.MatchedCount, nil
}
