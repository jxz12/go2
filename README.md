# Goception
Asynchronous multiplayer go

## Howto Client
```
cd client
npm run dev
```

## Howto Server
```
cd server
go run go2
```

# TODO list
* the UI does not display the score
* the game does not end when no more valid moves are possible
* the UI only chooses between two colours and does not highlight the current user's stones
* there is not reconnect logic (how can we maintain the playerId? Maybe via ip address?)
  * the current logic is actually leaky because players who disconnect do not close their goroutines

* store an expiration time on the server
  * player can hit a button which is "I'm still here"
  * real expiration time is 10 seconds, but only display final 5
  * every state of the board has an expiration time, not just the board state
  * how to handle when someone has left the game?