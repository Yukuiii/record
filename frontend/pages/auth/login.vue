<template>
  <div class="min-h-screen bg-gray-100">
    <!-- 登录表单容器 -->
    <div class="px-4 pt-8">
      <!-- 欢迎标题 -->
      <div class="text-center mb-8">
        <h1 class="text-2xl font-bold text-gray-900 mb-2">欢迎回来</h1>
        <p class="text-gray-600">登录您的记账账户</p>
      </div>

      <!-- 登录表单卡片 -->
      <div class="bg-white rounded-2xl p-6 shadow-sm">
        <form @submit.prevent="handleLogin">
          <!-- 邮箱/手机号输入 -->
          <div class="mb-4">
            <label class="block text-sm font-medium text-gray-700 mb-2">
              邮箱或手机号
            </label>
            <input v-model="loginForm.username" type="text" placeholder="请输入邮箱或手机号"
              class="w-full px-4 py-3 border border-gray-200 rounded-xl focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              required />
          </div>

          <!-- 密码输入 -->
          <div class="mb-6">
            <label class="block text-sm font-medium text-gray-700 mb-2">
              密码
            </label>
            <div class="relative">
              <input v-model="loginForm.password" :type="showPassword ? 'text' : 'password'" placeholder="请输入密码"
                class="w-full px-4 py-3 border border-gray-200 rounded-xl focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent pr-12"
                required />
              <button type="button" @click="togglePassword"
                class="absolute right-3 top-1/2 transform -translate-y-1/2 text-gray-400 hover:text-gray-600">
                <Icon :name="showPassword ? 'material-symbols:visibility-off' : 'material-symbols:visibility'"
                  size="20" />
              </button>
            </div>
          </div>

          <!-- 记住我和忘记密码 -->
          <div class="flex items-center justify-between mb-6">
            <label class="flex items-center">
              <input v-model="loginForm.remember" type="checkbox"
                class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500" />
              <span class="ml-2 text-sm text-gray-600">记住我</span>
            </label>
            <button type="button" @click="handleForgotPassword" class="text-sm text-blue-600 hover:text-blue-700">
              忘记密码？
            </button>
          </div>

          <!-- 登录按钮 -->
          <button type="submit" :disabled="isLoading"
            class="w-full bg-blue-500 text-white py-3 rounded-xl font-medium hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 disabled:opacity-50 disabled:cursor-not-allowed transition-colors">
            <span v-if="!isLoading">登录</span>
            <span v-else class="flex items-center justify-center">
              <Icon name="material-symbols:progress-activity" size="20" class="animate-spin mr-2" />
              登录中...
            </span>
          </button>
        </form>
      </div>

      <!-- 注册链接 -->
      <div class="text-center mt-6">
        <p class="text-gray-600">
          还没有账户？
          <NuxtLink to="/auth/register" class="text-blue-600 hover:text-blue-700 font-medium">
            立即注册
          </NuxtLink>
        </p>
      </div>

      <!-- 第三方登录 -->
      <div class="mt-8">
        <div class="relative">
          <div class="absolute inset-0 flex items-center">
            <div class="w-full border-t border-gray-200"></div>
          </div>
          <div class="relative flex justify-center text-sm">
            <span class="px-4 bg-gray-100 text-gray-500">或者使用</span>
          </div>
        </div>

        <div class="mt-6 grid grid-cols-2 gap-3">
          <!-- 微信登录 -->
          <button @click="handleWechatLogin"
            class="flex items-center justify-center px-4 py-3 border border-gray-200 rounded-xl hover:bg-gray-100 transition-colors">
            <Icon name="material-symbols:wechat" size="20" class="text-green-500 mr-2" />
            <span class="text-sm text-gray-700">微信</span>
          </button>

          <!-- QQ登录 -->
          <button @click="handleQQLogin"
            class="flex items-center justify-center px-4 py-3 border border-gray-200 rounded-xl hover:bg-gray-50 transition-colors">
            <Icon name="material-symbols:chat" size="20" class="text-blue-500 mr-2" />
            <span class="text-sm text-gray-700">QQ</span>
          </button>
        </div>
      </div>
    </div>

    <!-- 底部安全区域 -->
    <div class="h-20"></div>
  </div>
</template>

<script setup>
// 使用移动端布局
definePageMeta({
  layout: 'mobile'
})

// 设置页面元信息
useHead({
  title: '登录 - 个人记账应用',
  meta: [
    { name: 'description', content: '登录您的个人记账账户' },
    { name: 'viewport', content: 'width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no' }
  ]
})

// 登录表单数据
const loginForm = ref({
  username: '',
  password: '',
  remember: false
})

// 状态管理
const isLoading = ref(false)
const showPassword = ref(false)

// 切换密码显示
const togglePassword = () => {
  showPassword.value = !showPassword.value
}

// 处理登录
const handleLogin = async () => {
  if (!loginForm.value.username || !loginForm.value.password) {
    // 这里可以添加表单验证提示
    console.log('请填写完整的登录信息')
    return
  }

  isLoading.value = true

  try {
    // 这里添加实际的登录逻辑
    console.log('登录信息:', loginForm.value)

    // 模拟登录请求
    await new Promise(resolve => setTimeout(resolve, 2000))

    // 登录成功后跳转到首页
    await navigateTo('/')

  } catch (error) {
    console.error('登录失败:', error)
    // 这里可以添加错误提示
  } finally {
    isLoading.value = false
  }
}

// 处理忘记密码
const handleForgotPassword = () => {
  console.log('忘记密码')
  // 这里可以跳转到忘记密码页面或显示重置密码弹窗
}

// 处理微信登录
const handleWechatLogin = () => {
  console.log('微信登录')
  // 这里添加微信登录逻辑
}

// 处理QQ登录
const handleQQLogin = () => {
  console.log('QQ登录')
  // 这里添加QQ登录逻辑
}
</script>

<style scoped>
/* 登录页面样式 */
/* 如果需要特殊样式可以在这里添加 */
</style>
