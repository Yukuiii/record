<template>
  <div class="min-h-screen bg-gradient-to-br from-blue-50 via-white to-sky-50 relative overflow-hidden">
    <!-- 背景装饰 -->
    <div class="absolute inset-0 overflow-hidden">
      <div class="absolute -top-40 -right-32 w-80 h-80 bg-gradient-to-br from-blue-200/30 to-sky-200/30 rounded-full blur-3xl"></div>
      <div class="absolute -bottom-40 -left-32 w-80 h-80 bg-gradient-to-tr from-cyan-200/30 to-blue-200/30 rounded-full blur-3xl"></div>
    </div>

    <!-- 登录表单容器 -->
    <div class="relative z-10 px-6 pt-16">
      <!-- 品牌Logo区域 -->
      <div class="text-center mb-12">
        <div class="inline-flex items-center justify-center w-16 h-16 bg-gradient-to-r from-blue-500 to-sky-600 rounded-2xl mb-6 shadow-lg">
          <Icon name="material-symbols:account-balance-wallet" size="32" class="text-white" />
        </div>
        <h1 class="text-3xl font-bold bg-gradient-to-r from-gray-900 to-gray-700 bg-clip-text text-transparent mb-3">
          欢迎回来
        </h1>
        <p class="text-gray-500 text-lg">登录您的记账账户</p>
      </div>

      <!-- 登录表单卡片 -->
      <div class="bg-white/80 backdrop-blur-xl rounded-3xl p-8 shadow-xl border border-white/20">
        <form @submit.prevent="handleLogin" class="space-y-6">
          <!-- 邮箱/手机号输入 -->
          <div class="space-y-2">
            <label class="block text-sm font-semibold text-gray-700">
              邮箱或手机号
            </label>
            <div class="relative">
              <Icon name="material-symbols:person" size="20" class="absolute left-4 top-1/2 transform -translate-y-1/2 text-gray-400" />
              <input v-model="loginForm.username" type="text" placeholder="请输入邮箱或手机号"
                class="w-full pl-12 pr-4 py-4 bg-gray-50/50 border border-gray-200/50 rounded-2xl focus:outline-none focus:ring-2 focus:ring-blue-500/50 focus:border-blue-500/50 focus:bg-white transition-all duration-200"
                required />
            </div>
          </div>

          <!-- 密码输入 -->
          <div class="space-y-2">
            <label class="block text-sm font-semibold text-gray-700">
              密码
            </label>
            <div class="relative">
              <Icon name="material-symbols:lock" size="20" class="absolute left-4 top-1/2 transform -translate-y-1/2 text-gray-400" />
              <input v-model="loginForm.password" :type="showPassword ? 'text' : 'password'" placeholder="请输入密码"
                class="w-full pl-12 pr-14 py-4 bg-gray-50/50 border border-gray-200/50 rounded-2xl focus:outline-none focus:ring-2 focus:ring-blue-500/50 focus:border-blue-500/50 focus:bg-white transition-all duration-200"
                required />
              <button type="button" @click="togglePassword"
                class="absolute right-4 top-1/2 transform -translate-y-1/2 text-gray-400 hover:text-gray-600 transition-colors">
                <Icon :name="showPassword ? 'material-symbols:visibility-off' : 'material-symbols:visibility'"
                  size="20" />
              </button>
            </div>
          </div>

          <!-- 记住我和忘记密码 -->
          <div class="flex items-center justify-between">
            <label class="flex items-center group cursor-pointer">
              <input v-model="loginForm.remember" type="checkbox"
                class="w-5 h-5 text-blue-600 border-2 border-gray-300 rounded-lg focus:ring-blue-500 focus:ring-2" />
              <span class="ml-3 text-sm text-gray-600 group-hover:text-gray-800 transition-colors">记住我</span>
            </label>
            <button type="button" @click="handleForgotPassword" 
              class="text-sm font-medium text-blue-600 hover:text-blue-700 transition-colors">
              忘记密码？
            </button>
          </div>

          <!-- 登录按钮 -->
          <button type="submit" :disabled="isLoading"
            class="w-full bg-gradient-to-r from-blue-500 to-sky-600 text-white py-4 rounded-2xl font-semibold hover:from-blue-600 hover:to-sky-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 disabled:opacity-50 disabled:cursor-not-allowed transition-all duration-200 shadow-lg hover:shadow-xl transform hover:-translate-y-0.5">
            <span v-if="!isLoading" class="flex items-center justify-center">
              <Icon name="material-symbols:login" size="20" class="mr-2" />
              登录
            </span>
            <span v-else class="flex items-center justify-center">
              <Icon name="material-symbols:progress-activity" size="20" class="animate-spin mr-2" />
              登录中...
            </span>
          </button>
        </form>
      </div>

      <!-- 注册链接 -->
      <div class="text-center mt-8">
        <p class="text-gray-600">
          还没有账户？
          <NuxtLink to="/auth/register" 
            class="font-semibold text-blue-600 hover:text-blue-700 transition-colors ml-1">
            立即注册
          </NuxtLink>
        </p>
      </div>

      <!-- 第三方登录 -->
      <div class="mt-10">
        <div class="relative">
          <div class="absolute inset-0 flex items-center">
            <div class="w-full border-t border-gray-200"></div>
          </div>
          <div class="relative flex justify-center text-sm">
            <span class="px-6 bg-gradient-to-br from-blue-50 via-white to-sky-50 text-gray-500 font-medium">或者使用</span>
          </div>
        </div>

        <div class="mt-8 grid grid-cols-2 gap-4">
          <!-- 微信登录 -->
          <button @click="handleWechatLogin"
            class="flex items-center justify-center px-6 py-4 bg-white/60 backdrop-blur-sm border border-gray-200/50 rounded-2xl hover:bg-white/80 hover:shadow-md transition-all duration-200 group">
            <Icon name="material-symbols:wechat" size="24" class="text-green-500 mr-3 group-hover:scale-110 transition-transform" />
            <span class="text-sm font-medium text-gray-700">微信</span>
          </button>

          <!-- QQ登录 -->
          <button @click="handleQQLogin"
            class="flex items-center justify-center px-6 py-4 bg-white/60 backdrop-blur-sm border border-gray-200/50 rounded-2xl hover:bg-white/80 hover:shadow-md transition-all duration-200 group">
            <Icon name="material-symbols:chat" size="24" class="text-blue-500 mr-3 group-hover:scale-110 transition-transform" />
            <span class="text-sm font-medium text-gray-700">QQ</span>
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
