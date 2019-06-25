import React, { useState } from 'react'
import * as api from '../../apis/dashboards'
import THead from '../../componentes/THead/THead'
import Td from '../../componentes/Td/Td'
import TFoot from '../../componentes/TFoot/TFoot'
import './dashboard.css'

export default class Dashboard extends React.Component { 
    constructor(props) {
        super(props)
        this.state = {
            offset: 0,
            limit: 5,
            customers: [],
            paginacao: 10,
            setStart:0,
            setEnd: 10,
            length: 0,
            usuarios: [],
            cabecalho: [
                "Tipo",
                "Nome",
                "Data",
                "Salário"
            ]
        }
    }

    async componentDidMount() {
        const resp = await api.getLastAlerts(this.state.offset,this.state.limit);
        this.setState({
            customers: resp.data.data
        });
    }

    render() {

        return (
            <div>
                <div>
                    <h1>Dashboard</h1>
                </div>
                <div className="row">
                    <div className="column">
                    <table className="card table">
                        <THead cabecalho={this.state.cabecalho}> </THead>
                        <tbody>
                            {        
                            this.state.customers.map(customer => {
                                return (
                                    <tr key={customer.id}>
                                        <Td>{customer.type}</Td>
                                        <Td>{customer.name}</Td>
                                        <Td>{customer.datetime}</Td>
                                        <Td>{customer.salary}</Td>
                                    </tr>
                                )
                            })}
                        </tbody>
                        <TFoot quantidade={Math.ceil(this.state.length/this.state.paginacao)} paginacao={this.state.paginacao} setEnd={this.state.setEnd} setStart={this.state.setStart}></TFoot>
                    </table>
                    </div>
                    <div className="column"></div>
                </div>
            </div>
        );
    }

} 