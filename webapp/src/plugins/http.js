import axios from "axios"

export default {
    install(Vue, options) {
        Vue.prototype.$http = axios.create({
            baseUrl: "http://server:8089",
            headers: options.headers || null
        })
    }
}