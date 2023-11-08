import React from "react";
import ReactDOM from "react-dom";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import pp, {App} from "./App";
import Lgin from "./Login";

ReactDOM.render(
    <React.StrictMode>
        <BrowserRouter>
            <Routes>
                {/* The "Login" component is rendered at the root URL ("/") */}
                <Route index element={<Lgin />} />

                {/* The "App" component is rendered when the URL matches "/app" */}
                <Route path="app" element={<App />} />
            </Routes>
        </BrowserRouter>
    </React.StrictMode>,
    document.getElementById("root")
);
