package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"

	"github.com/victorgoecking/go-rest-api-basic/database"
)

type Game struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var games []Game

func GetGames(w http.ResponseWriter, r *http.Request) {
	conn := database.SetConnection()
	defer conn.Close()

	/*
		var datas []interface{}

		datas = append(datas,10)

		selDB, err := conn.Query("SELECT * FROM game where price=?", datas...)
	*/
	selDB, err := conn.Query("SELECT * FROM game")

	if err != nil {
		fmt.Println("Error to fetch", err)
	}

	for selDB.Next() {
		var game Game
		err = selDB.Scan(&game.ID, &game.Name, &game.Price)
		games = append(games, game)
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(games)
}

func GetGamesById(w http.ResponseWriter, r *http.Request) {
	conn := database.SetConnection()
	defer conn.Close()

	var game Game

	vars := mux.Vars(r)
	id := vars["gameID"]

	selDB := conn.QueryRow("SELECT * FROM game where id=" + id)

	selDB.Scan(&game.ID, &game.Name, &game.Price)

	encoder := json.NewEncoder(w)
	encoder.Encode(game)
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	conn := database.SetConnection()
	defer conn.Close()

	var register Game

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprint(w, "Bad Request")
	}

	json.Unmarshal(body, &register)

	action, err := conn.Prepare("INSERT INTO game (name, price) VALUES(?,?)")
	action.Exec(register.Name, register.Price)

	encoder := json.NewEncoder(w)
	encoder.Encode(games)
}

func LookForGame(w http.ResponseWriter, r *http.Request) {
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

func UpdateGame(w http.ResponseWriter, r *http.Request) {
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
