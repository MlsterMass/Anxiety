import './App.css';
import React, {Component} from "react";
import Menu from "./components/menu";
import AddSlider from "./components/addSlider";
import Services from "./components/Services";
import About from "./components/about";
import Specialists from "./components/specialists";
import Footer from "./components/footer";
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Home from "./components/home";
import UserRegistration from "./forms/userRegistration";

class App extends Component{
    render() {
        return (
            <div>
                <Router>
                    <Routes>
                        <Route path="/" element={<Home />} />
                        <Route path="/user_registration" element={<UserRegistration />} />
                    </Routes>
                </Router>
            </div>

        );
    }
}

export default App;
