package main

import "github.com/gorilla/websocket"

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
	players map[*Player]bool
	pubs    chan []Move
	subs    chan *Player
	unsubs  chan *Player
}

type Player struct {
	hub  *Hub
	conn *websocket.Conn
	pubs chan []byte
}

func NewHub(size int) *Hub {
	return &Hub{
		board:   NewBoard(19),
		players: make(map[*Player]bool),
		pubs:    make(chan []Move),
		subs:    make(chan *Player),
		unsubs:  make(chan *Player),
	}
}

// func (h *Hub) run() {
// 	for {
// 		select {
// 		case client := <-h.register:
// 			h.clients[client] = true
// 		case client := <-h.unregister:
// 			if _, ok := h.clients[client]; ok {
// 				delete(h.clients, client)
// 				close(client.send)
// 			}
// 		case message := <-h.broadcast:
// 			for client := range h.clients {
// 				select {
// 				case client.send <- message:
// 				default:
// 					close(client.send)
// 					delete(h.clients, client)
// 				}
// 			}
// 		}
// 	}
// }
