package mongoDb

import (
	"context"

	database "gin-mongo/configuration"
	models "gin-mongo/src/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var roleCollection *mongo.Collection = database.GetCollection(database.Client, "roles")

func GetRole(ctx context.Context, roleCode int) (models.Role, error) {
	filter := bson.M{
		"role_code": roleCode,
	}

	var res models.Role
	err := roleCollection.FindOne(ctx, filter).Decode(&res)

	if err != nil {
		return models.Role{}, err
	}

	return res, err
}
