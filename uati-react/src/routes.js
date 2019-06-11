import React from 'react';

import { BrowserRouter, Route, Switch } from 'react-router-dom';
import Login from './pages/Login/Login'
import Dashboard from './pages/Dashboard/dashboard'
import Conta from './pages/Conta/Conta'

const Routes = () => (
    <BrowserRouter>
        <Switch>
            <Route exact path="/" component={Login} />
            <Route exact path="/login" component={Login} />
            <Route exact path="/dashboard" component={Dashboard} />
            <Route exact path="/conta" component={Conta} />
        </Switch>
    </BrowserRouter>
);

export default Routes;