import React from 'react';
import Dropzone from 'react-dropzone';

import './uploadFile.css'
import Botao from '../../componentes/Botao/Botao';
import  * as api from '../../apis/customer';

export default class uploadFile extends React.Component {
    constructor(props) {
        super(props)
        this.state = {}
    }
    onDrop(files) {

        this.setState({file: files[0]})

    }

    async uploadFile(e) {
        e.preventDefault();
        try {
            await api.upload(this.state.file);
            alert("Clientes importados com sucesso!");
        } catch (e) {
            alert(e.response.data[0].message);
        }
    }

    render() {
        return (
            <div className="UploadFileContainer">
                <div className="UploadFileBox">
                    <div className="Fields">
                        <center>
                            <form onSubmit={this.uploadFile.bind(this)}>
                                <h1 className="UploadFileTitle">Envio de arquivos</h1>
                                <br />

                                <div className="Dropzone">
                                    <Dropzone
                                        accept=".csv"
                                        onDrop={this.onDrop}
                                        onDropAccepted={this.onDrop.bind(this)}
                                    >
                                        <div>
                                            <h1 className="DropFile"> {this.state.file ? this.state.file.name : "Solte o arquivo csv aqui ou click para selecionar."}</h1>
                                        </div>
                                    </Dropzone>
                                </div>

                                <div className="Upload">
                                    <Botao> Enviar </Botao>
                                </div>
                            </form>
                        </center>
                    </div>
                </div>
            </div>
        );
    }
}

