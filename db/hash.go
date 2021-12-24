package db

import "golang.org/x/crypto/bcrypt"

func Hash(password string) (string, error) {
	salt := 8
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), salt)
	return string(hashed), err
}
