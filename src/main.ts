import './assets/css/tailwindcss.css';
import './assets/css/main.scss';

import { createApp } from 'vue';
import { setupI18n } from './i18n';
import { setupStore } from './stores';
import { setupRouter } from './router';

import App from './App.vue';

async function bootstrap() {
  const app = createApp(App);

  setupStore(app);

  setupI18n(app);

  await setupRouter(app);

  app.mount('#app');
}

bootstrap();
