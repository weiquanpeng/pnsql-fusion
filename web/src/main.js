import './assets/normalize.css'
import { createApp } from 'vue'
import { createPinia } from 'pinia'
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate';
import App from './App.vue'
import router from '@/router'
import Element from 'element-plus';
import 'element-plus/dist/index.css'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import ArcoVue from '@arco-design/web-vue';
import '@arco-design/web-vue/dist/arco.css';
import AnimatedNumber from 'animated-number-vue3'
import ArcoVueIcon from '@arco-design/web-vue/es/icon';

const app = createApp(App)

//注册图标库
for (const [key,component] of Object.entries(ElementPlusIconsVue)){
    app.component(key,component)
}

const pinia = createPinia()
pinia.use(piniaPluginPersistedstate);

app
    .use(pinia)
    .use(router)
    .use(Element)
    .use(ArcoVue)
    .use(AnimatedNumber)
    .use(ArcoVueIcon)
app.mount('#app')
