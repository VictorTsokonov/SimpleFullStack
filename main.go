package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
)

// Game struct represents the structure of each game
type Game struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ReleaseDate int    `json:"release_date"`
}

// Database connection string (adjust as needed for your setup)
const (
	host     = "postgres"
	port     = 5432
	user     = "admin"
	password = "adminpassword"
	dbname   = "VOT"
)

var db *sql.DB

func main() {
	// Connect to PostgreSQL database
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	// Verify connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}
	fmt.Println("Connected to the PostgreSQL database")

	// Create the games table if it does not exist
	createTable := `CREATE TABLE IF NOT EXISTS games (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100),
		release_date INT
	);`

	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}
	fmt.Println("Ensured 'games' table exists")

	// Initialize the mux router
	router := mux.NewRouter()

	// Register routes
	router.HandleFunc("/games", getGames).Methods("GET")
	router.HandleFunc("/games", createGame).Methods("POST")
	router.HandleFunc("/games", deleteGame).Methods("DELETE")

	// Start the server
	fmt.Println("Server is listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Handler for GET /games - Returns a list of games from the database
func getGames(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name, release_date FROM games")
	if err != nil {
		http.Error(w, "Failed to retrieve games", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var games []Game
	for rows.Next() {
		var game Game
		if err := rows.Scan(&game.ID, &game.Name, &game.ReleaseDate); err != nil {
			http.Error(w, "Failed to scan game", http.StatusInternalServerError)
			return
		}
		games = append(games, game)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(games)
	fmt.Println("GET /games: Returning list of games")
}

// Handler for POST /games - Adds a new game to the database
func createGame(w http.ResponseWriter, r *http.Request) {
	var newGame Game

	// Decode the incoming JSON request
	err := json.NewDecoder(r.Body).Decode(&newGame)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Insert the new game into the database
	sqlStatement := `INSERT INTO games (name, release_date) VALUES ($1, $2) RETURNING id`
	err = db.QueryRow(sqlStatement, newGame.Name, newGame.ReleaseDate).Scan(&newGame.ID)
	if err != nil {
		http.Error(w, "Failed to create game", http.StatusInternalServerError)
		return
	}

	fmt.Printf("POST /games: Added new game: %+v\n", newGame)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newGame)
}

// Handler for DELETE /games?id=<id> - Deletes a game from the database by id
func deleteGame(w http.ResponseWriter, r *http.Request) {
	// Retrieve 'id' from query parameters
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	// Convert id to an integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	// Delete the game from the database
	sqlStatement := `DELETE FROM games WHERE id = $1`
	result, err := db.Exec(sqlStatement, id)
	if err != nil {
		http.Error(w, "Failed to delete game", http.StatusInternalServerError)
		return
	}

	// Check if a game was actually deleted
	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		http.Error(w, "Game not found", http.StatusNotFound)
		return
	}

	fmt.Printf("DELETE /games: Deleted game with id %d\n", id)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Game with id %d deleted", id)
}
