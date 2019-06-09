import React from 'react';

import { BrowserRouter, Route, Switch } from 'react-router-dom';
import Login from './pages/Login/login'
import Dashboard from './pages/Dashboard/dashboard'
import UploadFile from './pages/UploadFile/uploadFile'

const Routes = () => (
    <BrowserRouter>
        <Switch>
            <Route exact path="/" component={Login} />
            <Route exact path="/dashboard" component={Dashboard} />
            <Route exact path="/upload" component={UploadFile} />
        </Switch>
    </BrowserRouter>
);

export default Routes;