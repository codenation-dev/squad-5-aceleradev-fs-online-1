import React from 'react'
import './Td.css'

const Td = ({children, id, classe}) => <td className={classe} id={id}> {children} </td>

export default Td;