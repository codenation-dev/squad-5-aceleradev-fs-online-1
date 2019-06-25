import axios from 'axios'

const localStorageWrapper = {
  set: (key, value) => localStorage.setItem(key, JSON.stringify(value)),
  get: (key) => {
    const value = localStorage.getItem(key)

    return value ? JSON.parse(value) : value
  }
}

const NS_TOKEN = 'JWT_TOKEN'

export function get() {
    return localStorageWrapper.get(NS_TOKEN)
}

export function set(token) {
  return localStorageWrapper.set(NS_TOKEN,token)
}