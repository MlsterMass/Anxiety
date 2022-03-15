import React, {Component} from 'react';

class Services extends Component {
    constructor(props) {
        super(props);

    }

    render() {
        return (
            <div>
                <h1>Сервисы</h1>
                <div className="services">
                    <div className="row">
                        <div className="column"><h3>Ментальное здоровье</h3>
                        Контент
                        </div>
                        <div className="column"><h3>Реабилитация</h3>
                            Контент
                        </div>
                        <div className="column"><h3>Безопасность</h3>
                            Контент
                        </div>

                    </div>
                    <div className="row">
                        <div className="column"><h3>Транспорт</h3>
                            Контент
                        </div>
                        <div className="column"><h3>Поиск</h3>
                            Контент
                        </div>
                        <div className="column"><h3>О нас</h3>
                            Контент
                        </div>
                    </div>
                </div>
            </div>
        );
    }
}

Services.propTypes = {};

export default Services;