import React, { useState, useEffect } from 'react'
import * as apiAlerts from '../../apis/alertas'
import * as apiClientes from '../../apis/clientes'
import THead from '../THead/THead'
import TBody from '../TBody/TBody'
import TFoot from '../TFoot/TFoot'
import './Table.css'

function Table(props) {
  const [start, setStart] = useState(0);
  const [edit, setEdit] = useState(false)
  const [end, setEnd] = useState(props.paginacao);
  const { clientes, usuarios, setClientes, setAlertas, paginacao, cabecalho, length, alertas, filtro, setLength } = props;

  useEffect(() => {
    if (filtro !== undefined && (!filtro.type || filtro.type === '') && (!filtro.customer_name || filtro.customer_name === '') && (!filtro.datetime || filtro.datetime === '')) {
      apiAlerts.getAlerts()
        .then(response => {
          setAlertas(response.data.data.slice(start, end))
          setLength(response.data.data.length)
        })
        .catch(error => {
          if (error.response) {
            alert(error.response.data[0].message)
          }
        })
    }
    else {
      apiAlerts.getAlerts()
        .then(response => {
          const result = response.data.data.filter(item => {
            return ((filtro.type && item.type !== undefined ? item.type.toLowerCase().includes(filtro.type.toLowerCase()) : false) ||
                (filtro.customer_name && item.customer_name !== undefined ? item.customer_name.toLowerCase().includes(filtro.customer_name.toLowerCase()) : false) ||
                (filtro.datetime && item.datetime !== undefined ? item.datetime.includes(filtro.datetime) : false)
            )
          }).slice(start, end)

          setAlertas(result)
          setLength(result.length)
        })
        .catch(error => {
          if (error.response) {
            alert(error.response.data[0].message)
          }
        })
    }
  }, [setAlertas, start, end, paginacao, length, filtro, setLength, usuarios])

  useEffect(() => {
    if (filtro !== undefined && (!filtro.name || filtro.name === '') && (!filtro.salary || filtro.salary === '')) {
      const getClientes = () => {
        apiClientes.getCustomers()
          .then(response => {
            setClientes(response.data.data.slice(start, end))
            setLength(response.data.data.length)
          })
          .catch(error => {
            if (error.response) {
              alert(error.response.data[0].message)
            }
          })
      }
      getClientes()
    }
    else {
      apiClientes.getCustomers()
        .then(response => {
          const result = response.data.data.filter(item => {
            return  (filtro.salary && item.salary !== undefined ? item.salary === filtro.salary : false) || (filtro.name && item.name !== undefined ? item.name.toLowerCase().indexOf(filtro.name.toLowerCase()) !== -1 : false)
          }).slice(start, end)
          setClientes(result)
          setLength(result.length)
        })
        .catch(error => {
          if (error.response) {
            alert(error.response.data[0].message)
          }
        })
    }

  }, [setClientes, start, end, paginacao, length, filtro, setLength, edit])

  return (
    <table className="table">
      <THead cabecalho={cabecalho}> </THead>
      <TBody alertas={alertas} clientes={clientes} usuarios={usuarios} setEdit={setEdit} edit={edit}></TBody>
      <TFoot quantidade={Math.ceil(length / paginacao)} paginacao={paginacao} setEnd={setEnd} setStart={setStart}></TFoot>
    </table>
  )
}

export default Table