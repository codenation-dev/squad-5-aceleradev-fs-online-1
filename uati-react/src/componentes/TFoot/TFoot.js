import React from 'react'
import Botao from '../Botao/Botao'
import './TFoot.css'

const TFoot = (props) => {
    let rows = []

    for (let i = 1; i <= props.quantidade; i++) {
        let end = props.paginacao * i
        let start = end - props.paginacao

        rows.push(<td key={i}>
            <Botao
                classe="paginacao"
                click={() => {
                    props.setStart(start)
                    props.setEnd(end)
                }}>{i}</Botao>
        </td>
        )
    }
    return (
        <tfoot className="tFoot">
            <tr className="tFoot">
                {rows.length > 1 ? rows : ''}
            </tr>
        </tfoot>

    )
}
export default TFoot;