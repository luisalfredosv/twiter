package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/luisalfredosv/twiter/db"
	"github.com/luisalfredosv/twiter/jwt"
	"github.com/luisalfredosv/twiter/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Usuario y/o contraseña incorrectos"+err.Error(), 400)
	}

	if len(t.Correo) == 0 {
		http.Error(w, "El correo es requerido", 400)
		return
	}

	usr, find := db.Login(t.Correo, t.Password)

	if !find {
		http.Error(w, "Usuario y/o contraseña incorrectos", 400)
		return
	}

	AccessToken, err := jwt.Sign(usr)
	if err != nil {
		http.Error(w, "Ocurrio un error al interenar generar el Token"+err.Error(), 400)
	}

	resp := models.ResponseLogin{
		AccessToken: AccessToken,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "AccessToken",
		Value:   AccessToken,
		Expires: expirationTime,
	})
}
