import { createApp } from 'vue'
import App from './App.vue'
import router from "./router";

// 应用实例对象
const app = createApp( App )


// 应用路由
app.use(router)

// 挂载
app.mount('#app')

// 暴露
export default app