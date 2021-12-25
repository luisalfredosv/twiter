package jwt

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/luisalfredosv/twiter/db"
	"github.com/luisalfredosv/twiter/models"
)

var Correo string
var IDUser string

func Verify(accessToken string) (*models.Claim, bool, string, error) {
	key := []byte("TestKey")
	claims := &models.Claim{}

	splitToken := strings.Split(accessToken, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")

	}

	accessToken = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err == nil {
		_, find, _ := db.FindUser(claims.Correo)

		if find {
			Correo = claims.Correo
			IDUser = claims.ID.Hex()
		}
		return claims, find, IDUser, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}

	return claims, false, string(""), err
}
