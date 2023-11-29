import { defineConfig, splitVendorChunkPlugin } from 'vite'
import react from '@vitejs/plugin-react'
import { fileURLToPath } from 'node:url'

// https://vitejs.dev/config/
export default defineConfig({
  server: {
    port: 7702,
    host: '0.0.0.0'
  },
  resolve: {
    alias: [
      {
        find: '@',
        replacement: fileURLToPath(new URL('./src', import.meta.url))
      }
    ]
  },
  plugins: [react(), splitVendorChunkPlugin()],
  build: {
    outDir: process.env.BUILD_GO ? '../backend/web' : 'dist'
  }
})
