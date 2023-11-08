import React from "react";
import { useFormik } from "formik";
import { useNavigate } from "react-router-dom";

export default function Login() {
    const navigate = useNavigate();

    const formik = useFormik({
        initialValues: {
            name: "",
            email: "",
        },
        onSubmit: values => {
            localStorage.setItem('dataa', values.name);
             navigate("/app");
        }
    });

    return (
        <div >
            <h1>LOGIN TO PLAY THE GAME</h1>

        <form onSubmit={formik.handleSubmit}>
            <label>
                ENTER NAME:
                <input
                    type="text"
                    name="name"
                    onChange={formik.handleChange}
                    value={formik.values.name}
                />
            </label>
            <br /><br />
            <label>
                ENTER EMAIL:
                <input
                    type="text"
                    name="email"
                    onChange={formik.handleChange}
                    value={formik.values.email}
                />
            </label>
            <br /><br />

            <button type="submit">ENTER</button>
        </form>
            </div>
    );
}


