import { computed } from 'vue';
import { enUS, zhCN } from 'naive-ui';
import { useAppStore } from '@/stores/app';
import { setLocale } from '@/i18n';

export function useLanguage() {
  const appStore = useAppStore();

  const language = computed(() => {
    switch (appStore.language) {
      case 'en-US':
        setLocale('en-US');
        return enUS;
      case 'zh-CN':
        setLocale('zh-CN');
        return zhCN;
      default:
        setLocale('en-US');
        return enUS;
    }
  });

  return { language };
}
