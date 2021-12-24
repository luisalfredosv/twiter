package middlewares

import (
	"net/http"

	"github.com/luisalfredosv/twiter/db"
)

func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !db.CheckConnection() {
			http.Error(w, "Conexión perdida con la base de datos", 500)
		}

		next.ServeHTTP(w, r)
	}
}
