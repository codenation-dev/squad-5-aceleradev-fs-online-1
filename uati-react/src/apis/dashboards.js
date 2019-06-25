import protocolo from './configuracao'
let url = "/dashboard/customer"

export function getLastAlerts(offset,limit) {
    url = `${url}?limit=${limit}&offset=${offset}`
    return protocolo.get(url)   
   }