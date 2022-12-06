package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/atennyson/DBTest/entities"
	"github.com/gorilla/mux"
)

var DB *sql.DB

func GetGamesHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Query("SELECT * FROM games")
	if err != nil {
		w.WriteHeader(400)
		return
	}
	defer rows.Close()

	games := []entities.Game{}]
	for rows.Next() {
		game := entities.Game{}
		err := rows.Scan(&game.ID, &game.Title, &game.Developer, &game.Started, &game.Finished)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		games = append(games, game)
	}

	if err = rows.Err(); err != nil {
		w.WriteHeader(500)
		return
	}

	for _, game := range games {
		fmt.Fprintf(w, "ID: %d Title: %s Developer: %s Started?: %t Finished?: %t\n", game.ID, game.Title, game.Developer, game.Started, game.Finished)
	}
}
func GetSortedGamesHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Query("SELECT * FROM games ORDER BY title ASC")
	if err != nil {
		w.WriteHeader(400)
		return
	}
	defer rows.Close()

	games := []entities.Game{}
	for rows.Next() {
		game := entities.Game{}
		err := rows.Scan(&game.ID, &game.Title, &game.Developer, &game.Started, &game.Finished)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		games = append(games, game)
	}

	if err = rows.Err(); err != nil {
		w.WriteHeader(500)
		return
	}

	for _, game := range games {
		fmt.Fprintf(w, "ID: %d Title: %s Developer: %s Started?: %t Finished?: %t\n", game.ID, game.Title, game.Developer, game.Started, game.Finished)
	}
}

func GetUnPlayedGamesHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Query("SELECT * FROM games WHERE started=false")
	if err != nil {
		w.WriteHeader(400)
		return
	}
	defer rows.Close()

	games := []entities.Game{}
	for rows.Next() {
		game := entities.Game{}
		err := rows.Scan(&game.ID, &game.Title, &game.Developer, &game.Started, &game.Finished)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		games = append(games, game)
	}

	if err = rows.Err(); err != nil {
		w.WriteHeader(500)
		return
	}

	for _, game := range games {
		fmt.Fprintf(w, "ID: %d Title: %s Developer: %s Started?: %t Finished?: %t\n", game.ID, game.Title, game.Developer, game.Started, game.Finished)
	}
}

func GetStartedUnfinishedGamesHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Query("SELECT * FROM games WHERE started=true AND finished=false")
	if err != nil {
		w.WriteHeader(400)
		return
	}
	defer rows.Close()

	games := []entities.Game{}
	for rows.Next() {
		game := entities.Game{}
		err := rows.Scan(&game.ID, &game.Title, &game.Developer, &game.Started, &game.Finished)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		games = append(games, game)
	}

	if err = rows.Err(); err != nil {
		w.WriteHeader(500)
		return
	}

	for _, game := range games {
		fmt.Fprintf(w, "ID: %d Title: %s Developer: %s Started?: %t Finished?: %t\n", game.ID, game.Title, game.Developer, game.Started, game.Finished)
	}
}

func GetFinishedGamesHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Query("SELECT * FROM games WHERE finished=true")
	if err != nil {
		w.WriteHeader(400)
		return
	}
	defer rows.Close()

	games := []entities.Game{}
	for rows.Next() {
		game := entities.Game{}
		err := rows.Scan(&game.ID, &game.Title, &game.Developer, &game.Started, &game.Finished)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		games = append(games, game)
	}

	if err = rows.Err(); err != nil {
		w.WriteHeader(500)
		return
	}

	for _, game := range games {
		fmt.Fprintf(w, "ID: %d Title: %s Developer: %s Started?: %t Finished?: %t\n", game.ID, game.Title, game.Developer, game.Started, game.Finished)
	}
}

func GetSpecificGameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]

	if !IterateData(title) {
		w.WriteHeader(404)
		return
	}

	rows, err := DB.Query("SELECT * FROM games WHERE title=$1", title)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	defer rows.Close()
	game := entities.Game{}
	for rows.Next() {
		err := rows.Scan(&game.ID, &game.Title, &game.Developer, &game.Started, &game.Finished)
		if err != nil {
			w.WriteHeader(500)
			return
		}
	}

	if err = rows.Err(); err != nil {
		w.WriteHeader(500)
		return
	}

	fmt.Fprintf(w, "ID: %d Title: %s Developer: %s Started?: %t Finished?: %t\n", game.ID, game.Title, game.Developer, game.Started, game.Finished)
}

func AddGameHandler(w http.ResponseWriter, r *http.Request) {
	game := entities.Game{}

	err := json.NewDecoder(r.Body).Decode(&game)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	if IterateData(game.Title) {
		w.WriteHeader(400)
		fmt.Fprint(w, "Game already exists.")
		return
	}
	result, err := DB.Exec("INSERT INTO games (title, developer, started, finished) VALUES ($1, $2, $3, $4)", game.Title, game.Developer, game.Started, game.Finished)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		w.WriteHeader(500)
		return
	}

	fmt.Fprintf(w, "Game %s created successfully (%d row affected)\n", game.Title, rowsAffected)
}

func UpdateGameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]

	game := entities.Game{}

	err := json.NewDecoder(r.Body).Decode(&game)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	if !IterateData(title) {
		w.WriteHeader(404)
		return
	}
	result, err := DB.Exec("UPDATE games SET title=$2, developer=$3, started=$4, finished=$5 WHERE title=$1", title, game.Title, game.Developer, game.Started, game.Finished)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		w.WriteHeader(500)
		return
	}

	fmt.Fprintf(w, "Game %s was updated successfully (%d row affected)\n", title, rowsAffected)
}

func DeleteGameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]

	if !IterateData(title) {
		w.WriteHeader(404)
		return
	}

	result, err := DB.Exec("DELETE FROM games WHERE title=$1", title)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		w.WriteHeader(500)
		return
	}

	fmt.Fprintf(w, "Game %s deleted successfully (%d row affected)\n", title, rowsAffected)
}

func IterateData(title string) bool {
	results, err := DB.Query("SELECT * FROM games")
	if err != nil {
		return false
	}
	defer results.Close()

	games := []entities.Game{}
	for results.Next() {
		game := entities.Game{}
		err := results.Scan(&game.ID, &game.Title, &game.Developer, &game.Started, &game.Finished)
		if err != nil {
			return false
		}

		games = append(games, game)
	}

	for _, game := range games {
		if game.Title == title {
			return true
		}
	}

	return false
}
