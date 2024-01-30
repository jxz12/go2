import { useState, useEffect } from 'react'
import './App.css'

const BOARD_WIDTH = 9;
const BOARD_START = Array(BOARD_WIDTH).fill(Array(BOARD_WIDTH).fill(0));

function App() {
  const [board, setBoard] = useState([[0, 0, 0], [0, 0, 0], [0, 0, 0]]);
  const [socket, setSocket] = useState(null);

  useEffect(() => {
    console.log("connecting");
    var socket = new WebSocket("ws://localhost:8080/echo");
    console.log("connected");
    socket.onmessage = function (e) {
      console.log(e.data);
    };
    setSocket(socket);
  }, [])

  function click(i, j) {
    setBoard(prevBoard => {
      prevBoard[i][j] = 1;
      return prevBoard.map(row => row.map(cell => cell));
    });
  }

  return (
    <>
      <button onClick={() => socket.send("hello")}>socket</button>
      <table>
        <tbody>
          {
            board.map((row, i) => {
              return (
                <tr key={i}>
                  {
                    row.map((cell, j) => {
                      // TODO: use emojis hehe
                      return <td key={j}><div onClick={() => click(i, j)}>{cell}</div></td>
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
