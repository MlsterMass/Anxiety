import React, {Component} from 'react';
import Menu from "./menu";
import AddSlider from "./addSlider";
import Services from "./Services";
import About from "./about";
import Specialists from "./specialists";
import Footer from "./footer";

class Home extends Component {
    render() {
        return (
            <div>
                <Menu />
                <AddSlider />
                <Services />
                <About />
                <Specialists />
                <Footer />
            </div>
        );
    }
}

export default Home;