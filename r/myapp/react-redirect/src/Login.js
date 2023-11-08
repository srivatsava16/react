import { useNavigate } from "react-router-dom";
import { useState } from "react";

function Login() {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState(""); // Add state for password
    const navigate = useNavigate();

    const handleSubmit = async (e) => {
        e.preventDefault();

        const data = {
            username: username,
            password: password, // Include the password
        };

        try {
            const response = await fetch("/auth", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(data),
            });

            if (response.ok) {
                // Successful authentication
                navigate("/dashboard");
            } else {
                // Authentication failed
                console.error("Authentication failed");
            }
        } catch (error) {
            console.error("Network error:", error);
        }
    };

    return (
        <div className="App">
            <p><b>ENTER YOUR CREDENTIALS</b></p>
            <form className="Inp" onSubmit={handleSubmit}>
                <label>NAME:
                    <input
                        type="text"
                        value={username}
                        onChange={(e) => setUsername(e.target.value)}
                    />
                </label><br/><br/>

                <label>PASSWORD:
                    <input
                        type="password" // Change the input type to "password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                    />
                </label><br/><br/>

                <input type="submit" value="Submit" />
            </form>
        </div>
    );
}

export default Login;
