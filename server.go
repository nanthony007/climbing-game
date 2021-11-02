package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	core "github.com/nanthony007/climbing-game/pkg"
)

func main() {
	http.HandleFunc("/game", gameHandler)
	log.Panic(http.ListenAndServe(":8080", nil))
}

func gameHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		game, err := core.NewGame()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		decodeErr := json.NewDecoder(r.Body).Decode(&game)
		if decodeErr != nil {
			http.Error(w, decodeErr.Error(), http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(game)
		pts, _ := game.TotalPoints()
		fmt.Fprint(w, "Game was worth ", pts, " points.")
	default:
		fmt.Fprintf(w, "Sorry, only POST methods are supported.")
	}
}
