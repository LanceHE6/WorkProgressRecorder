import { createApp } from 'vue'
import { router } from './router/index'
import axios from "axios"
import './style.css'
import App from './App.vue'
import naive from 'naive-ui'
import 'vfonts/Lato.css'
import 'vfonts/FiraCode.css'

const app = createApp(App)
// axios.defaults.baseURL = 'http://26.107.171.13:8070'
axios.defaults.baseURL = 'http://127.0.0.1:8070'

app.provide("axios", axios)

app.use(router)
    .use(naive)

app.mount('#app')
