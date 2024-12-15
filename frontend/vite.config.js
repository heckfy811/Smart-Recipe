import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';
import path from 'path';

export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      '@': '/src',
    },
  },
  build: {
    outDir: 'dist', // Указываем папку для сборки
    rollupOptions: {
      input: path.resolve(__dirname, 'index.html'), // Указываем путь к index.html
    },
  },
  server: {
    watch: {
      usePolling: true, // Включение polling (необязательно, зависит от окружения)
    },
    historyApiFallback: true, // Перенаправление для маршрутов Vue Router
  },
  base: '/', // Базовый URL (если проект деплоится в корень)
});
