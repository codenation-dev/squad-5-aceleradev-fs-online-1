import React, { useState } from 'react'
import Label from '../Label/Label'
import Botao from '../Botao/Botao'
import './Ordenar.css'

const Ordenar = ({lista, setLista, setFiltro}) => {
    const [itens, setItens] = useState(lista.map(item => (
        { ...item, value: '' })
    ))

    const filter = (value, name) => {
         let novosItens = itens.map(item => {
            return item.name === name ? { ...item, value: value } : {...item}
        })
        setItens(novosItens)
    }

    return (
        <div className="card">
            <div className="card-ordenar">
                {/* <Label>Ordenar:
            </Label> */}
            <form onSubmit={(e) => {
                e.preventDefault()
                setFiltro(itens)
                }}>
                {lista.map((item, i) => {
                    return (
                        <Label key={i}>{item.name}
                            <input type="text" onChange={(e) => filter(e.target.value, item.name)}></input>
                        </Label>)
                })}
                <Botao>Filtrar</Botao>
               
                </form>
            </div>
        </div>

    )

}

export default Ordenar;