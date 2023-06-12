import Vue from 'vue'
import App from './App.vue'
import router from './router'
import { BootstrapVue, IconsPlugin } from 'bootstrap-vue'

// Import Bootstrap and BootstrapVue CSS files (order is important)
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import http from "@/plugins/http";
// Make BootstrapVue available throughout your project
Vue.use(BootstrapVue)

// Optionally install the BootstrapVue icon components plugin
Vue.use(IconsPlugin)
Vue.config.productionTip = false

Vue.prototype.$baseUrl="http://localhost:8085"
Vue.use(http, {
  baseUrl: Vue.prototype.$baseUrl
})

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')

export const url = "wss://localhost:8085"