import protocolo from './configuracao'

export function getAlerts() {
    const url = "/alerts"
    return protocolo.get(url)   
}

export function getIdAlert(id) {
    const url =`/alerts/${id}`
    return protocolo.get(url)
}

// export function paginator(start, end){
//     const url = `?_start=${start}&_end=${end}`
//     return protocolo.get(url)
// }

// export function ordenar(filtro, start, limit){
//     const url = `?_sort=${filtro}&_order=asc&_start=${start}&_limit=${limit}`
//     return protocolo.get(url)
// }