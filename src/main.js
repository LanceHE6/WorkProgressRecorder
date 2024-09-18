import { createApp } from 'vue'
import { router } from './router/index'
import axios from "axios"
import './style.css'
import App from './App.vue'
import naive from 'naive-ui'
import 'vfonts/Lato.css'
import 'vfonts/FiraCode.css'

const app = createApp(App)

app.provide("axios", axios)

app.use(router)
    .use(naive)

app.mount('#app')
