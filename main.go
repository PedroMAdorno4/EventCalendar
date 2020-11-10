package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PedroMAdorno4/Desafio/pkg/auth"
	"github.com/PedroMAdorno4/Desafio/pkg/create"
	"github.com/PedroMAdorno4/Desafio/pkg/delete"
	"github.com/PedroMAdorno4/Desafio/pkg/http/rest"
	"github.com/PedroMAdorno4/Desafio/pkg/read"
	"github.com/PedroMAdorno4/Desafio/pkg/storage/mongodb"
	"github.com/PedroMAdorno4/Desafio/pkg/update"
)

func main() {
	st, _ := mongodb.NewStorage()

	auther := auth.NewService(st)
	// parser := parser.NewService(st)
	creater := create.NewService(st)
	reader := read.NewService(st)
	updater := update.NewService(st)
	deleter := delete.NewService(st)

	router := rest.Handler(auther, creater, reader, updater, deleter)

	fmt.Println("Listening on port 4444")
	log.Fatal(http.ListenAndServe(":4444", router))

}
