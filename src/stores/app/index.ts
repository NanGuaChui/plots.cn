import { defineStore } from 'pinia';
import { getDefaultSetting } from './helper';
import type { AppState, Language, Theme } from './helper';
import i18n from '@/i18n';

export const useAppStore = defineStore('app-store', {
  persist: {
    omit: ['theme'],
  },
  state: (): AppState => getDefaultSetting(),
  getters: {
    isDark(state): boolean {
      return state.theme === 'dark';
    },
  },
  actions: {
    setTheme(theme: Theme) {
      console.log('🚀 ~ setTheme ~ theme:', theme);
      this.theme = theme;
    },

    setLanguage(language: Language) {
      if (this.language !== language) {
        this.language = language;
        i18n.global.locale.value = language;
      }
    },
  },
});
