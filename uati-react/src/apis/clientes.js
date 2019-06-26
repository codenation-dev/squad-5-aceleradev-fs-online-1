import protocolo from './configuracao'

export function getCustomers() {
    const url = "/customers?limit=-1&offset=0"
    return protocolo.get(url)   
}

export function postCustomer(payload) {
    const url ="/customer"
    return protocolo.post(url, payload)
}

export function putCustomer(id, payload) {
    const url =`/customer/${id}`
    return protocolo.put(url, payload)
}

