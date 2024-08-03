import { createI18n } from 'vue-i18n';
import en from './lang/en.json';
import hu from './lang/hu.json';

const i18n = createI18n({
  legacy: false,
  locale: localStorage.getItem('lang') || 'hu',
  fallbackLocale: 'hu',
  messages: {
    en,
    hu,
  },
});

export default i18n;
