import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import Input from '@material-ui/core/Input';
import InputLabel from '@material-ui/core/InputLabel';
import InputAdornment from '@material-ui/core/InputAdornment';
import FormControl from '@material-ui/core/FormControl';
import TextField from '@material-ui/core/TextField';
import Grid from '@material-ui/core/Grid';



import './styles/login.css'


export default class Login extends React.Component {
    render() {
        return (
            <div className="LoginContainer">
                <div className="LoginBox">
                    <center>
                        <h1 className="LoginTitle">Login</h1>
                        <TextField
                            id="standard-with-placeholder"
                            label="Username"
                            placeholder="userName"
                            style = {{width: 250}} 
                            margin="normal"
                        />
                        <br></br>
                        <TextField
                            id="outlined-password-input"
                            label="Password"
                            type="password"
                            autoComplete="current-password"
                            style = {{width: 250}} 
                            margin="normal"

                            
                        />

                    </center>
                </div>
            </div>

        );
    }

} 