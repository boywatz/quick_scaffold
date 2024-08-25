// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },
  ssr: true,
  modules: ["@nuxt/ui"],
  compatibilityDate: "2024-08-24",
  app: {
    baseURL: '/second-frontend/'
  },
  runtimeConfig: {
    public: {
      API_BASE_URL: process.env.NUXT_PUBLIC_API_BASE_URL || "http://localhost/api"
    }
  }
});
