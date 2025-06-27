import { createApp } from 'vue';
import { createPinia } from 'pinia';
import App from './App.vue';
import router from './router';
import './assets/main.css';
import '@fortawesome/fontawesome-free/css/all.css';

document.title = import.meta.env.VITE_APP_TITLE;

const pinia = createPinia();

const app = createApp(App);

app.config.devtools = true;

app.use(pinia);
app.use(router);
app.mount('#app');
