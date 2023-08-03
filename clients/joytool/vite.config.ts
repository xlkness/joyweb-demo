import { fileURLToPath, URL } from 'node:url'
import path from 'path'
import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import { ElementPlusResolver } from 'unplugin-vue-components/resolvers'
import Icons from 'unplugin-icons/vite'
import IconsResolver from 'unplugin-icons/resolver'
// import Inspect from 'vite-plugin-inspect'

const pathSrc = path.resolve(__dirname, 'src')

// https://vitejs.dev/config/
export default defineConfig({
  css: {
    preprocessorOptions: {
      scss: {
        additionalData: "@use './src/style/element/index.scss' as *;",
      }
    }
  },
  plugins: [
      vue(),
      AutoImport({
        resolvers: [
            ElementPlusResolver(),
            IconsResolver({
              prefix: 'Icon',
            }),
        ],
        dts: path.resolve(pathSrc, 'auto-imports.d.ts'),
      }),
      Components({
        resolvers: [
            ElementPlusResolver(),
            IconsResolver({
              enabledCollections: ['ep'],
            }),
        ],
        dts: path.resolve(pathSrc, 'components.d.ts'),
      }),

    Icons({
      autoInstall: true,
    }),

    // Inspect(),
  ],
  define: {
    'process.env': {}
  },
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },
  server: {
    open: false,  //启动项目后打开浏览器
    port: 8080,   //端口
    host: '192.168.1.60',
    https: false,
    lintOnSave: false,
    proxy: {
        '/user': {
            target: 'http://192.168.1.60:9000/',  //API服务地址
            changeOrigin: true,             //开启跨域
            rewrite: (path) => path.replace(/^\/user/, '')
        },
        '/gmtoolapi': {
            target: 'http://192.168.1.60:9001/',  //API服务地址
            changeOrigin: true,             //开启跨域
            rewrite: (path) => path.replace(/^\/gmtoolapi/, '')
        },
    }
  }
})
