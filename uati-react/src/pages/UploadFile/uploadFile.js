// import React from 'react';
// import CSVReader from 'react-csv-reader'
// import Dropzone from 'react-dropzone';
// import csv from 'csv';

// import './uploadFile.css'
// import Botao from '../../componentes/Botao/Botao';

// export default class uploadFile extends React.Component {
//     onDrop(files) {

//         this.setState({ files });

//         var file = files[0];

//         const reader = new FileReader();
//         reader.onload = () => {
//             csv.parse(reader.result, (err, data) => {

//                 var userList = [];

//                 for (var i = 0; i < data.length; i++) {
//                     const name = data[i][0];
//                     const newUser = { "name": name };
//                     userList.push(newUser);

//                     /*fetch('https://', {
//                       method: 'POST',
//                       headers: {
//                         'Accept': 'application/json',
//                         'Content-Type': 'application/json',
//                       },
//                       body: JSON.stringify(newUser)
//                     })*/
//                 };
//             });
//         };
//         reader.readAsBinaryString(file);
//     }
//     render() {
//         return (
//             <div className="UploadFileContainer">
//                 <div className="UploadFileBox">
//                     <div className="Fields">
//                         <center>
//                             <h1 className="UploadFileTitle">Envio de arquivos</h1>
//                             <CSVReader
//                                 cssClass="csv-reader-input"
//                                 onFileLoaded={this.handleForce}
//                                 onError={this.handleDarkSideForce}
//                                 inputId="Upload"
//                                 inputStyle={{ color: 'red' }}
//                             />
//                             <br />

//                             <div className="Dropzone">
//                                 <Dropzone
//                                     accept=".csv"
//                                     onDropAccepted={this.onDrop.bind(this)}
//                                 >
//                                     <div>
//                                         <h1 className="DropFile"> Try dropping some files here, or click to select files to upload.</h1>
//                                     </div>
//                                 </Dropzone>
//                             </div>

//                             <div className="Upload">
//                                 <Botao> Enviar </Botao>
//                             </div>
//                         </center>
//                     </div>
//                 </div>
//             </div>
//         );
//     }
// }

