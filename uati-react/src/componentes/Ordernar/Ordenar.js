import React, { useState } from 'react'
import Label from '../Label/Label'
import Botao from '../Botao/Botao'
import './Ordenar.css'

const Ordenar = ({ lista, setLista, setFiltro }) => {
    const [itens, setItens] = useState({})

    const filter = (value, name) => {
        setItens({ ...itens, [name]: value })
    }

    return (
        <div className="card">
            <div className="card-ordenar">
                {/* <Label>Ordenar:
            </Label> */}
                <form className="card-ordenar"
                    onSubmit={(e) => {
                        e.preventDefault()
                        setFiltro(itens)
                    }}>
                    {lista.map((item, i) => {
                        return (
                            <Label key={i}>{item.name}
                                <input type="text" onChange={(e) => filter(e.target.value, item.column)}></input>
                            </Label>)
                    })}
                    <Botao classe={'paginacao'}>Filtrar</Botao>

                </form>
            </div>
        </div>

    )

}

export default Ordenar;