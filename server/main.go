package main

import "net/http"
import "encoding/json"
import "fmt"

func main() {
	// TODO
	// open a goroutine for each user
	// need a mutex for the board?
	// matchmaking and new game making
	playerId := 0
	board := NewBoard(19)

	// websockets dont work because of this error
	// "websocket: RSV1 set, bad opcode 7, bad MASK"
	// so just use request-response instead
	http.HandleFunc("/player", func(w http.ResponseWriter, r *http.Request) {
		// return index of player
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]int{"status": playerId})
		playerId += 1
	})
	http.HandleFunc("/board", func(w http.ResponseWriter, r *http.Request) {
		// should return state of board, and also if game is over
		// or maybe just always call play with nothing as heartbeat?
		if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			// TODO include scores here as well
			json.NewEncoder(w).Encode(board)
		} else if r.Method == "PUT" {
			type Placement struct {
				Player int
				Row    int
				Col    int
			}
			var p Placement
			err := json.NewDecoder(r.Body).Decode(&p)
			if err != nil {
				fmt.Println(err)
				w.WriteHeader(400) // Bad Request
				return
			}
			if board.Play(p.Player, p.Row, p.Col) {
				w.WriteHeader(200) // Success
			} else {
				fmt.Println("Could not place stone at", p.Player, p.Row, p.Col)
				// TODO return reason why e.g. stone already there, Ko, self-capture
				w.WriteHeader(409) // Conflict
			}
			json.NewEncoder(w).Encode(board)
		}
	})
	http.HandleFunc("/finish", func(w http.ResponseWriter, r *http.Request) {
		// take x,y location and return true or not
	})
	http.ListenAndServe(":8080", nil)
}
