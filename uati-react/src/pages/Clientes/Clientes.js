import React, { useEffect, useState, Fragment } from 'react'
import './clientes.css';
import * as apiClientes from '../../apis/clientes'
import Table from '../../componentes/Table/Table'
import Ordenar from '../../componentes/Ordernar/Ordenar'
import FormClientes from '../../componentes/FormClientes/FormClientes'
import Navbar from '../../componentes/Navbar/Navbar'
import { estAutenticado } from '../../routes'
import Botao from '../../componentes/Botao/Botao';

function Clientes() {
  const [length, setLength] = useState(0)
  const [clientes, setClientes] = useState([])
  const [filtro, setFiltro] = useState({})
  const [cadastrar, setCadastrar] = useState(false)
  const [filtrar, setFiltrar] = useState(false)
  const lista = [
    {
      name: 'Nome',
      column: 'name'
    },
    {
      name: 'Salário',
      column: 'salary'
    }
  ]
  const paginacao = 10
  const cabecalho = [
    "Nome",
    "Salário",
    "Ações"
  ]


  useEffect(() => {
    apiClientes.getCustomers()
      .then(response => {
        setLength(response.data.data.length)
        //setCarregando(false)
      })
      .catch(error => {
        if (error.response) {
          alert(error.response.data[0].message)
        }
      })
  }, [])

  return (
    <>
      {estAutenticado() ? <Navbar></Navbar> : ''}
      <div className="container">
        <section className="card">
          {/* {!carregando ?   */}
          <Fragment>
            <div className="buttons">
              <Botao
                classe="clientes"
                click={() => {
                  setFiltrar(true)
                  setCadastrar(false)
                }}>Pesquisar</Botao>
              <Botao
                classe="clientes"
                click={() => {
                  setFiltrar(false)
                  setCadastrar(true)
                }}>Cadastrar</Botao>
            </div>
            {cadastrar ? <FormClientes></FormClientes> : ''}
            {filtrar ? <Ordenar lista={lista} setFiltro={setFiltro}></Ordenar> : ''}
            <Table cabecalho={cabecalho} length={length} filtro={filtro} clientes={clientes} setClientes={setClientes} paginacao={paginacao} setLength={setLength} />
          </Fragment>
          {/* : 'carregando...'} */}
        </section>
      </div>
    </>
  )
}

export default Clientes;
