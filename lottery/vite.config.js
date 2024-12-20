import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";

// https://vitejs.dev/config/
export default defineConfig({
  base: "./",
  plugins: [vue()],
  resolve: {
    alias: {
      src: "/src",
    },
  },
  build: {
    assetsInlineLimit: 4 * 1024,
  },
});
