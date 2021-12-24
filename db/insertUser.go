package db

import (
	"context"
	"time"

	"github.com/luisalfredosv/twiter/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertUser(u models.Usuario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twiter")
	col := db.Collection("usuarios")

	u.Password, _ = Hash(u.Password)
	result, err := col.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
