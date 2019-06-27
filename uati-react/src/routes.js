import React from 'react'

import { BrowserRouter, Route, Switch, Redirect } from 'react-router-dom'
import Login from './pages/Login/Login'
import Dashboard from './pages/Dashboard/dashboard'
import UploadFile from './pages/UploadFile/uploadFile'
import Conta from './pages/Conta/Conta'
import Alertas from './pages/Alertas/Alertas'
import AlertasInfo from './pages/AlertasInfo/AlertasInfo'
import Clientes from './pages/Clientes/Clientes'

export function estAutenticado() {
   return localStorage.getItem('TOKEN') ? true : false
}

class PrivateRoute extends React.Component {
    render() {
        const Componente = this.props.component
        if (estAutenticado()) {
            return <Route render={() => <Componente {...this.props}></Componente>}></Route>
            }
        else {
            return <Redirect to="/login"></Redirect>
        }
    }
}

class PrivateRouteLogin extends React.Component {
    render() {
        const Componente = this.props.component
        if (!estAutenticado()) {
            return <Route render={(props) => <Componente {...this.props} {...props}></Componente>}></Route>
            }
        else {
            return <Redirect to="/"></Redirect>
        }
    }
}

const Routes = () => (
    <BrowserRouter>
        <Switch>
            <PrivateRoute exact path="/" component={Dashboard} />
            <PrivateRouteLogin path="/login" component={Login} />
            <PrivateRoute path="/dashboard" component={Dashboard} />
            <PrivateRoute path="/upload" component={UploadFile} />
            <PrivateRouteLogin path="/conta" component={Conta} />
            <PrivateRoute exact path="/alertas" component={Alertas} />
            <Route exact path="/alertas/:alerta" component={({match}) => (<AlertasInfo id={match.params.alerta} />)} />
            <PrivateRoute path="/clientes" component={Clientes} />
         
        </Switch>
    </BrowserRouter>
);

export default Routes;