package main

import (
	"fmt"
    "net/http"
    "github.com/gorilla/websocket"
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

    http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
        conn, err := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("connected")
        for {
            // Read message from browser
            msgType, msg, err := conn.ReadMessage()
            if err != nil {
				fmt.Printf("Error: %s\n", err)
                // return
            }

            // Print the message to the console
            fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

            // Write message back to browser
            if err = conn.WriteMessage(msgType, []byte("HELLO")); err != nil {
                return
            }

            // Write message back to browser
            if err = conn.WriteMessage(msgType, msg); err != nil {
                return
            }
        }
    })
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "websockets.html")
    })
    http.ListenAndServe(":8080", nil)
    // http.ListenAndServe("localhost:8080", nil)
    // http.ListenAndServe("127.0.0.1:8080", nil)
}