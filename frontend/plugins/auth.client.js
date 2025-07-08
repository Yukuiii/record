/**
 * 认证插件 - 客户端
 * 在应用启动时初始化认证状态
 */

export default defineNuxtPlugin(async () => {
  const { checkAuth } = useAuth()
  
  // 在客户端启动时检查认证状态
  await checkAuth()
})