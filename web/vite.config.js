import { defineConfig, loadEnv } from 'vite'; // 导入 loadEnv
import vue from '@vitejs/plugin-vue';
import AutoImport from 'unplugin-auto-import/vite';
import Components from 'unplugin-vue-components/vite';
import { ElementPlusResolver } from 'unplugin-vue-components/resolvers';
import { fileURLToPath, URL } from 'url';

export default defineConfig(({ command, mode }) => {
    const env = loadEnv(mode, process.cwd(), ''); // 加载环境变量

    return {
        // 插件配置
        plugins: [
            vue(),
            AutoImport({
                imports: ['vue', 'pinia'],
                resolvers: [ElementPlusResolver()],
            }),
            Components({
                resolvers: [ElementPlusResolver()],
            }),
        ],

        // 开发服务器配置
        server: {
            open: true, // 服务启动时自动打开浏览器
            port: env.VITE_CLI_PORT, // 前端服务端口
            host: '0.0.0.0', // 监听所有网络接口
            // proxy: {
            //     // 代理配置
            //     [env.VITE_BASE_API]: {
            //         target: env.VITE_BASE_PATH, // 后端服务地址
            //         changeOrigin: true, // 允许跨域
            //         rewrite: (path) => path.replace(new RegExp(`^${env.VITE_BASE_API}`), ''), // 重写路径
            //     },
            // },
        },

        // 路径别名配置
        resolve: {
            alias: {
                '@': fileURLToPath(new URL('./src', import.meta.url)),
            },
        },
    };
});