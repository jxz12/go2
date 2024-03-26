package main

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

// subscribe to the client
// publish to the client

type IBoard interface {
	// Andrea says "tl;dr move the interface to the place where you use it"
	Width() int
	Get(int, int) int
	ToString() string
	Score() map[int]int
	Play(int, int, int) bool
}

type Move struct {
	Player int
	Row    int
	Col    int
}

type Hub struct {
	board   IBoard
	sub     chan Move
	players map[*Player]bool
}

type Player struct {
	id   int
	conn *websocket.Conn
	pub  chan []byte
	hub  *Hub
}

func NewHub(boardSize int) *Hub {
	hub := &Hub{
		board:   NewBoard(boardSize),
		players: make(map[*Player]bool),
		sub:     make(chan Move),
	}
	go hub.Fanout()
	return hub
}
func (hub *Hub) NewPlayer(conn *websocket.Conn, id int) *Player {
	player := &Player{
		id:   id,
		conn: conn,
		hub:  hub,
		pub:  make(chan []byte),
	}
	hub.players[player] = true

	go player.Sub()
	go player.Pub()

	// TODO: this reference to hub.board may require a mutex
	//       this is probably why the gorilla chat example has channels for register and unregister
	//       https://github.com/gorilla/websocket/blob/main/examples/chat/hub.go
	str, _ := json.Marshal(hub.board)
	player.pub <- str
	return player
}

func (player *Player) Sub() {
	// TODO: not sure if I actually need to clean up since Golang has GC
	// defer func() {
	// 	if _, ok := player.hub.players[player]; ok {
	// 		delete(player.hub.players, player)
	// 		close(player.pub)
	// 	}
	// 	player.conn.Close()
	// }()
	for {
		msgType, message, err := player.conn.ReadMessage()
		if err != nil {
			fmt.Println(msgType, message, err)
			// if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
			// 	log.Printf("error: %v", err)
			// }
			break
		}
		var move Move
		err = json.Unmarshal(message, &move)

		// the client should not need to be aware of its own id
		move.Player = player.id

		if err != nil {
			fmt.Println(err)
			// TODO: return error to client
		}
		player.hub.sub <- move
	}
}
func (player *Player) Pub() {
	// defer func() {
	// 	if _, ok := player.hub.players[player]; ok {
	// 		delete(player.hub.players, player)
	// 	}
	// 	// TODO: not sure if I actually need to clean up since Golang has GC
	// 	// player.conn.Close()
	// }()
	for {
		message := <-player.pub
		err := player.conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			return
		}
	}
}

func (hub *Hub) Fanout() {
	for {
		move := <-hub.sub
		hub.board.Play(move.Player, move.Row, move.Col)
		str, _ := json.Marshal(hub.board)
		for player := range hub.players {
			select {
			case player.pub <- str:
			default:
				// close(player.pub)
				delete(hub.players, player)
			}
		}
	}
}
