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