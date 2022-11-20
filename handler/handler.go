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

func GetGames(w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Query("SELECT * FROM games")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()

	gme := make([]*entities.Game, 0)
	for rows.Next() {
		ge := new(entities.Game)
		err := rows.Scan(&ge.ID, &ge.Title, &ge.Developer, &ge.Started, &ge.Finished)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

		gme = append(gme, ge)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, g := range gme {
		fmt.Fprintf(w, "%d %s %s %t %t\n", g.ID, g.Title, g.Developer, g.Started, g.Finished)
	}
}
func GetSortedGames(w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Query("SELECT * FROM games ORDER BY title ASC")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()

	gme := make([]*entities.Game, 0)
	for rows.Next() {
		ge := new(entities.Game)
		err := rows.Scan(&ge.ID, &ge.Title, &ge.Developer, &ge.Started, &ge.Finished)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

		gme = append(gme, ge)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, g := range gme {
		fmt.Fprintf(w, "%d %s %s %t %t\n", g.ID, g.Title, g.Developer, g.Started, g.Finished)
	}
}

func GetUnPlayedGames(w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Query("SELECT * FROM games WHERE started=false")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()

	gme := make([]*entities.Game, 0)
	for rows.Next() {
		ge := new(entities.Game)
		err := rows.Scan(&ge.ID, &ge.Title, &ge.Developer, &ge.Started, &ge.Finished)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

		gme = append(gme, ge)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, g := range gme {
		fmt.Fprintf(w, "%d %s %s %t %t\n", g.ID, g.Title, g.Developer, g.Started, g.Finished)
	}
}

func GetStartedUnfinishedGames(w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Query("SELECT * FROM games WHERE started=true AND finished=false")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()

	gme := make([]*entities.Game, 0)
	for rows.Next() {
		ge := new(entities.Game)
		err := rows.Scan(&ge.ID, &ge.Title, &ge.Developer, &ge.Started, &ge.Finished)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

		gme = append(gme, ge)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, g := range gme {
		fmt.Fprintf(w, "%d %s %s %t %t\n", g.ID, g.Title, g.Developer, g.Started, g.Finished)
	}
}

func GetFinishedGames(w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Query("SELECT * FROM games WHERE finished=true")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()

	gme := make([]*entities.Game, 0)
	for rows.Next() {
		ge := new(entities.Game)
		err := rows.Scan(&ge.ID, &ge.Title, &ge.Developer, &ge.Started, &ge.Finished)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

		gme = append(gme, ge)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, g := range gme {
		fmt.Fprintf(w, "%d %s %s %t %t\n", g.ID, g.Title, g.Developer, g.Started, g.Finished)
	}
}

func GetSpecificGame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]

	rows, err := DB.Query("SELECT * FROM games WHERE title=$1", title)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()
	var game *entities.Game
	for rows.Next() {
		game = new(entities.Game)
		err := rows.Scan(&game.ID, &game.Title, &game.Developer, &game.Started, &game.Finished)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
	}

	if game.Title == "" {
		http.Error(w, http.StatusText(404), 404)
		return
	}

	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	fmt.Fprintf(w, "%d %s %s %t %t\n", game.ID, game.Title, game.Developer, game.Started, game.Finished)
}

func PostGame(w http.ResponseWriter, r *http.Request) {
	var game entities.Game

	err := json.NewDecoder(r.Body).Decode(&game)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}
	result, err := DB.Exec("INSERT INTO games (title, developer, started, finished) VALUES ($1, $2, $3, $4)", game.Title, game.Developer, game.Started, game.Finished)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	if game.Title == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	fmt.Fprintf(w, "Game %s created successfully (%d row affected)\n", game.Title, rowsAffected)
}

func PutGame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]

	var game entities.Game

	err := json.NewDecoder(r.Body).Decode(&game)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}
	result, err := DB.Exec("UPDATE games SET title=$2, developer=$3, started=$4, finished=$5 WHERE title=$1", title, game.Title, game.Developer, game.Started, game.Finished)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	if game.Title == "" {
		http.Error(w, http.StatusText(404), 404)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	fmt.Fprintf(w, "Game %s was updated successfully (%d row affected)\n", title, rowsAffected)
}

func DeleteGame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]

	result, err := DB.Exec("DELETE FROM games WHERE title=$1", title)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	fmt.Fprintf(w, "Game %s deleted successfully (%d row affected)\n", title, rowsAffected)
}
