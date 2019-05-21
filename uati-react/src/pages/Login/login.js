import React from 'react';
import TextField from '@material-ui/core/TextField';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';


import './login.css'


export default class Login extends React.Component {
    render() {
        return (
            <div className="LoginContainer">
                <div className="LoginBox">
                    <div className="Fields">
                        <center>
                            <h1 className="LoginTitle">Login</h1>
                            <TextField
                                id="standard-with-placeholder"
                                label="Username"
                                placeholder="userName"
                                style={{ width: '70%' }}
                                margin="normal"
                            />
                            <br></br>
                            <TextField
                                id="outlined-password-input"
                                label="Password"
                                type="password"
                                autoComplete="current-password"
                                style={{ width: '70%' }}
                                margin="normal"
                            />
                        </center>
                    </div>
                    <div className="Buttons">
                        <center>
                            <Button variant="contained" color="primary" >Enter</Button>
                            <Button variant="contained" color="primary" style={{ marginLeft: 10 }} >
                                register</Button>
                        </center>
                    </div>

                </div>
            </div>

        );
    }

} 