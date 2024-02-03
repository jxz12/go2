# Goception
Asynchronous multiplayer go

# TODO list
* Currently the websocket goroutine blocks on reading a new move
  * we should have a 'game controller' goroutine which listens to a channel for moves, then fans out new state
* Currently the user chooses which player id to play as, this should be automatically assigned on connect
* The UI is ugly, we can use either coloured counters or emojis 