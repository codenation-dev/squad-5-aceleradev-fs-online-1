import React, { useEffect, useState } from 'react'
import './Alertas.css';
import * as apiAlerts from '../../apis/alertas'
import Table from '../../componentes/Table/Table'

function Alertas() {
  const [length, setLength] = useState(0)
  const usuarios = []
  const [alertas, setAlertas] = useState([])
 // const [filtro, setFiltro] = useState([])
  let [carregando, setCarregando] = useState(true)
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
         {/* {!carregando ?  */}
          <Table cabecalho={cabecalho} length={length} alertas={alertas} usuarios={usuarios} setAlertas={setAlertas} paginacao={paginacao}  />
         
          {/* : 'carregando...'} */}
       
    </div>
  )
}

export default Alertas;
