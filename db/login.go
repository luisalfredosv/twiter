package db

import (
	"github.com/luisalfredosv/twiter/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(correo string, password string) (models.Usuario, bool) {
	usr, find, _ := FindUser(correo)

	if !find {
		return usr, false
	}

	passwordBytes := []byte(password)
	hashed := []byte(usr.Password)

	err := bcrypt.CompareHashAndPassword(hashed, passwordBytes)

	if err != nil {
		return usr, false
	}

	return usr, true

}
