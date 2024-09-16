// src/main.js
import { createApp } from 'vue';
import App from './App.vue';
import router from './router'; // 导入路由配置
import store from './store';
import ElementPlus from 'element-plus';
import 'element-plus/dist/index.css';

const app = createApp(App);
app.use(router); // 使用路由
app.use(store); // 使用状态管理
app.use(ElementPlus);
app.mount('#app');
