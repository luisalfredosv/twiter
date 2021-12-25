package middlewares

import (
	"net/http"

	"github.com/luisalfredosv/twiter/jwt"
)

func VerifyJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := jwt.Verify(r.Header.Get("Authorization"))

		if err != nil {
			http.Error(w, "Error en el token de autorizacion, "+err.Error(), http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	}
}
