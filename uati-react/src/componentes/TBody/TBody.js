import React from 'react'
import { Link } from 'react-router-dom'
import Td from '../Td/Td'
import Botao from '../Botao/Botao'
import * as apiClientes from '../../apis/clientes'

import './TBody.css'

class TBody extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            nome: '',
            salary: ''
        }
    }

    editar = () => {
        this.props.setEdit(true)
    }

    salvar = (cliente) => {
        const { nome, salary } = this.state
        if (nome !== '' || salary !== '') {
            const payload = {
                name: nome ? nome : cliente.name,
                salary: salary ? parseFloat(salary) : parseFloat(cliente.salary)
            }
            apiClientes.putCustomer(cliente.id, payload)
                .then(response => {
                    console.log(response)
                })
        }
        this.props.setEdit(false)


    }

    render() {
        const { alertas, clientes } = this.props
        return (
            <tbody>

                {alertas ?
                    alertas.map(alerta => {
                        return (
                            <tr key={alerta.id}>
                                <Td classe="column" >{alerta.type}</Td>
                                <Td id="name" classe="column">{alerta.customer_name}</Td>
                                <Td classe="column">{alerta.datetime.substring(0, 19).replace('T', ' ')}</Td>
                                <Td classe="column">
                                    <Link className="link-mais_info" to={{
                                        pathname: `/alertas/${alerta.id}`,
                                        state: {
                                            id: alerta.id
                                        }
                                    }}>+ informações</Link>
                                </Td>
                            </tr>
                        )
                    }) : <tr></tr>}

                {clientes ?
                    clientes.map(cliente => (
                        <tr key={cliente.id}>
                            <Td id="name" classe="column-clientes">
                                <input
                                    className={!this.props.edit ? "input" : ""}
                                    readOnly={!this.props.edit}
                                    type="text"
                                    defaultValue={cliente.name}
                                    onChange={(e) => this.setState({ nome: e.target.value })}
                                >
                                </input>
                            </Td>
                            <Td classe="column-clientes">
                                <input
                                    className={!this.props.edit ? "input" : ""}
                                    type="text"
                                    readOnly={!this.props.edit}
                                    defaultValue={cliente.salary}
                                    onChange={(e) => this.setState({ salary: e.target.value })}>
                                </input>
                            </Td>
                            <Td> <Botao classe='paginacao' click={this.props.edit ? () => this.salvar(cliente) : this.editar}> {this.state.editar ? 'Salvar' : 'Editar'} </Botao> </Td>


                        </tr>
                    )) : <tr></tr>
                }

            </tbody>
        )
    }
}
export default TBody
