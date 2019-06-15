import React from 'react';

import { BrowserRouter, Route, Switch } from 'react-router-dom';
import Login from './pages/Login/Login'
import Dashboard from './pages/Dashboard/dashboard'
<<<<<<< HEAD
import UploadFile from './pages/UploadFile/uploadFile'
=======
import Conta from './pages/Conta/Conta'
>>>>>>> 215a9d6d2ee1c0606d7802d228d94759ea959e53

const Routes = () => (
    <BrowserRouter>
        <Switch>
            <Route exact path="/" component={Login} />
            <Route exact path="/login" component={Login} />
            <Route exact path="/dashboard" component={Dashboard} />
<<<<<<< HEAD
            <Route exact path="/upload" component={UploadFile} />
=======
            <Route exact path="/conta" component={Conta} />
>>>>>>> 215a9d6d2ee1c0606d7802d228d94759ea959e53
        </Switch>
    </BrowserRouter>
);

export default Routes;