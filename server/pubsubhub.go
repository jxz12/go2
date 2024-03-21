package main

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

const SUB_BUF = 8
const PUB_BUF = 64

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
	// TODO: maybe this is cleaner instead of reference to players?
	// pubs map[chan<- Move]bool
}

type Player struct {
	conn *websocket.Conn
	pub  chan []byte
	hub  *Hub
	// TODO: maybe this is cleaner instead of reference to hub?
	// sub <-chan Move
}

func NewHub(boardSize int) *Hub {
	return &Hub{
		board:   NewBoard(19),
		players: make(map[*Player]bool),
		sub:     make(chan Move, SUB_BUF),
	}
}
func (hub *Hub) NewPlayer(conn *websocket.Conn) *Player {
	player := &Player{
		conn: conn,
		hub:  hub,
		pub:  make(chan []byte, PUB_BUF),
	}
	hub.players[player] = true
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
