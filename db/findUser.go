package db

import (
	"context"
	"time"

	"github.com/luisalfredosv/twiter/models"
	"go.mongodb.org/mongo-driver/bson"
)

func FindUser(correo string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twiter")
	col := db.Collection("usuarios")
	with := bson.M{"correo": correo}

	var result models.Usuario
	err := col.FindOne(ctx, with).Decode(&result)
	ID := result.ID.Hex()

	if err != nil {
		return result, false, ID
	}

	return result, true, ID
}
