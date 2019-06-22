import React from 'react'
import Td from '../Td/Td'
import { Link } from 'react-router-dom'
import './TBody.css'

const TBody = ({ alertas, usuarios }) => (

    <tbody>
        {alertas ?        
            alertas.map(alerta => {
                return (
                    <tr key={alerta.id}>
                        <Td>{alerta.type}</Td>
                        <Td id="name">{alerta.customer_name}</Td>
                        <Td>{alerta.datetime}</Td>
                        <Td>
                            <Link className="link-mais_info" to={{
                                pathname: `alertas/${alerta.id}`,
                                state: {
                                    id: alerta.id
                                }
                            }}>+ informações</Link>
                        </Td>
                    </tr>
                )
            }) : usuarios ?
                usuarios.map(user => (
                    <tr key={user.id}>
                        <Td> {user.username} </Td>
                        <Td> {user.name} </Td>
                        <Td> {user.email} </Td>
                    </tr>
                ))


                : ''}



    </tbody>
)
export default TBody
