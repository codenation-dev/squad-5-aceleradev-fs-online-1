import React from 'react'
import './Label.css'

const Label = (props) => {
    return <label className={props.classe}> {props.children} </label>
}
export default Label