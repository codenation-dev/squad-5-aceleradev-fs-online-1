import React from 'react';

import { BrowserRouter, Route, Switch, NavLink } from 'react-router-dom';
import Login from './pages/Login/Login'
import Dashboard from './pages/Dashboard/dashboard'
//import UploadFile from './pages/UploadFile/uploadFile'
import Conta from './pages/Conta/Conta'
import Alertas from './pages/Alertas/Alertas'
import AlertasInfo from './pages/AlertasInfo/AlertasInfo'
import Clientes from './pages/Clientes/Clientes'

const Routes = () => (
    <BrowserRouter>
        <div className="menu">
            <ul className="nav nav-pills">
                <li role="presentation">
                    <NavLink to="/login" activeClassName="selected">Login</NavLink>
                </li>
                <li role="presentation">
                    <NavLink to="/dashboard" activeClassName="selected">Dashboard</NavLink>
                </li>
                <li role="presentation">
                    <NavLink to="/upload" activeClassName="selected">Upload de Clientes</NavLink>
                </li>
                <li role="presentation">
                    <NavLink to="/alertas" activeClassName="selected">Alertas</NavLink>
                </li>
            </ul>
        </div>
        <Switch>
            <Route exact path="/" component={Login} />
            <Route exact path="/login" component={Login} />
            <Route exact path="/dashboard" component={Dashboard} />
            {/* <Route exact path="/upload" component={UploadFile} /> */}
            <Route exact path="/conta" component={Conta} />
            <Route exact path="/alertas" component={Alertas} />
            <Route exact path="/alertas/:alerta" component={({match}) => (<AlertasInfo id={match.params.alerta} />)} />
            <Route exact path="/clientes" component={Clientes} />
         
        </Switch>
    </BrowserRouter>
);

export default Routes;