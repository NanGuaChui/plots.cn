import { createI18n } from 'vue-i18n';
import enUS from './en-US.json';
import zhCN from './zh-CN.json';
import type { App } from 'vue';
import type { Language } from '@/stores/app/helper';
import { useAppStore } from '@/stores/app';

const i18n = createI18n({
  legacy: false, // 强制使用 Composition API 模式
  locale: 'zh-CN',
  fallbackLocale: 'en-US',
  messages: {
    'en-US': enUS,
    'zh-CN': zhCN,
  },
});

// 推荐在组件内通过 useI18n() 获取 t
// 如需全局 t，可使用 i18n.global.t

export function setLocale(locale: Language) {
  i18n.global.locale.value = locale;
}

export function setupI18n(app: App) {
  const appStore = useAppStore();
  setLocale(appStore.language || 'zh-CN');
  app.use(i18n);
}

export default i18n;
