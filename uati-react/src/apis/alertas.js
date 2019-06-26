import protocolo from './configuracao'

export function getAlerts() {
    const url = "/alerts?limit=-1&offset=0"
    return protocolo.get(url)   
}

export function getIdAlert(id) {
    const url =`/alerts/${id}`
    return protocolo.get(url)
}
