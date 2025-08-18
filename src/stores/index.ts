import type { App } from 'vue';
import { createPinia } from 'pinia';
import { createPersistedState } from 'pinia-plugin-persistedstate';
import { useDebounceFn } from '@vueuse/core';

// 自定义防抖的 Storage 适配器
const debounceStorage = {
  getItem: (key: string) => localStorage.getItem(key),
  setItem: useDebounceFn((key: string, value: string) => {
    localStorage.setItem(key, value);
  }, 1000),
};

export const store = createPinia();
store.use(createPersistedState({ storage: debounceStorage }));

export function setupStore(app: App) {
  // localStorage.removeItem('chat-store');
  localStorage.removeItem('app-store');
  app.use(store);
}
