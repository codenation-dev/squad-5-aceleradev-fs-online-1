import React from 'react';

import { BrowserRouter, Route, Switch, Redirect } from 'react-router-dom';
import Login from './pages/Login/login'
import Dashboard from './pages/Dashboard/dashboard'

const Routes = () => (
    <BrowserRouter>
        <Switch>
            <Route exact path = "/" component = {Login}/>
            <Route exact path = "/dashboard" component = {Dashboard}/>
        </Switch>
    </BrowserRouter>
);

export default Routes;