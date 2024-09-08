// src/main.js
import { createApp } from 'vue';
import App from './App.vue';
import router from './router'; // 导入路由配置
import store from './store';

const app = createApp(App);
app.use(router); // 使用路由
app.use(store); // 使用状态管理
app.mount('#app');
