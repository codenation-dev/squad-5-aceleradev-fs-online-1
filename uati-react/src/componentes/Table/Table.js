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
  const { usuarios, setAlertas, paginacao, cabecalho, length, alertas, filtro, setLength } = props;

    useEffect(() => {
      if (filtro !== undefined && (!filtro.type || filtro.type==='') && (!filtro.customer_name || filtro.customer_name==='') && (!filtro.datetime || filtro.datetime==='')) {
        apiAlerts.getAlerts()
          .then(response => {
            setAlertas(response.data.data.slice(start, end))
            setLength(response.data.data.length)
          })
              .catch(error => {
                if (error.response) {
                  alert(error.response.data.erro)
                }
              })          
      } 
      else{
        apiAlerts.getAlerts()
          .then(response => {
            const result = response.data.data.filter(item => {
              return (filtro.type ? item.type.toLowerCase().includes(filtro.type.toLowerCase()) : false ||
                filtro.customer_name ? item.customer_name.toLowerCase().includes(filtro.customer_name.toLowerCase()) : false ||
                  filtro.datetime ? item.datetime.includes(filtro.datetime) : false
              )
            }).slice(start, end)
            setAlertas(result)
            setLength(result.length)
          })
              .catch(error => {
                if (error.response) {
                  alert(error.response.data.erro)
                }
              })
          
      }

    }, [setAlertas, start, end, paginacao, length, filtro, setLength])


    return (
      <table className="table">
        <THead cabecalho={cabecalho}> </THead>
        <TBody alertas={alertas} usuarios={usuarios}></TBody>
        {/* {alertas.length > paginacao -1 || usuarios.length > paginacao -1 ? <TFoot quantidade={Math.ceil(length / paginacao)} paginacao={paginacao} setEnd={setEnd} setStart={setStart}></TFoot> : ''} */}
        <TFoot quantidade={Math.ceil(length / paginacao)} paginacao={paginacao} setEnd={setEnd} setStart={setStart}></TFoot>
      </table>
    )
  }

export default Table