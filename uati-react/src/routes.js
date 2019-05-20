import React from 'react';

import { BrowserRouter, Route, Switch, Redirect } from 'react-router-dom';
import Login from './components/login'
import Dashboard from './components/daschboard'

const Routes = () => (
    <BrowserRouter>
        <Switch>
            <Route exact path = "/" component={Login}/>
            <Route exact path = "/dashboard" component = {Dashboard}/>
        </Switch>
    </BrowserRouter>
);

export default Routes;