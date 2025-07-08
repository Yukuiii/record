<template>
  <div class="min-h-screen bg-gray-50">
    <!-- 头部导航组件 -->
    <LayoutHeaderNav />

    <!-- 注册表单容器 -->
    <div class="px-4 pt-8">
      <!-- 欢迎标题 -->
      <div class="text-center mb-8">
        <h1 class="text-2xl font-bold text-gray-900 mb-2">创建账户</h1>
        <p class="text-gray-600">开始您的记账之旅</p>
      </div>

      <!-- 注册表单卡片 -->
      <div class="bg-white rounded-2xl p-6 shadow-sm">
        <form @submit.prevent="handleRegister">
          <!-- 邮箱输入 -->
          <div class="mb-4">
            <label class="block text-sm font-medium text-gray-700 mb-2">
              邮箱地址
            </label>
            <input
              v-model="registerForm.email"
              type="email"
              placeholder="请输入邮箱地址"
              class="w-full px-4 py-3 border border-gray-200 rounded-xl focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              required
            />
          </div>

          <!-- 手机号输入 -->
          <div class="mb-4">
            <label class="block text-sm font-medium text-gray-700 mb-2">
              手机号码
            </label>
            <input
              v-model="registerForm.phone"
              type="tel"
              placeholder="请输入手机号码"
              class="w-full px-4 py-3 border border-gray-200 rounded-xl focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              required
            />
          </div>

          <!-- 密码输入 -->
          <div class="mb-4">
            <label class="block text-sm font-medium text-gray-700 mb-2">
              设置密码
            </label>
            <div class="relative">
              <input
                v-model="registerForm.password"
                :type="showPassword ? 'text' : 'password'"
                placeholder="请设置密码（至少6位）"
                class="w-full px-4 py-3 border border-gray-200 rounded-xl focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent pr-12"
                required
                minlength="6"
              />
              <button
                type="button"
                @click="togglePassword"
                class="absolute right-3 top-1/2 transform -translate-y-1/2 text-gray-400 hover:text-gray-600"
              >
                <Icon 
                  :name="showPassword ? 'material-symbols:visibility-off' : 'material-symbols:visibility'" 
                  size="20" 
                />
              </button>
            </div>
          </div>

          <!-- 确认密码输入 -->
          <div class="mb-6">
            <label class="block text-sm font-medium text-gray-700 mb-2">
              确认密码
            </label>
            <div class="relative">
              <input
                v-model="registerForm.confirmPassword"
                :type="showConfirmPassword ? 'text' : 'password'"
                placeholder="请再次输入密码"
                class="w-full px-4 py-3 border border-gray-200 rounded-xl focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent pr-12"
                required
              />
              <button
                type="button"
                @click="toggleConfirmPassword"
                class="absolute right-3 top-1/2 transform -translate-y-1/2 text-gray-400 hover:text-gray-600"
              >
                <Icon 
                  :name="showConfirmPassword ? 'material-symbols:visibility-off' : 'material-symbols:visibility'" 
                  size="20" 
                />
              </button>
            </div>
          </div>

          <!-- 用户协议 -->
          <div class="mb-6">
            <label class="flex items-start">
              <input
                v-model="registerForm.agreeTerms"
                type="checkbox"
                class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500 mt-0.5"
                required
              />
              <span class="ml-2 text-sm text-gray-600">
                我已阅读并同意
                <button type="button" class="text-blue-600 hover:text-blue-700">用户协议</button>
                和
                <button type="button" class="text-blue-600 hover:text-blue-700">隐私政策</button>
              </span>
            </label>
          </div>

          <!-- 注册按钮 -->
          <button
            type="submit"
            :disabled="isLoading"
            class="w-full bg-blue-500 text-white py-3 rounded-xl font-medium hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 disabled:opacity-50 disabled:cursor-not-allowed transition-colors"
          >
            <span v-if="!isLoading">创建账户</span>
            <span v-else class="flex items-center justify-center">
              <Icon name="material-symbols:progress-activity" size="20" class="animate-spin mr-2" />
              注册中...
            </span>
          </button>
        </form>
      </div>

      <!-- 登录链接 -->
      <div class="text-center mt-6">
        <p class="text-gray-600">
          已有账户？
          <NuxtLink to="/auth/login" class="text-blue-600 hover:text-blue-700 font-medium">
            立即登录
          </NuxtLink>
        </p>
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
  title: '注册 - 个人记账应用',
  meta: [
    { name: 'description', content: '创建您的个人记账账户' },
    { name: 'viewport', content: 'width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no' }
  ]
})

// 使用认证组合函数
const { register, isLoading } = useAuth()

// 注册表单数据
const registerForm = ref({
  email: '',
  phone: '',
  password: '',
  confirmPassword: '',
  agreeTerms: false
})

// 状态管理
const showPassword = ref(false)
const showConfirmPassword = ref(false)

// 切换密码显示
const togglePassword = () => {
  showPassword.value = !showPassword.value
}

const toggleConfirmPassword = () => {
  showConfirmPassword.value = !showConfirmPassword.value
}

// 处理注册
const handleRegister = async () => {
  // 表单验证
  if (!registerForm.value.email || !registerForm.value.phone ||
      !registerForm.value.password || !registerForm.value.confirmPassword) {
    ElMessage.warning('请填写完整的注册信息')
    return
  }

  if (registerForm.value.password !== registerForm.value.confirmPassword) {
    ElMessage.warning('两次输入的密码不一致')
    return
  }

  if (!registerForm.value.agreeTerms) {
    ElMessage.warning('请同意用户协议和隐私政策')
    return
  }

  try {
    await register({
      email: registerForm.value.email,
      phone: registerForm.value.phone,
      password: registerForm.value.password
    })
    // 注册成功后会自动跳转到首页，无需额外处理
  } catch (error) {
    console.error('注册失败:', error)
    // 错误会通过useApi自动显示toast，无需手动处理
  }
}
</script>

<style scoped>
/* 注册页面样式 */
/* 如果需要特殊样式可以在这里添加 */
</style>
