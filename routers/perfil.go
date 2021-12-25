package routers

import (
	"encoding/json"
	"net/http"

	"github.com/luisalfredosv/twiter/db"
)

func Perfil(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	perfil, err := db.FindProfile(ID)

	if err != nil {
		http.Error(w, "No se encontro el usuario, "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(perfil)

}
