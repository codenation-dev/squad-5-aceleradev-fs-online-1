import protocolo from './configuracao'

export function getCustomers() {
    const url = "/customers"
    return protocolo.get(url)   
}

export function postCustomer(payload) {
    const url ="/customer"
    return protocolo.post(url, payload)
}

export function putCustomer(payload) {
    const url =`/customer/${payload.id}`
    return protocolo.put(url, payload)
}

