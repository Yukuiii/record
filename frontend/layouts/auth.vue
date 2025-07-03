<template>
  <div class="auth-layout">
    <div class="auth-container">
      <!-- Logo 区域 -->
      <div class="auth-header">
        <NuxtLink to="/" class="logo-link">
          <h1 class="logo">记账本</h1>
        </NuxtLink>
      </div>

      <!-- 主要内容 -->
      <div class="auth-content">
        <slot />
      </div>

      <!-- 页脚链接 -->
      <div class="auth-footer">
        <div class="footer-links">
          <NuxtLink to="/privacy" class="footer-link">隐私政策</NuxtLink>
          <span class="separator">·</span>
          <NuxtLink to="/terms" class="footer-link">使用条款</NuxtLink>
          <span class="separator">·</span>
          <NuxtLink to="/help" class="footer-link">帮助</NuxtLink>
        </div>
        <p class="copyright">© 2025 个人记账应用</p>
      </div>
    </div>
  </div>
</template>

<style scoped>
.auth-layout {
  min-height: 100vh;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 1rem;
}

.auth-container {
  width: 100%;
  max-width: 400px;
  background: white;
  border-radius: 0.75rem;
  box-shadow: 0 20px 25px -5px rgba(0, 0, 0, 0.1), 0 10px 10px -5px rgba(0, 0, 0, 0.04);
  padding: 2rem;
}

.auth-header {
  text-align: center;
  margin-bottom: 2rem;
}

.logo-link {
  text-decoration: none;
}

.logo {
  font-size: 2rem;
  font-weight: bold;
  color: #1f2937;
  margin: 0;
}

.auth-content {
  margin-bottom: 2rem;
}

.auth-footer {
  text-align: center;
  border-top: 1px solid #e5e7eb;
  padding-top: 1.5rem;
}

.footer-links {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 0.5rem;
  margin-bottom: 1rem;
}

.footer-link {
  color: #6b7280;
  text-decoration: none;
  font-size: 0.875rem;
  transition: color 0.2s;
}

.footer-link:hover {
  color: #374151;
}

.separator {
  color: #d1d5db;
  font-size: 0.875rem;
}

.copyright {
  color: #9ca3af;
  font-size: 0.75rem;
  margin: 0;
}

@media (max-width: 480px) {
  .auth-container {
    padding: 1.5rem;
  }

  .footer-links {
    flex-direction: column;
    gap: 0.25rem;
  }

  .separator {
    display: none;
  }
}
</style>
