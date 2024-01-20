package main

import (
    "fmt"
    "net/http"
)
var upgrader = websocket.Upgrader{
    ReadBufferSize: 1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool { return true },
}

func main() {

    // TODO?
    // open a goroutine for each user
    // need a mutex for the board?

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "websockets.html")
    })
    http.HandleFunc("/new", func(w http.ResponseWriter, r *http.Request) {
        // return index of player
    })
    http.HandleFunc("/heartbeat", func(w http.ResponseWriter, r *http.Request) {
        // websockets dont work because of this error, so just use request-response instead
        // "websocket: RSV1 set, bad opcode 7, bad MASK"
        // should return if game is over
    })
    http.HandleFunc("/play", func(w http.ResponseWriter, r *http.Request) {
        // take x,y location and return true or not
    })
    http.HandleFunc("/finish", func(w http.ResponseWriter, r *http.Request) {
        // take x,y location and return true or not
    })
    http.ListenAndServe(":8080", nil)
}