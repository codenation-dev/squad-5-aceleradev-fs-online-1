import React, { useEffect, useState, Fragment } from 'react'
import './clientes.css';
import * as apiClientes from '../../apis/clientes'
import Table from '../../componentes/Table/Table'
import Ordenar from '../../componentes/Ordernar/Ordenar'
import FormClientes from '../../componentes/FormClientes/FormClientes';

function Clientes() {
  const [length, setLength] = useState(0)
  const [clientes, setClientes] = useState([])
  const [filtro, setFiltro] = useState({})
  let [carregando, setCarregando] = useState(true)
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
    ""
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
    <div className="container">
      <section className="card">
         {/* {!carregando ?   */}
         <Fragment>
         <FormClientes></FormClientes>
         <Ordenar lista={lista} setFiltro={setFiltro}></Ordenar>
         <Table cabecalho={cabecalho} length={length} filtro={filtro} clientes={clientes} setClientes={setClientes} paginacao={paginacao} setLength={setLength} />
         </Fragment>
          {/* : 'carregando...'} */}
       </section>
    </div>
  )
}

export default Clientes;
