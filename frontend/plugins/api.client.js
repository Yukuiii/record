/**
 * API 客户端插件
 * 在客户端初始化时配置 API 相关设置
 */

export default defineNuxtPlugin(() => {
  // 可以在这里添加全局 API 配置
  // 例如：拦截器、错误处理等
  
  // 示例：添加全局错误处理
  const handleApiError = (error) => {
    console.error('API 错误:', error)
    
    // 可以在这里添加全局错误提示
    // 例如使用 toast 组件显示错误信息
  }
  
  // 将错误处理函数提供给应用
  return {
    provide: {
      handleApiError
    }
  }
})
