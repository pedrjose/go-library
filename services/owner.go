package services

import (
	"context"
	"fmt"

	"github.com/pedrjose/go-library/db"
	"github.com/pedrjose/go-library/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetOwner(ctx context.Context, owner models.Owner) error {
	insertResult, err := db.OwnersCollection.InsertOne(ctx, owner)
	if err != nil {
		return err
	}
	fmt.Println(insertResult.InsertedID)
	return nil
}

func GetOwner(ctx context.Context, owner models.Owner) models.Owner {
	filter := bson.M{"email": owner.Email}
	var bookOwner models.Owner

	err := db.OwnersCollection.FindOne(context.TODO(), filter).Decode(&bookOwner)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("Query owner error:", err.Error())
		}
		panic(err)
	}

	return bookOwner
}
