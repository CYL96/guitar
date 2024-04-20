import {createRouter, createWebHashHistory, RouteRecordRaw} from 'vue-router'

import {PageHome} from "../mod/home/home.js";
import Home from "../mod/home/home.vue";

import {PageDown} from "../mod/down/down.js";
import Down from "../mod/down/down.vue";


const routes: Readonly<RouteRecordRaw[]> = [
    {
        path: "/",
        redirect: "Home",
    },
    {
        path: PageHome,
        name: 'Home',
        component: Home
    },
    {
        path: PageDown,
        name: 'Down',
        component: Down
    },
]

const router = createRouter({
    history: createWebHashHistory(),
    routes: routes,
})
export default router
