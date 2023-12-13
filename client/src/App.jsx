import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'

const BOARD_WIDTH = 9;
const BOARD_START = Array(BOARD_WIDTH).fill(Array(BOARD_WIDTH).fill(0));

function App() {
  const [board, setBoard] = useState([[0,0,0],[0,0,0],[0,0,0]]);

  function click(i, j) {
    // board[i][j] = 1;
    // setBoard(board.map(row => row.map(cell => cell)));

    setBoard(prevBoard => {
      prevBoard[i][j] = 1;
      return prevBoard.map(row => row.map(cell => cell));
    });
  }

  return (
    <>
      <table>
        <tbody>
          {
            board.map((row, i) => {
              return (
                <tr key={i}>
                  {
                    row.map((cell, j) => {
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
