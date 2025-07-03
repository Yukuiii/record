<template>
  <div class="app-layout">
    <!-- 头部导航 -->
    <LayoutHeader />

    <!-- 主要内容区域 -->
    <main class="main-content">
      <div class="container">
        <slot />
      </div>
    </main>

    <!-- 页脚 -->
    <footer class="footer">
      <div class="container">
        <div class="footer-content">
          <p class="footer-text">
            © 2025 个人记账应用. 保留所有权利.
          </p>
          <div class="footer-links">
            <NuxtLink to="/privacy" class="footer-link">
              隐私政策
            </NuxtLink>
            <NuxtLink to="/terms" class="footer-link">
              使用条款
            </NuxtLink>
            <NuxtLink to="/help" class="footer-link">
              帮助中心
            </NuxtLink>
          </div>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup lang="ts">
// 检查用户认证状态
const { checkAuth } = useAuth()

// 页面加载时检查认证状态
onMounted(() => {
  checkAuth()
})
</script>

<style scoped>
.app-layout {
  @apply min-h-screen flex flex-col;
}

.main-content {
  @apply flex-1 py-6;
}

.footer {
  @apply bg-gray-100 border-t border-gray-200 py-6 mt-auto;
}

.footer-content {
  @apply flex flex-col md:flex-row justify-between items-center space-y-4 md:space-y-0;
}

.footer-text {
  @apply text-gray-600 text-sm;
}

.footer-links {
  @apply flex space-x-6;
}

.footer-link {
  @apply text-gray-500 hover:text-gray-700 text-sm transition-colors;
}
</style>
