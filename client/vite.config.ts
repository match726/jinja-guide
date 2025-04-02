import path from "path";
import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react-swc'
import env from 'vite-plugin-env-compatible';

// https://vite.dev/config/
export default defineConfig({
  plugins: [react(), env({ prefix: "VITE", mountedPath: "process.env" })],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./src")
    }
  }
});
