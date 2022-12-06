package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/atennyson/DBTest/handler"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func init() {
	var err error
	host := os.Getenv("HOST")
	po := os.Getenv("PORT")
	user := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")

	port, _ := strconv.Atoi(po)

	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

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

	fmt.Println("Listening on port 6089")
	http.ListenAndServe(":6089", r)
}
