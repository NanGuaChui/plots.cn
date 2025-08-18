export type Theme = 'light' | 'dark' | 'auto';

export type Language = 'zh-CN' | 'en-US';

export interface AppState {
  theme: Theme;
  language: Language;
}

export function getDefaultSetting(): AppState {
  return { theme: 'dark', language: 'zh-CN' };
}
