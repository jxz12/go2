export function cellClass(cell) {
  switch (cell) {
    case 0:
      return "cell empty";
    case 1:
      return "cell player-one";
    case 2:
      return "cell player-two";
  }
}

export function cellHorizontalLine(j, board) {
  const style = { top: "50%", width: "100%" };
  if (j === 0) {
    style.left = "50%";
  } else if (j === board.length - 1) {
    style.left = "-50%";
  }
  return style;
}

export function cellVerticalLine(i, board) {
  const style = { left: "50%", height: "100%" };
  if (i === 0) {
    style.top = "50%";
  } else if (i === board.length - 1) {
    style.top = "-50%";
  }
  return style;
}
