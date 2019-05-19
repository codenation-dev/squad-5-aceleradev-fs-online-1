import React from 'react';

import { BrowserRouter, Route, Switch, Redirect } from 'react-router-dom';
import Login from './components/login'
import Dashboard from './components/daschboard'

const Routes = () => (
    <BrowserRouter>
        <Switch>
            <Route exact path = "/" component={Login}/>
        </Switch>
    </BrowserRouter>
);

export default Routes;