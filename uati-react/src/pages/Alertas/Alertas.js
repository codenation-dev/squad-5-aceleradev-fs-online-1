import React, { useEffect, useState } from 'react'
import './Alertas.css';
import * as apiAlerts from '../../apis/alertas'
import Table from '../../componentes/Table/Table'

function Alertas() {
  const [length, setLength] = useState(0)
  const usuarios = []
  const [alertas, setAlertas] = useState([])
  const [filtro, setFiltro] = useState([])
//   let [carregando, setCarregando] = useState(true)
  const paginacao = 10
  const cabecalho = [
    "Tipo",
    "Cliente",
    "Data",
    "Detalhes"
  ]
//   Uma tela para listar e detalhar os alertas, listar os envios de emails 
//   e para quem foi enviado, data, hora e outras funcionalidades que o grupo 
//   julgar interessantes;
//   const lista_ordenar = [
//     {
//       name: 'Nome',
//       value: 'name'
//     },
//     {
//       name: 'E-mail',
//       value: 'email'
//     }
//   ]

//   useEffect(() => {
//     apiUsers.getUsers()
//       .then(response => {
//         setLength(response.data.length)
//         setCarregando(false)
//       })
//       .catch(error => {
//         if (error.response) {
//           alert(error.response.data.erro)
//         }
//       })
//   }, [])

useEffect(() => {
    apiAlerts.getAlerts()
      .then(response => {
        setLength(response.data.data.length)
        // setCarregando(false)
      })
      .catch(error => {
        if (error.response) {
          alert(error.response.data.erro)
        }
      })
  }, [])
  return (
    <div className="container">
     
          {/* <Ordenar lista={lista_ordenar} filtrar={setFiltro}></Ordenar> */}
          <Table cabecalho={cabecalho} length={length} alertas={alertas} usuarios={usuarios} setAlertas={setAlertas} paginacao={paginacao} filtro={filtro} />
       
    </div>
  )
}

export default Alertas;
