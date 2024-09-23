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

const ENV_SERVER_URL = import.meta.env.SERVER_URL;
const SERVER_URL = 'https://8.137.87.234:9001'
axios.defaults.baseURL = ENV_SERVER_URL ? ENV_SERVER_URL : SERVER_URL

app.provide("axios", axios)

app.use(router)
    .use(naive)

app.mount('#app')
