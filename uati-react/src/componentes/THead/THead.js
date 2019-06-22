import React from 'react'

const THead = ({cabecalho}) => (
    <thead>
        <tr>
            {cabecalho.map((value, i) => (
                <th className="column" id={value==='Cliente'? 'name' : ''} key={i}> {value} </th>
            ))}
        </tr>
    </thead>
)

export default THead