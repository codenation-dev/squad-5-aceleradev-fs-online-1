import React from 'react'
import './FormClientes.js'
import Label from '../Label/Label'
import Botao from '../Botao/Botao'
import * as apiClientes from '../../apis/clientes'

const FormClientes = () => {
    let nome = React.createRef()
    let salario = React.createRef()
    const cadastrarCliente = (e) => {
        e.preventDefault()
        const payload = {
            name: nome.value.toUpperCase(),
            salary: parseFloat(salario.value) || 0
        }
        apiClientes.postCustomer(payload)
        .then(response => {
            alert('Cliente cadastrado com sucesso.')
          })
          .catch(error => {
            if (error.response) {
              alert(error.response.data[0].message)
            }
          })
    }

    return (
        <div className="card">
            <div className="card-ordenar">
                {/* <Label>Ordenar:
            </Label> */}
                <form className="card-ordenar" onSubmit={cadastrarCliente}>
                    <Label>
                        Nome:  <input ref={(e) => nome = e}></input>
                    </Label>
                    <Label>
                        Salário: <input ref={(e) => salario = e}></input>
                    </Label>
                    <Botao classe={'paginacao'}>Cadastrar</Botao>

                </form>
            </div>
        </div>

    )
}



export default FormClientes;
