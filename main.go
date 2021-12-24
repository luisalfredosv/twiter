package main

import (
	"log"

	"github.com/luisalfredosv/twiter/db"
	"github.com/luisalfredosv/twiter/handlers"
)

func main() {
	if !db.CheckConnection() {
		log.Fatal("No hay conexión con la base de datos")
		return
	}

	handlers.Handlers()
}
