import { useState, useEffect } from 'react'
import './App.css'

function App() {
  const [board, setBoard] = useState([]);
  const [socket, setSocket] = useState(null);
  const [player, setPlayer] = useState(1)

  useEffect(() => {
    console.log("connecting");
    var socket = new WebSocket("ws://localhost:8081/play");
    console.log("connected");
    socket.onmessage = function (e) {
      setBoard(JSON.parse(e.data));
    };
    setSocket(socket);
  }, [])

  function click(i, j) {
    socket.send(JSON.stringify({player: player, row: i, col: j}));
  }

  return (
    <>
      <p>player 1 or 2: <input type="checkbox" onChange={() => setPlayer((player % 2) + 1)}/></p>
      <table>
        <tbody>
          {
            board.map((row, i) => {
              return (
                <tr key={i}>
                  {
                    row.map((cell, j) => {
                      // TODO: use emojis hehe
                      return <td key={j}><div onClick={() => click(i, j)}>&nbsp;{cell}&nbsp;</div></td>
                    })
                  }
                </tr>
              )
            })
          }
        </tbody>
      </table>
    </>
  )
}

export default App
