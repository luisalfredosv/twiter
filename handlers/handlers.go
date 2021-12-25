package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/luisalfredosv/twiter/middlewares"
	"github.com/luisalfredosv/twiter/routers"
	"github.com/rs/cors"
)

func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlewares.CheckDB(routers.Registro)).Methods("POST")

	router.HandleFunc("/login", middlewares.CheckDB(routers.Login)).Methods("POST")

	router.HandleFunc("/perfil", middlewares.CheckDB(middlewares.VerifyJWT(routers.Perfil))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
