package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/atennyson/DBTest/cmd/entities"
	"github.com/gorilla/mux"
	"net/http"
)

var DB *sql.DB

func GetGames(w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Query("SELECT * FROM video_games")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()

	gme := make([]*entities.Game, 0)
	for rows.Next() {
		ge := new(entities.Game)
		err := rows.Scan(&ge.ID, &ge.Title, &ge.Genre)
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
		fmt.Fprintf(w, "%d %s %s\n", g.ID, g.Title, g.Genre)
	}
}

func GetSpecificGame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]

	rows, err := DB.Query("SELECT * FROM video_games WHERE title = $1", title)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()
	var game *entities.Game
	for rows.Next() {
		game = new(entities.Game)
		err := rows.Scan(&game.ID, &game.Title, &game.Genre)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
	}

	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	fmt.Fprintf(w, "%d %s %s\n", game.ID, game.Title, game.Genre)
}

func PostGame(w http.ResponseWriter, r *http.Request) {
	var game entities.Game

	err := json.NewDecoder(r.Body).Decode(&game)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}
	result, err := DB.Exec("INSERT INTO video_games (title, genre) VALUES ($1, $2)", game.Title, game.Genre)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
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
	result, err := DB.Exec("UPDATE video_games SET title = $2, genre = $3 WHERE title = $1", title, game.Title, game.Genre)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
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

	result, err := DB.Exec("DELETE FROM video_games WHERE title = $1", title)
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
