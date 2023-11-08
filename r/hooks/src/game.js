import React from "react";
import { useFormik } from "formik";

export default function Game() {
    const formik = useFormik({
        initialValues: {
            PlayerOne: "",
            PlayerTwo: "",
        },
        onSubmit: values => {
            alert(JSON.stringify(values, null, 2));
        }

    }); // Remove the extra closing parenthesis here

    const inputStyle = {
        marginBottom: "20px", // Add more space between the two fields
    };

    const buttonStyle = {
        fontSize: "24px", // Increase the font size
        // Center-align the text
        padding: "10px 20px",
        marginLeft:"500px",
        // Add some padding to the button
    };



    return (
        <div>
            <h1>ENTER PLAYERS NAMES:</h1>
            <form onSubmit={formik.handleSubmit}>
                <div>
                    <label>
                        PLAYER 1:
                        <input
                            style={inputStyle} // Apply style to the input
                            type="text"
                            name="PlayerOne"
                            onChange={formik.handleChange}
                            value={formik.values.PlayerOne}
                        />
                    </label>
                </div>


                <br />

                <div>
                    <label>
                        PLAYER 2:
                        <select
                            style={inputStyle} // Apply style to the select
                            name="PlayerTwo"
                            onChange={formik.handleChange}
                            value={formik.values.PlayerTwo}
                        >
                            <option value="BOT" label="bot1" />
                            <option value="Player 2" label="bot2" />
                            <option value="Player 3" label="bot3" />
                        </select>
                    </label>

                    <br/><br/>

                    <button >click me</button> {/* Apply the button style */}

                </div>
            </form>
        </div>
    );
}
