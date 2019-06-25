import React from 'react'
import './Botao.css'

function Botao(props) {
  let classes = "botao"
  if (props.desabilitado) {
    classes += " botao--desabilitado"
  }

  if (props.classe==="paginacao") {
    classes += "-paginacao"
  }
  
  return (
    <button className={classes} disabled={props.desabilitado} onClick={props.click}>
      {props.children}
    </button>
  )
}

export default Botao