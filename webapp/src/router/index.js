import VueRouter from 'vue-router'
import Vue from 'vue'
import AllView from "@/views/AllView.vue";

Vue.use(VueRouter)

const routes = [
    {
        path: '/',
        component: AllView
    }
]

const router = new VueRouter({
    mode: 'history',
    routes
})

export default router