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
	password = "sk8er4life"
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
	r.HandleFunc("/games", handler.GetGames).Methods("GET")
	r.HandleFunc("/games/sorted", handler.GetSortedGames).Methods("GET")
	r.HandleFunc("/games/unplayed", handler.GetUnPlayedGames).Methods("GET")
	r.HandleFunc("/games/started/unfinished", handler.GetStartedUnfinishedGames).Methods("GET")
	r.HandleFunc("/games/finished", handler.GetFinishedGames).Methods("GET")
	r.HandleFunc("/games/{title}", handler.GetSpecificGame).Methods("GET")
	r.HandleFunc("/games/newgame", handler.PostGame).Methods("POST")
	r.HandleFunc("/games/{title}", handler.PutGame).Methods("PUT")
	r.HandleFunc("/games/{title}", handler.DeleteGame).Methods("DELETE")

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", r)
}
