import { useState, useEffect } from "react";
import { cellClass, cellHorizontalLine, cellVerticalLine } from "./utils";
import "./App.css";

function App() {
  const [board, setBoard] = useState([]);
  const [socket, setSocket] = useState(null);

  useEffect(() => {
    var socket = new WebSocket("ws://localhost:8081/play");
    // TODO: implement socket.onopen and socket.onclose for reconnect logic
    socket.onmessage = function (e) {
      setBoard(JSON.parse(e.data));
    };
    setSocket(socket);
  }, []);

  function click(row, col) {
    socket.send(JSON.stringify({ row: row, col: col }));
  }

  return (
    <>
      <header>
        <h1>
          <span className="heading-go-one">Go</span>
          <span className="heading-in">in</span>
          <span className="heading-go-two">Go</span>
        </h1>
      </header>
      <main>
        <div
          className="board"
          style={{
            gridTemplateColumns: `repeat(${board.length}, 1fr)`,
            gridTemplateRows: `repeat(${board.length}, 1fr)`,
          }}>
          {board.map((playerIds, row) => {
            return (
              <div key={row} className="row">
                {playerIds.map((playerId, col) => {
                  return (
                    <div
                      key={col}
                      className={cellClass(playerId)}
                      onClick={() => click(row, col)}>
                      <div
                        className="cell-horizontal-line"
                        style={cellHorizontalLine(col, board)}></div>
                      <div
                        className="cell-vertical-line"
                        style={cellVerticalLine(row, board)}></div>
                      {playerId}
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
