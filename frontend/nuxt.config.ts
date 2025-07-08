import tailwindcss from "@tailwindcss/vite";

// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: "2025-05-15",
  devtools: { enabled: true },

  // 应用配置
  app: {
    head: {
      title: "个人记账应用",
      meta: [
        { charset: "utf-8" },
        { name: "viewport", content: "width=device-width, initial-scale=1" },
        { name: "description", content: "基于 Nuxt.js 的个人记账应用" },
      ],
    },
  },

  vite: {
    plugins: [tailwindcss()],
  },

  // CSS 配置
  css: ["~/assets/css/main.css", "element-plus/dist/index.css"],

  // 运行时配置
  runtimeConfig: {
    // 私有配置（仅在服务端可用）
    apiSecret: process.env.API_SECRET,

    // 公共配置（客户端和服务端都可用）
    public: {
      apiBase: process.env.API_BASE || "http://localhost:8080/api",
    },
  },

  // 开发服务器配置
  devServer: {
    port: 3000,
  },

  // TypeScript 配置（禁用）
  typescript: {
    typeCheck: false,
  },

  modules: ["@nuxt/icon", "@element-plus/nuxt"],
});
