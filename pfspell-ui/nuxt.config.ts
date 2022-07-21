import { defineNuxtConfig } from 'nuxt'

// https://v3.nuxtjs.org/api/configuration/nuxt.config
export default defineNuxtConfig({
    ssr: false,
    runtimeConfig: {
        public: {
            API_BASE_URL: process.env.NODE_ENV !== 'production' ? process.env.API_DEV_URL: '/api/'
        }
    },
    typescript: {
        shim: false
    },
    experimental: {
        reactivityTransform: true
    },
    modules: [
        './modules/fonts',
        '@nuxtjs/tailwindcss',
        '@nuxtjs/color-mode'
    ],
    googleFonts: {
        families: {
            Inter: true,
            Balthazar: true
        }
    },
    colorMode: {
        preference: 'light',
        dataValue: 'theme'
    }
})
