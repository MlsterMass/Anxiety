import React, {Component} from 'react';
import axios from "axios";

class Menu extends Component {
    constructor(props) {
        super(props);
        this.state = {
            value:'user',
            user: 'User',
        };
        this.handleUserType = this.handleUserType.bind(this)
    }

    handleUserType(event){
        this.setState({value: event.target.value});
    }
    getUser = () => {
        axios.get(
            "/user",
            {
                headers: {
                    "Content-Type": "multipart/form-data",
                    "Access-Control-Allow-Origin": "*",
                },
            })
            .then(response => {
                this.setState({user: response.data.user});
            })
            .catch(error => {
                console.log(error);
            });
    }
    render() {
        return (
        <ul className="nav">
            <span>Anxiety</span>
            <li><a href="">Сервисы</a></li>
            <li><a href="">Специалисты</a></li>
            <li><a href="">Быстрая помощь</a></li>
            <li><a href="">О нас</a></li>
            <li><a href="">Контакты</a></li>
            <li><span>Anxiety bot TG</span></li>
            <li><span>+38098000000</span></li>
            {/*<li><span className="user" onLoad={this.getUser}>{this.state.user}</span><button>Выход</button></li>*/}
            <span className="registration">
                <form action={`${this.state.value}_api_registration`}>
                    <button type="submit" className="registration_button">Регистрация</button>
                    <select defaultValue={this.state.value} onChange={this.handleUserType}>
                        <option value="user">Пользователь</option>
                        <option value="volunteer">Волонтер</option>
                        <option value="specialist">Специалист</option>
                        <option value="support">Поддержка</option>
                    </select>
                </form>
            </span>
        </ul>
        );
    }
}

export default Menu;