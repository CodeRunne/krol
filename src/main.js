import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'

import Toast from 'vue3-toastify'
import { createPinia } from 'pinia'
import 'vue3-toastify/dist/index.css'

const pinia = createPinia()
const app = createApp(App)

app.use(Toast)
app.use(pinia)
app.mount('#app')
