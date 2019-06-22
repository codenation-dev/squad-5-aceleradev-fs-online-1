import React from 'react'
import './Td.css'

const Td = ({children, id}) => <td className="column" id={id}> {children} </td>

export default Td;