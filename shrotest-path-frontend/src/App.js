import { useState, useCallback } from 'react';
import './App.css';

const GRID_SIZE = 20;

function App() {
  const [start, setStart] = useState(null); 
  const [stop, setStop] = useState(null);  

  const handleCellClick = useCallback((x, y) => {
    if (!start) {
      setStart({ x, y });
    } else if (!stop) {
      setStop({ x, y });
    } else {
      setStart({ x, y });
      setStop(null);
    }
  }, [start, stop]);

  return (
    <div
      className="App"
      style={{
        width: '100vw',
        height: '100vh',
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center',
      }}
    >
      <div
        style={{
          display: 'grid',
          gridTemplateColumns: `repeat(${GRID_SIZE}, 1fr)`,
        }}
      >
        {Array.from({ length: GRID_SIZE }).map((_, x) =>
          Array.from({ length: GRID_SIZE }).map((_, y) => (
            <SingleGridCell
              key={`${x}-${y}`}
              x={x}
              y={y}
              isSelected={
                (start?.x === x && start?.y === y) ||
                (stop?.x === x && stop?.y === y)
              }
              onSelectHandler={handleCellClick}
            />
          ))
        )}
      </div>
    </div>
  );
}

function SingleGridCell({ x, y, isSelected, onSelectHandler }) {
  const handleClick = useCallback(() => {
    onSelectHandler(x, y);
  }, [x, y, onSelectHandler]);

  return (
    <div
      onClick={handleClick}
      style={{
        border: '1px solid black',
        width: '40px',
        height: '40px',
        cursor: 'pointer',
        backgroundColor: isSelected ? 'lightblue' : 'white',
      }}
    />
  );
}

export default App;
