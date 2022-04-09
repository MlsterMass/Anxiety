import './App.css';
import React, {Component} from "react";
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Home from "./components/home";
import Login from "./forms/login";
import UserRegistration from "./forms/userRegistration";

class App extends Component{
    render() {
        return (
            <div>
                <Router>
                    <Routes>
                        <Route path="/" element={<Home />} />
                        <Route path="/user_registration" element={<UserRegistration />} />
                        {/* TODO: init endpoints
                        <Route path="/volunteer_registration" element={<UserRegistration />} />
                        <Route path="/specialist_registration" element={<UserRegistration />} />
                        <Route path="/support_registration" element={<UserRegistration />} />
                        */}
                        <Route path="/login" element={<Login />} />
                    </Routes>
                </Router>
            </div>

        );
    }
}

export default App;
