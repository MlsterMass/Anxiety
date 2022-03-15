import React, {Component} from 'react';

class Menu extends Component {
    constructor(props) {
        super(props);
        this.state = {
            value:'user'
        };
        this.handleUserType = this.handleUserType.bind(this)
    }

    handleUserType(event){
        this.setState({value: event.target.value});
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
            <span>Anxiety bot TG   </span>
            <span>+38098000000</span>
            <li className="registration">
                <form action={`${this.state.value}_registration`}>
                    <button type="submit">Регистрация</button>
                    <select defaultValue={this.state.value} onChange={this.handleUserType}>
                        <option value="user">Пользователь</option>
                        <option value="volunteer">Волонтер</option>
                        <option value="specialist">Специалист</option>
                        <option value="support">Поддержка</option>
                    </select>
                </form>
            </li>
        </ul>
        );
    }
}

export default Menu;