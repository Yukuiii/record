/**
 * 认证中间件
 * 检查用户是否已登录，未登录则重定向到登录页
 */

export default defineNuxtRouteMiddleware((to, from) => {
  const { isLoggedIn } = useAuth();

  // 如果用户未登录且不是访问认证相关页面，则重定向到登录页
  if (!isLoggedIn.value && !to.path.startsWith("/auth")) {
    return navigateTo("/auth/login");
  }

  // 如果用户已登录且访问认证页面，则重定向到首页
  if (isLoggedIn.value && to.path.startsWith("/auth")) {
    return navigateTo("/");
  }
});
