import React, { useEffect, useState, Fragment } from 'react'
import './Alertas.css';
import * as apiAlerts from '../../apis/alertas'
import Table from '../../componentes/Table/Table'
import Ordenar from '../../componentes/Ordernar/Ordenar'

function Alertas() {
  const [length, setLength] = useState(0)
  const usuarios = []
  const [alertas, setAlertas] = useState([])
  const [filtro, setFiltro] = useState([])
  let [carregando, setCarregando] = useState(true)
  let [lista, setLista] = useState([
    {
      name: 'Tipo',
      column: 'type'
    },
    {
      name: 'Cliente',
      column: 'customer_name'
    },
    {
      name: 'Data',
      column: 'datetime'
    }
  ])
  const paginacao = 10
  const cabecalho = [
    "Tipo",
    "Cliente",
    "Data",
    "Detalhes"
  ]
 

useEffect(() => {
    apiAlerts.getAlerts()
    .then(response => {
        setLength(response.data.data.length)
        setCarregando(false)
      })
      .catch(error => {
        if (error.response) {
          alert(error.response.data.erro)
        }
      })
  }, [])

  return (    
    <div className="container">
      <section className="card">
         {!carregando ?  
         <Fragment>
         <Ordenar lista={lista} setFiltro={setFiltro}></Ordenar>
         <Table cabecalho={cabecalho} length={length} filtro={filtro} alertas={alertas} usuarios={usuarios} setAlertas={setAlertas} paginacao={paginacao}  />
         </Fragment>
          : 'carregando...'}
       </section>
    </div>
  )
}

export default Alertas;
