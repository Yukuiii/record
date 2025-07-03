<template>
  <div class="mobile-layout">
    <!-- 主要内容区域 -->
    <main class="main-content">
      <slot />
    </main>
  </div>
</template>

<script setup>
// 移动端布局专用逻辑
</script>

<style scoped>
.mobile-layout {
  /* 移动端专用样式 */
  -webkit-overflow-scrolling: touch;
  overflow-x: hidden;
}

.main-content {
  @apply w-full;
  /* 移除默认的容器限制，让内容占满屏幕 */
}

/* 移动端视口优化 */
@media (max-width: 768px) {
  .mobile-layout {
    min-height: 100vh;
    min-height: 100dvh;
    /* 动态视口高度 */
  }
}

/* 禁用双击缩放 */
* {
  touch-action: manipulation;
}

/* 优化滚动性能 */
.main-content {
  -webkit-overflow-scrolling: touch;
  transform: translateZ(0);
}
</style>
