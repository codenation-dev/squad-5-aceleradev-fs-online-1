import protocolo from './configuracao'
const url = "/customers"

export function upload(file) {
    const config = { headers: { 'Content-Type': 'multipart/form-data' } };
    let fd = new FormData();
    fd.append('file',file);

    return protocolo.post(url,fd,config);
}