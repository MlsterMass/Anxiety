import React, {Component} from 'react';
import axios from "axios";
class Api extends Component {
    constructor(props) {
        super(props)
        this.state = {
            posts:[]
        }
    }
    apiCall(){
        axios.get()
    }
    render() {
        return (
            <div onClick={this.apiCall}>

            </div>
        );
    }
}

export default Api;