import protocolo from './configuracao'
let url = "/alerts"

export function getAlerts() {
 return protocolo.get(url)   
}

export function getIdAlert(id) {
    url +=`/${id}`
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