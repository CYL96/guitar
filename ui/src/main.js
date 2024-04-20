import ElementPlus from 'element-plus'
import { createApp } from 'vue'
import 'element-plus/dist/index.css'
import App from './App.vue'
import router from "@/components/router/router.ts";
import {zhCn} from "element-plus/es/locale/index";
import {runConfig} from "@/components/common/config.ts";

export  let baseUrl = "http://10.5.10.87:55002"
if (import.meta.env.PROD){
    runConfig.server =  window.location.origin
}else {
    runConfig.server = baseUrl
}

const app = createApp(App)

app.use(ElementPlus, {
    locale: zhCn,
})
app.use(router)

app.mount('#app')


