import { createApp, createVNode  } from 'vue'
import App from './App.vue'
import router from './router'
import { createPinia } from 'pinia'
import "element-plus/theme-chalk/el-notification.css"
import "element-plus/theme-chalk/el-message.css"
import "element-plus/theme-chalk/el-message-box.css"

// 完整引入element-plus
//import ElementPlus from 'element-plus'
//import 'element-plus/dist/index.css'

// 按需引入element-plus

// 导入icon图标
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import EnvView from "@/views/gmtool/EnvView.vue";

const app = createApp(App)

app.config.globalProperties.$gmProject = ''

app.use(createPinia())
app.use(router)

for (const [key, icon] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, icon)
}

// const Icons = (props: { icon: string }) => {
//     const { icon } = props
//     return createVNode(ElementPlusIconsVue[icon as keyof typeof ElementPlusIconsVue])
// }
// app.component('Icons', Icons)

// Object.keys(ElementPlusIconsVue).forEach((key) => {
//     app.component(key, ElementPlusIconsVue[key as keyof typeof ElementPlusIconsVue])
// })

app.mount('#app')

