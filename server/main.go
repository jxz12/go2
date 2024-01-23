package main

import (
    "net/http"
    "encoding/json"
)

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
    http.HandleFunc("/new", func(w http.ResponseWriter, r *http.Request) {
        // return index of player
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]int{"status": playerId})
        playerId += 1
    })
    http.HandleFunc("/board", func(w http.ResponseWriter, r *http.Request) {
        // should return state of board, and also if game is over
        // or maybe just always call play with nothing as heartbeat?
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(board)
    })
    http.HandleFunc("/play", func(w http.ResponseWriter, r *http.Request) {
        // take x,y location and return true or not
        // should return true or false and also state of board
        json.NewEncoder(w).Encode(r.URL.Query().Get("param1"))
    })
    http.HandleFunc("/finish", func(w http.ResponseWriter, r *http.Request) {
        // take x,y location and return true or not
    })
    http.ListenAndServe(":8080", nil)
}