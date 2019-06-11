import React from 'react';
import './button.css';

const Button = (props) => {
    return (
        <button className="button-login">{props.children} </button>        
    )
}

export default Button;