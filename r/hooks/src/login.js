import React from "react";
import { useFormik } from "formik";

export default function Login() {
    const formik = useFormik({
        initialValues: {
            name: "",
            email: "",
        },
        onSubmit: values => {
            alert(JSON.stringify(values, null, 2));
        }
    });

    return (
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
            <br/><br/>
            <label>
                ENTER EMAIL:
                <input
                    type="text"
                    name="email"
                    onChange={formik.handleChange}
                    value={formik.values.email}
                />
            </label>
            <br />            <br/><br/>

            <button type="submit">ENTER</button>
        </form>
    );
}
