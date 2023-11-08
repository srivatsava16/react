import React, { useState } from "react";
import { useFormik } from "formik";
import { useParams } from "react-router-dom";
import './styles.css';


export  function App() {
    const { name } = useParams(); // Get the "name" parameter from the URL

    const formik = useFormik({
        initialValues: {
            PlayerOne: name, // Set the value of PlayerOne to the received name
            PlayerTwo: "",
        },
        onSubmit: (values) => {
            alert(JSON.stringify(values, null, 2));
        },
    });

    const inputStyle = {
        marginBottom: "20px",
    };

    return (
        <>
            <div className='main'>
                <form>
                    <div>
                        <label>
                            PLAYER 1:
                            <input
                                style={inputStyle}
                                type="text"
                                name="PlayerOne"
                                onChange={formik.handleChange}
                                value={localStorage.getItem("dataa")}
                            />
                        </label>

                        <br />

                        <label>
                            PLAYER 2:
                            <select
                                style={inputStyle}
                                name="PlayerTwo"
                                onChange={(e) => {
                                    formik.setFieldValue("PlayerTwo", e.target.value);
                                    localStorage.setItem('second', e.target.value)

                                }}
                                value={formik.values.PlayerTwo}
                            >

                                <option  value = "bot1" label="bot1" />
                                <option  value = "bot2" label="bot2" />
                                <option  value = "bot3" label="bot3" />
                            </select></label>

                    </div>
                    <br /><br />

                </form>

                <div className='puz'>
                    <Game/>
                </div>
            </div>
        </>


    );
}

function Square({ value, onSquareClick }) {
    return (
        <button className="square" onClick={onSquareClick}>
           {value}
        </button>
    );
}

function Board({ xIsNext, squares, onPlay }) {
    function handleClick(i) {
        if (calculateWinner(squares) || squares[i]) {
            return;
        }
        const nextSquares = squares.slice();
        if (xIsNext) {
            nextSquares[i] = 'X';
        } else {
            nextSquares[i] = 'O';
        }
        onPlay(nextSquares);
    }

    const winner = calculateWinner(squares);
    let status;
    if (winner) {
        const t=localStorage.getItem('dataa')
        const x=localStorage.getItem('second')

        if(winner==="X") {
            status = 'Winner: ' + t;
        }else {
            status= 'Winner: ' + x;
        }
    } else {
        status = 'Next player: ' + (xIsNext ? 'X' : 'O');
    }

    return (
        <>
            <div className="status">{status}</div>
            <div className="board-row">
                <Square value={squares[0]} onSquareClick={() => handleClick(0)} />
                <Square value={squares[1]} onSquareClick={() => handleClick(1)} />
                <Square value={squares[2]} onSquareClick={() => handleClick(2)} />
            </div>
            <div className="board-row">
                <Square value={squares[3]} onSquareClick={() => handleClick(3)} />
                <Square value={squares[4]} onSquareClick={() => handleClick(4)} />
                <Square value={squares[5]} onSquareClick={() => handleClick(5)} />
            </div>
            <div className="board-row">
                <Square value={squares[6]} onSquareClick={() => handleClick(6)} />
                <Square value={squares[7]} onSquareClick={() => handleClick(7)} />
                <Square value={squares[8]} onSquareClick={() => handleClick(8)} />
            </div>
        </>
    );
}

function Game() {
    const [history, setHistory] = useState([Array(9).fill(null)]);
    const [currentMove, setCurrentMove] = useState(0);
    const xIsNext = currentMove % 2 === 0;
    const currentSquares = history[currentMove];

    function handlePlay(nextSquares) {
        const nextHistory = [...history.slice(0, currentMove + 1), nextSquares];
        setHistory(nextHistory);
        setCurrentMove(nextHistory.length - 1);
    }

    function jumpTo(nextMove) {
        setCurrentMove(nextMove);
    }

    const moves = history.map((squares, move) => {
        let description;
        if (move > 0) {
            description = 'Go to move #' + move;
        } else {
            description = 'Go to game start';
        }
        return (
            <div>
            <li key={move}>
                <button onClick={() => jumpTo(move)}>{description}</button>
            </li>
            </div>
        );
    });

    return (
        <div className="game">
            <div className="game-board">
                <Board xIsNext={xIsNext} squares={currentSquares} onPlay={handlePlay} />
            </div>
            <div className="game-info">
                <ol>{moves}</ol>
            </div>
        </div>
    );
}

function calculateWinner(squares) {
    const lines = [
        [0, 1, 2],
        [3, 4, 5],
        [6, 7, 8],
        [0, 3, 6],
        [1, 4, 7],
        [2, 5, 8],
        [0, 4, 8],
        [2, 4, 6],
    ];
    for (let i = 0; i < lines.length; i++) {
        const [a, b, c] = lines[i];
        if (squares[a] && squares[a] === squares[b] && squares[a] === squares[c]) {
            return squares[a];
        }
    }
    return null;
}
