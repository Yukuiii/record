/**
 * 认证中间件
 * 检查用户是否已登录，未登录则重定向到登录页
 */

export default defineNuxtRouteMiddleware(async (to, from) => {
  const { isLoggedIn, checkAuth } = useAuth();

  // 先检查认证状态，确保状态是最新的
  await checkAuth();

  // 如果用户未登录且不是访问认证相关页面，则重定向到登录页
  if (!isLoggedIn.value && !to.path.startsWith("/auth")) {
    // 根据访问的页面生成不同的提示消息
    let message = "请先登录后再访问该页面";
    let messageType = "default";

    if (to.path.includes("/profile")) {
      message = "请先登录后再查看个人资料";
      messageType = "profile";
    } else if (to.path.includes("/records")) {
      message = "请先登录后再管理您的记账记录";
      messageType = "records";
    } else if (to.path.includes("/settings")) {
      message = "请先登录后再访问设置页面";
      messageType = "settings";
    }

    // 保存原始访问路径，登录后可以跳转回来
    const redirectPath = to.fullPath;

    return navigateTo(
      `/auth/login?redirect=${encodeURIComponent(
        redirectPath
      )}&message=${encodeURIComponent(message)}&type=${messageType}`
    );
  }

  // 如果用户已登录且访问认证页面，则重定向到首页
  if (isLoggedIn.value && to.path.startsWith("/auth")) {
    return navigateTo("/");
  }
});
