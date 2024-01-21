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

    http.HandleFunc("/new", func(w http.ResponseWriter, r *http.Request) {
        // return index of player
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]int{"status": playerId})
        playerId += 1
    })
    http.HandleFunc("/heartbeat", func(w http.ResponseWriter, r *http.Request) {
        // websockets dont work because of this error, so just use request-response instead
        // "websocket: RSV1 set, bad opcode 7, bad MASK"
        // should return state of board, and also if game is over
        // or maybe just always call play with nothing as heartbeat?
    })
    http.HandleFunc("/play", func(w http.ResponseWriter, r *http.Request) {
        // take x,y location and return true or not
        // should return state of board
    })
    http.HandleFunc("/finish", func(w http.ResponseWriter, r *http.Request) {
        // take x,y location and return true or not
    })
    http.ListenAndServe(":8080", nil)
}