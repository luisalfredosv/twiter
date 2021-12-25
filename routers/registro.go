package routers

import (
	"encoding/json"
	"net/http"

	"github.com/luisalfredosv/twiter/db"
	"github.com/luisalfredosv/twiter/models"
)

func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}

	if len(t.Correo) == 0 {
		http.Error(w, "El correo es requerido", 400)
	}

	if len(t.Password) < 6 {
		http.Error(w, "La contraseña debe ser mayor a 6 caracteres", 400)
	}

	_, find, _ := db.FindUser(t.Correo)

	if find == true {
		http.Error(w, "Ya existe un usuario con ese correo", 400)
		return
	}

	_, status, err := db.InsertUser(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar realizar el error"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se logró guardar el usuario"+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
