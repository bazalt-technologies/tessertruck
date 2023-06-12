import axios from "axios"

export default {
    install(Vue, options) {
        Vue.prototype.$http = axios.create({
            baseUrl: "http://localhost:8085",
            headers: options.headers || null
        })
    }
}