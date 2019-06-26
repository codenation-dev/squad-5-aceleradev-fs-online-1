import React from 'react'
import { NavLink } from 'react-router-dom';

const Navbar = () => (
    <nav className="menu">
        <ul className="nav nav-pills">
            <li role="presentation">
                <NavLink to="/dashboard" activeClassName="selected">Dashboard</NavLink>
            </li>
            <li role="presentation">
                <NavLink to="/upload" activeClassName="selected">Upload de Clientes</NavLink>
            </li>
            <li role="presentation">
                <NavLink to="/alertas" activeClassName="selected">Alertas</NavLink>
            </li>
            <li role="presentation">
                <NavLink to="/clientes" activeClassName="selected">Clientes</NavLink>
            </li>
            <li role="presentation">
                <NavLink to="/clientes" onClick={() => localStorage.removeItem('TOKEN')} activeClassName="selected"> Logout </NavLink>
            </li>
            
        </ul>
    </nav>
)
export default Navbar