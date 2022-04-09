import React, {Component} from 'react';
import axios from "axios";

class Login extends Component {
    constructor(props) {
        super(props);
        this.state = {
            nickname:'ghost',
            password:'qwerty',
        };
    }

    authUser = event => {
        const formData = new FormData();
        event.preventDefault();
        axios.post(
            `/login?nickname=${this.state.nickname}&password=${this.state.password}`,
            formData,
            {
                headers: {
                    "Content-Type": "multipart/form-data",
                    "Access-Control-Allow-Origin": "*",
                },
            })
                .then(response=> {
                    console.log(response);
                    if (response.status === 200) {
                        window.location = "/"
                    }
                })
                .catch(error => {
                    console.log(error);
                });
    }
    handleNickname = event => {
        this.setState({
            nickname: event.target.value
        });
    }
    handlePassword = event => {
        this.setState({
            password: event.target.value
        });
    }
    render() {
        return (
            <div>
                <h1>Авторизация пользователя</h1>
                <form onSubmit={this.authUser}>
                    <div>
                        <label/>Никнейм
                        <input type="text"
                               value={this.state.nickname}
                               onChange={this.handleNickname}
                        />
                    </div>
                    <div>
                        <label/>Пароль
                        <input type="password"
                               value={this.state.password}
                               onChange={this.handlePassword}
                        />
                    </div>
                    <button type="submit">Войти</button>
                </form>
            </div>
        );
    }
}

export default Login;