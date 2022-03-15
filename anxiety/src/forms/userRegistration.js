import React, {Component} from 'react';
import axios from "axios";

class UserRegistration extends Component {
    constructor(props) {
        super(props);
        this.state = {
            name:'Жанна',
            nickname:'ghost',
            gender:'Женский',
            status:'Введите информацию о вашем состоянии',
            children: 0,
            pets: 0,
            location:'Харьков',
            password:'qwerty',
        };
    }

    createUser = () => {
        axios.post(
            `/registration/user?name=${this.state.name}&nickname=${this.state.nickname}&gender=${this.state.gender}&status=${this.state.status}&children=${this.state.children}&pets=${this.state.pets}&location=${this.state.location}&password=${this.state.password}`,
            "",
            {
            headers: {
                "Content-Type": "application/x-www-form-urlencoded",
                "Access-Control-Allow-Origin": "*",
            },
        })
            .then(response=> {
                console.log(response);
            })
            .catch(error => {
                console.log(error);
            });
    }
    handleName = event => {
        this.setState({
            name: event.target.value
        });
    }
    handleNickname = event => {
        this.setState({
            nickname: event.target.value
        });
    }
    handleGender = event => {
        this.setState({
            gender: event.target.value
        });
    }
    handleStatus = event => {
        this.setState({
            status: event.target.value
        });
    }
    handleChildren = event => {
        this.setState({
            children: 1
        });
    }
    handlePets = event => {
        this.setState({
            pets: 1
        });
    }
    handleLocation = event => {
        this.setState({
            location: event.target.value
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
                <h1>Регистрация пользователя</h1>
                <form onSubmit={this.createUser}>
                    <div>
                        <label/>Имя
                        <input type="text"
                               value={this.state.name}
                               onChange={this.handleName}
                        />
                    </div>
                    <div>
                        <label/>Никнейм
                        <input type="text"
                               value={this.state.nickname}
                               onChange={this.handleNickname}
                        />
                    </div>
                    <div>
                        <label/>Пол
                        <input type="text"
                               value={this.state.gender}
                               onChange={this.handleGender}
                        />
                    </div>
                    <div>
                        <label/>Статус
                        <textarea
                            value={this.state.status}
                            onChange={this.handleStatus}
                        />
                    </div>
                    <div>
                        <label/>Дети
                        <input type="radio"
                               checked={this.state.children}
                               onChange={this.handleChildren}
                        />
                    </div>
                    <div>
                        <label/>Питомцы
                        <input type="radio"
                               checked={this.state.pets}
                               onChange={this.handlePets}
                        />
                    </div>
                    <div>
                        <label/>Локация
                        <input type="text"
                               value={this.state.location}
                               onChange={this.handleLocation}
                        />
                    </div>
                    <div>
                        <label/>Пароль
                        <input type="password"
                               value={this.state.password}
                               onChange={this.handlePassword}
                        />
                    </div>
                    <button type="submit">Зарегистрироваться</button>
                </form>
            </div>
        );
    }
}

export default UserRegistration;