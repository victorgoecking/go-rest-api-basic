package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Game struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var games []Game = []Game{
	{
		ID:    1,
		Name:  "Nioh",
		Price: 12.50,
	},
	{
		ID:    2,
		Name:  "Red Dead Redemption",
		Price: 240,
	},
	{
		ID:    3,
		Name:  "FIFA 21",
		Price: 180,
	},
	{
		ID:    4,
		Name:  "GTA V",
		Price: 150.50,
	},
	{
		ID:    5,
		Name:  "Cyberpunk",
		Price: 400.50,
	},
}

func majorRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func getGames(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.Encode(games)
}

func newgame(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)

	var register Game

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprint(w, "Bad Request")
	}

	json.Unmarshal(body, &register)
	games = append(games, register)

	encoder := json.NewEncoder(w)
	encoder.Encode(games)
}

func lookForGame(w http.ResponseWriter, r *http.Request) {
	var fragments = strings.Split(r.URL.Path, "/")

	//fmt.Println(r.Method)
	id, err := strconv.Atoi(fragments[2])
	if err != nil {
		fmt.Println("Erro to Convert")
	}

	for _, valor := range games {
		if valor.ID == id {
			json.NewEncoder(w).Encode(valor)
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func updateGame(w http.ResponseWriter, r *http.Request) {
	var register Game

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, "Bad Request")
	}
	json.Unmarshal(body, &register)

	var findIndex = -1

	for chave, valor := range games {
		if valor.ID == register.ID {
			findIndex = chave
		}
	}

	if findIndex > 0 {
		games[findIndex] = register
		encoder := json.NewEncoder(w)
		encoder.Encode(register)
	}
}

func main() {
	log.Printf("Server is working on http://localhost:1602")

	http.HandleFunc("/", majorRoute)
	http.HandleFunc("/games", getGames)
	http.HandleFunc("/newgame", newgame)
	http.HandleFunc("/lookforgame/", lookForGame)
	http.HandleFunc("/updategame", updateGame)

	http.ListenAndServe(":1602", nil)
}