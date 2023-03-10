package mongoDb

import (
	"context"

	database "gin-mongo/configuration"
	models "gin-mongo/src/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var roleCollection *mongo.Collection = database.GetCollection(database.Client, "roles")

// func GetRole(ctx context.Context, roleCode int) (models.Role, error) {
// 	filter := bson.M{
// 		"role_code": roleCode,
// 	}

// 	var res models.Role
// 	err := roleCollection.FindOne(ctx, filter).Decode(&res)

// 	if err != nil {
// 		return models.Role{}, err
// 	}

// 	return res, err
// }

func GetRolesList(ctx context.Context) ([]models.Role, error) {
	res, err := roleCollection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	defer res.Close(ctx)

	roles := []models.Role{}
	for res.Next(ctx) {
		var role models.Role
		if err = res.Decode(&role); err != nil {
			return nil, err
		}

		roles = append(roles, role)
	}

	return roles, err
}
