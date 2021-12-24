package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Usuario struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre     string             `bson:"nombre,omitempty" json:"nombre"`
	Apellidos  string             `bson:"apellido,omitempty" json:"apellido"`
	Cumpleanos time.Time          `bson:"cumpleano,omitempty" json:"cumpleano"`
	Correo     string             `bson:"correo,omitempty" json:"correo"`
	Password   string             `bson:"password,omitempty" json:"password"`
	Avatar     string             `bson:"avatar,omitempty" json:"avatar"`
	Banner     string             `bson:"banner,omitempty" json:"banner"`
	Biografia  string             `bson:"biografia,omitempty" json:"biografia"`
	SitioWeb   string             `bson:"sitioWeb,omitempty" json:"sitioWeb"`
}
