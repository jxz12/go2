import { useState, useEffect } from "react";
import { cellClass, cellHorizontalLine, cellVerticalLine } from "./utils";
import "./App.css";

function App() {
  const [board, setBoard] = useState([]);
  const [socket, setSocket] = useState(null);
  const [player, setPlayer] = useState(1);

  useEffect(() => {
    console.log("connecting");
    var socket = new WebSocket("ws://localhost:8081/play");
    console.log("connected");
    socket.onmessage = function (e) {
      setBoard(JSON.parse(e.data));
    };
    setSocket(socket);
  }, []);

  function click(i, j) {
    socket.send(JSON.stringify({ player: player, row: i, col: j }));
  }

  function handlePlayerChange(event) {
    setPlayer(Number(event.target.value));
  }

  return (
    <>
      <header>
        <h1>
          <span className="heading-go-one">Go</span>
          <span className="heading-in">in</span>
          <span className="heading-go-two">Go</span>
        </h1>
        <form>
          <label>
            <div className="cell player-one">1</div>
            <input
              type="checkbox"
              value={1}
              checked={player === 1}
              onChange={handlePlayerChange}
            />
          </label>
          <label>
            <div className="cell player-two">2</div>
            <input
              type="checkbox"
              value={2}
              checked={player === 2}
              onChange={handlePlayerChange}
            />
          </label>
        </form>
      </header>
      <main>
        <div
          className="board"
          style={{
            gridTemplateColumns: `repeat(${board.length}, 1fr)`,
            gridTemplateRows: `repeat(${board.length}, 1fr)`,
          }}>
          {board.map((row, i) => {
            return (
              <div key={i} className="row">
                {row.map((cell, j) => {
                  return (
                    <div
                      key={j}
                      className={cellClass(cell)}
                      onClick={() => click(i, j)}>
                      <div
                        className="cell-horizontal-line"
                        style={cellHorizontalLine(j, board)}></div>
                      <div
                        className="cell-vertical-line"
                        style={cellVerticalLine(i, board)}></div>
                      {cell}
                    </div>
                  );
                })}
              </div>
            );
          })}
        </div>
      </main>
    </>
  );
}

export default App;
