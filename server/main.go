package main

import (
	"encoding/json"
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
	// matchmaking and new game making
	playerId := 0
	board := NewBoard(19)
	http.HandleFunc("/play", func(w http.ResponseWriter, r *http.Request) {

		// this is already a goroutine that has been spawned per player

		// Upgrade the HTTP connection to a WebSocket connection.
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		playerId += 1
		str, _ := json.Marshal(board)
		conn.WriteMessage(websocket.TextMessage, str)

		for {
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				fmt.Println(messageType, message, err)
				return
			}
			type Placement struct {
				Player int
				Row    int
				Col    int
			}
			var p Placement
			err = json.Unmarshal(message, &p)
			if err != nil {
				fmt.Println(err)
			}
			// TODO maybe do this?
			// channel <- p

			board.Play(p.Player, p.Row, p.Col)

			str, _ := json.Marshal(board)
			err = conn.WriteMessage(websocket.TextMessage, str)
		}
	})

	// port 8080 does not work on Mac, possibly due to parental controls
	// https://discussions.apple.com/thread/253069437
	http.ListenAndServe(":8081", nil)
}
