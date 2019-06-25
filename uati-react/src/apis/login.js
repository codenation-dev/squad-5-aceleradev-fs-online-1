import protocolo from './configuracao'
const url = "/auth"

export function login(credential) {
    return protocolo.post(url,credential)
}