package db

import (
	"context"
	"time"

	"github.com/luisalfredosv/twiter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindProfile(ID string) (models.Usuario, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)

	defer cancel()

	db := MongoCN.Database("twiter")
	col := db.Collection("usuarios")

	var perfil models.Usuario

	objID, _ := primitive.ObjectIDFromHex(ID)

	with := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, with).Decode(&perfil)

	perfil.Password = ""

	if err != nil {
		return perfil, err
	}

	return perfil, nil

}
