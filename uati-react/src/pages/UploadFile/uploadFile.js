import React from 'react';
import CSVReader from 'react-csv-reader'
import Dropzone from 'react-dropzone';
import './uploadFile.css'
import Botao from '../../componentes/Botao/Botao';
  
class uploadFile extends React.Component {  

    render() {
        return (
            <div className="UploadFileContainer">
                <div className="UploadFileBox">
                    <div className="Fields">
                        <center>
                            <h1 className="UploadFileTitle">Envio de arquivos</h1>
                            <CSVReader
                                cssClass="csv-reader-input"
                                onFileLoaded={this.handleForce}
                                onError={this.handleDarkSideForce}
                                inputId="Upload"
                                inputStyle={{ color: 'red' }}
                            />
                            <br />
                            <div className="Dropzone">
                                <Dropzone
                                    accept=".csv"
                                    onDropAccepted={this.onDrop.bind(this)}
                                >
                                    <div>
                                        <h1 className="DropFile"> Try dropping some files here, or click to select files to upload.</h1>
                                    </div>
                                </Dropzone>
                            </div>
                            <div className="Upload">
                                <Botao type="submit" onChange={this.onFileLoaded}> Enviar </Botao>
                            </div>
                        </center>
                    </div>
                </div>
            </div>
        );    
    }
}

export default uploadFile