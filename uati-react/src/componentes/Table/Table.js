import React, { useState, useEffect } from 'react'
//import * as apiUsers from '../../apis/users'
import * as apiAlerts from '../../apis/alertas'
import THead from '../THead/THead'
import TBody from '../TBody/TBody'
import TFoot from '../TFoot/TFoot'
import './Table.css'

function Table(props) {
  const [start, setStart] = useState(0);
  const [end, setEnd] = useState(props.paginacao);
  const { usuarios, setAlertas, paginacao, cabecalho, length, alertas } = props;


  useEffect(() => {
    apiAlerts.getAlerts()
      .then(response => {
        setAlertas(response.data.data.slice(start, end))
        //  setCarregando(false)
      })
      .catch(error => {
        if (error.response) {
          alert(error.response.data.erro)
        }
      })

  }, [setAlertas, start, end, paginacao, length])

  //   useEffect(() => {
  //     if (filtro.length > 0) {
  //         const limit = () => apiUsers.ordenar(filtro.join(), start, paginacao)
  //             .then(res => {
  //                 setUsuarios(res.data)
  //             })
  //             .catch(error => {
  //                 if (error.response) {
  //                     alert(error.response.data.erro)
  //                 }
  //             })

  //         limit()
  //     }
  //     else {
  //         const limit = (start, end) => (apiUsers.paginator(start, end)
  //             .then(res => {
  //                 setUsuarios(res.data)
  //             })
  //             .catch(error => {
  //                 if (error.response) {
  //                     alert(error.response.data.erro)
  //                 }
  //             })
  //         )
  //         limit(start, end)

  //     }

  // }, [start, end, paginacao, setUsuarios, filtro])


  return (
    <table className="card table">
      <THead cabecalho={cabecalho}> </THead>
      <TBody alertas={alertas} usuarios={usuarios}></TBody>
      <TFoot quantidade={Math.ceil(length / paginacao)} paginacao={paginacao} setEnd={setEnd} setStart={setStart}></TFoot>
    </table>
  )
}

export default Table