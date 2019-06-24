import axios from 'axios';

const api = axios.create({
    baseURL: 'http://bancouati.ga/api/',
});

export default api;