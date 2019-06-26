import axios from 'axios'

const configuracoes = {
    baseURL: 'http://bancouati.ga/api'
}

const protocolo = axios.create(configuracoes)

export default protocolo