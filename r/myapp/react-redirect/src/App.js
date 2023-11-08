import React from 'react';
import { Link } from 'react-router-dom';
import './App.css'; // Import the CSS file

window.__REACT_DEVTOOLS_GLOBAL_HOOK__ = { isDisabled: true };


function App() {
    return (
        <div className="main-page">
            <div className="button-container">
                <Link to="/login">
                    <button>Log In</button>
                </Link>
                <Link to="/register">
                    <button>Sign Up</button>
                </Link>
            </div>
        </div>
    );
}

export default App;
