import React, { useState } from "react";
import "./App.css";
import {useNavigate} from "react-router-dom";
import axios from "axios";

function Register() {

    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");

   const navigate = useNavigate();

    function handleSubmit(event) {
        event.preventDefault(); 
        const baseURL = "http://localhost:8080/create";
        const payload = JSON.stringify({ Username: username, Password: password });

        axios.post(baseURL, payload).then((response) => {
            // Handle the response here
            navigate("/login")
        });
    }


    return (
        <div className="register-container">
            <p id="n">CREATE NEW ACCOUNT</p>
            <form className="Inp" onSubmit={handleSubmit}>
                <label className="form-label">NAME:</label>
                <input
                    className="form-input"
                    type="text"
                    value={username}
                    onChange={(e) => setUsername(e.target.value)}
                />
                <br/><br/>
                <label className="form-label">PASSWORD:</label>
                <input
                    className="form-input"
                    type="text"
                    value={password}
                    onChange={(e)=>setPassword(e.target.value)}
                />
                <br/><br/>
                <input className="submit-button" type="submit" value="Submit" />
            </form>
        </div>
    );
}

export default Register;





// fetch("http://localhost:8080/create", {
//     method: 'POST',
//     headers: {
//         'Content-Type': 'application/json',
//     },
//     body: JSON.stringify({ Username: 'username',Password:"password" }), // Replace with your data
// })
//     .then(response => response.json()).catch(err => console.log(err))
// .then(data => {
//     console.log(data);
//     navigate("/login");
// })
// .catch(error => {
//     // Handle any errors
//     console.error(error);
// });