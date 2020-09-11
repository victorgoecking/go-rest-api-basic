package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/victorgoecking/go-rest-api-basic/middlewares"
	"github.com/victorgoecking/go-rest-api-basic/routes"

	"github.com/gorilla/mux"
)

func majorRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
	fmt.Println("Hello World")
}

func setRoutes(router *mux.Router) {
	router.HandleFunc("/", majorRoute)
	router.HandleFunc("/games", routes.GetGames)
	router.HandleFunc("/games/{gameID}", routes.GetGamesById)
	router.HandleFunc("/newgame", routes.NewGame)
	router.HandleFunc("/lookforgame", routes.LookForGame)
	router.HandleFunc("/updategame", routes.UpdateGame)
}

func main() {
	var router *mux.Router

	log.Printf("Server is working on http://localhost:1602")

	router = mux.NewRouter()

	router.Use(middlewares.JsonMiddleware)

	setRoutes(router)

	err := http.ListenAndServe(":1602", router)
	if err != nil {
		fmt.Println("Erro", err)
	}
}
