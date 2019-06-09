import React from 'react';
import CSVReader from 'react-csv-reader'
import './uploadFile.css'

import Button from '../../components/Button/Button';

export default class uploadFile extends React.Component {

    render() {
        return (
            <div className="UploadFileContainer">
                <div className="UploadFileBox">
                    <div className="Fields">
                        <center>
                            <h1 className="UploadFileTitle">Upload</h1>
                            <CSVReader
                                cssClass="csv-reader-input"
                                label="Selecione CSV "
                                onFileLoaded={this.handleForce}
                                onError={this.handleDarkSideForce}
                                inputId="Upload"
                                inputStyle={{ color: 'red' }}
                            />
                            <div className="Buttons">
                                <Button> Carregar </Button>
                            </div>
                        </center>
                    </div>
                </div>
            </div>
        );
    }
}

