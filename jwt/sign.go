package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/luisalfredosv/twiter/models"
)

func Sign(t models.Usuario) (string, error) {

	key := []byte("TestKey")

	payload := jwt.MapClaims{
		"correo":    t.Correo,
		"nombre":    t.Nombre,
		"apellidos": t.Apellidos,
		"cumpleano": t.Cumpleanos,
		"biografia": t.Biografia,
		// "ubicacion": t.,
		"sitioWeb": t.SitioWeb,
		"_id":      t.ID.Hex(),
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenStr, err := token.SignedString(key)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
