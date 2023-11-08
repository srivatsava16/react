import React from 'react';
import { Formik, Form, useField } from 'formik';
import * as Yup from 'yup';

// Create a Text component for input fields
function Text({ label, ...props }) {
    const [field, meta] = useField(props);

    return (
        <div>
            <label htmlFor={props.id || props.name}>{label}: </label>
            <input {...field} {...props} />
            {meta.touched && meta.error && <div className="error">{meta.error}</div>}
        </div>
    );
}

export default function App() {
    return (
        <div>
            <h1>Formik and Yup Example</h1>
            <Formik
                initialValues={{
                    Name: '',
                }}
                validationSchema={Yup.object({
                    Name: Yup.string()
                        .min(5, 'Must be 5 characters or more')
                        .required('Required'),
                })}
                onSubmit={(values, { setSubmitting }) => {
                    setTimeout(() => {
                        alert(JSON.stringify(values, null, 2));
                        setSubmitting(false);
                    }, 400);
                }}
            >
                <Form>
                    <div>
                        <Text label="Name" name="Name" type="text" />
                    </div>
                    <br></br>
                    <div>
                        <button type="submit">Submit</button>
                    </div>
                </Form>
            </Formik>
        </div>
    );
}
