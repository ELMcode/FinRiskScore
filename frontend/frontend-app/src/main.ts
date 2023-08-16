import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import './assets/base.css'


import { createVuetify } from 'vuetify'
import 'vuetify/styles'

const vuetify = createVuetify()

const app = createApp(App)

app.use(vuetify)
app.use(router)

app.mount('#app')