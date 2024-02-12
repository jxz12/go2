# Goception
Asynchronous multiplayer go

# TODO list
* Currently the websocket goroutine blocks on reading a new move
  * we should have a 'broadcaster' goroutine which listens to a channel for moves, then fans out new state
  * see chat example here: https://github.com/gorilla/websocket/blob/main/examples/chat/hub.go
* Currently the user chooses which player id to play as, this should be automatically assigned on connect
* Currently the game does not end when no more valid moves are possible
* Currently the UI does not display the score