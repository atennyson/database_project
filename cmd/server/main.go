package main

import (
	"database/sql"
	"fmt"
	"github.com/atennyson/DBTest/cmd/handler"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "sk8er4life"
	dbname   = "Video_Games"
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
	r.HandleFunc("/games/{title}", handler.GetSpecificGame).Methods("GET")
	r.HandleFunc("/games/newgame", handler.PostGame).Methods("POST")
	r.HandleFunc("/games/{title}", handler.PutGame).Methods("PUT")
	r.HandleFunc("/games/{title}", handler.DeleteGame).Methods("DELETE")
	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", r)
}
