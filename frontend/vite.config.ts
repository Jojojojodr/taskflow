import { fileURLToPath, URL } from 'node:url'

import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueDevTools from 'vite-plugin-vue-devtools'

// https://vite.dev/config/
export default defineConfig(({ mode }) => {
    const env = loadEnv(mode, process.cwd(), '')

    const host = env.VITE_HOST || '0.0.0.0'
    const port = Number(env.VITE_PORT || 3000)

    const apiServer = env.VITE_API_SERVER || 'http://localhost'
    const apiPort = env.VITE_API_PORT || '8080'

    return {
        plugins: [
            vue(),
            vueDevTools(),
        ],
        resolve: {
            alias: {
                '@': fileURLToPath(new URL('./src', import.meta.url))
            },
        },
        server: {
            host,
            port,
            proxy: {
                '/api': {
                    target: `${apiServer}:${apiPort}`,
                    changeOrigin: true,
                    secure: false,
                }
            }
        }
    }
})
