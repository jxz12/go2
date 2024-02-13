package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func main() {
	hub := NewHub(19)
	go hub.Fanout()

	http.HandleFunc("/play", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		// TODO: communicate to decide playerId, colour etc.

		player := hub.NewPlayer(conn)
		go player.Sub()
		go player.Pub()
	})

	// port 8080 does not work on Mac, possibly due to parental controls
	// https://discussions.apple.com/thread/253069437
	http.ListenAndServe(":8081", nil)
}
