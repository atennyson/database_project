package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/atennyson/DBTest/handler"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "W2xc7ig5GH!$32"
	dbname   = "owned_switch_games"
)

func init() {

	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	handler.DB, err = sql.Open("postgres", sqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	if err = handler.DB.Ping(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/games", handler.GetGamesHandler).Methods("GET")
	r.HandleFunc("/games/sorted", handler.GetSortedGamesHandler).Methods("GET")
	r.HandleFunc("/games/unplayed", handler.GetUnPlayedGamesHandler).Methods("GET")
	r.HandleFunc("/games/started/unfinished", handler.GetStartedUnfinishedGamesHandler).Methods("GET")
	r.HandleFunc("/games/finished", handler.GetFinishedGamesHandler).Methods("GET")
	r.HandleFunc("/games/{title}", handler.GetSpecificGameHandler).Methods("GET")
	r.HandleFunc("/games/newgame", handler.AddGameHandler).Methods("POST")
	r.HandleFunc("/games/{title}", handler.UpdateGameHandler).Methods("PUT")
	r.HandleFunc("/games/{title}", handler.DeleteGameHandler).Methods("DELETE")

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", r)
}
