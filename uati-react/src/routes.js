import React from 'react';

import { BrowserRouter, Route, Switch } from 'react-router-dom';
import Login from './pages/Login/Login'
import Dashboard from './pages/Dashboard/dashboard'
import UploadFile from './pages/UploadFile/uploadFile'
import Conta from './pages/Conta/Conta'
import Alertas from './pages/Alertas/Alertas'
import AlertasInfo from './pages/AlertasInfo/AlertasInfo'

const Routes = () => (
    <BrowserRouter>
        <Switch>
            <Route exact path="/" component={Login} />
            <Route exact path="/login" component={Login} />
            <Route exact path="/dashboard" component={Dashboard} />
            {/* <Route exact path="/upload" component={UploadFile} /> */}
            <Route exact path="/conta" component={Conta} />
            <Route exact path="/alertas" component={Alertas} />
            <Route exact path="/:alerta" component={({match}) => (<AlertasInfo id={match.params.alerta} />)} />
         
        </Switch>
    </BrowserRouter>
);

export default Routes;