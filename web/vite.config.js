import { fileURLToPath, URL } from 'node:url'
import {defineConfig, loadEnv} from 'vite'
import vue from '@vitejs/plugin-vue'
import axios from "axios";
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import { ElementPlusResolver } from 'unplugin-vue-components/resolvers'

export default defineConfig(({ command, mode }) => {
    const env = loadEnv(mode, process.cwd(), '')
    return {
        // vite 配置
        plugins: [
            vue(),
            AutoImport({
                imports: ['vue','pinia'],
                resolvers: [ElementPlusResolver()],
            }),
            Components({
                resolvers: [ElementPlusResolver()],
            })
        ],
        server:{
            // 服务启动网页自动打开，如果使用docker-compose开发模式，设置为false
            open: true,
            port: env.VITE_CLI_PORT,
            host: '0.0.0.0'
            // proxy: {
            //     [env.VITE_BASE_API]: { // 需要代理的路径   例如 '/api'
            //         target: `${env.VITE_BASE_PATH}/`, // 代理到 目标路径
            //         changeOrigin: true,
            //         //rewrite: path => path.replace(new RegExp('^' + env.VITE_BASE_API), ''),
            //     }
            // },
        },
        resolve: {
            alias: {
                '@': fileURLToPath(new URL('./src', import.meta.url))
            }
        }
    }
})
