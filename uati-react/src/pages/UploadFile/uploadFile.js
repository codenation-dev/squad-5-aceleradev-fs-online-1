import React from 'react';
import Dropzone from 'react-dropzone';
import './uploadFile.css'
import Botao from '../../componentes/Botao/Botao';
import api from '../../services/api'
  
class uploadFile extends React.Component {  
    constructor(props) {
        super(props)
        this.state = {}
    }

     upload(file) {
        let fd = new FormData();
        fd.append('file',file);
    
        return api.post("customers",fd,{ headers: { 'Content-Type': 'multipart/form-data' } });
    }

    onDrop(files) {

        this.setState({file: files[0]})

    }

    async uploadFile(e) {
        e.preventDefault();
        try {
            await this.upload(this.state.file);
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
                                    onDropAccepted={this.onDrop.bind(this)}
                                >
                                    <div>
                                        <h1 className="DropFile">{this.state.file ? this.state.file.name : "Solte o arquivo csv aqui ou click para selecionar."}</h1>
                                    </div>
                                </Dropzone>
                            </div>
                            <div className="Upload">
                                <Botao type="submit"> Enviar </Botao>
                            </div>
                            </form>
                        </center>
                    </div>
                </div>
            </div>
        );    
    }
}

export default uploadFile