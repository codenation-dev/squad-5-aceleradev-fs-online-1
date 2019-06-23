import React, { Component, Fragment } from 'react'
import { withRouter } from "react-router-dom"
import Label from '../../componentes/Label/Label'
import Botao from '../../componentes/Botao/Botao'
import Table from '../../componentes/Table/Table'
import * as apiAlertas from '../../apis/alertas'
import './AlertasInfo.css'

class AlertasInfo extends Component {
    constructor(props) {
        super(props)
        this.state = {
            carregando: true,
            alerta: {}
        }
        this.load()
    }

    load = () => {
        apiAlertas.getIdAlert(this.props.id)
            .then(response => {
                this.setState({ alerta: response.data })
                this.setState({ carregando: false })
            })
            .catch(error => {
                if (error.response) {
                    alert(error.response.data.erro)
                }
            })
    }

    render() {
        const cabecalho = [
            "Username",
            "Name",
            "Email"
        ]
        return (
            <div className="container">
                <section className="card">
                    {!this.state.carregando ?
                        <Fragment>
                            < fieldset >
                                <legend className="legend"> Cliente: </legend>
                                <Label> Tipo: <p> {this.state.alerta.type} </p>
                                </Label>
                                <Label> Descrição: <p> {this.state.alerta.description} </p>
                                </Label>
                                <Label> Nome: <p> {this.state.alerta.public_agent.name} </p>
                                </Label>
                                <Label> Cargo: <p> {this.state.alerta.public_agent.occupation} </p>
                                </Label>
                                <Label> Departamento: <p> {this.state.alerta.public_agent.department} </p>
                                </Label>
                                <Label> Salário: <p> {this.state.alerta.public_agent.salary} </p>
                                </Label>
                            </fieldset>

                            <div>
                                <p> Usuários que receberam alerta: </p>
                                <Table usuarios={this.state.alerta.users_received} length={this.state.alerta.users_received.length} paginacao="2" cabecalho={cabecalho} />                               
                            </div>

                            <div className="rigth">
                                <Botao
                                    classe="button-paginacao"
                                    click={() => {
                                        this.props.history.push('/alertas')
                                    }}>Voltar</Botao>
                            </div>
                        </Fragment>

                        : <p> carregando... </p>}
                </section>
            </div>
        )
    }
}

export default withRouter(AlertasInfo)
