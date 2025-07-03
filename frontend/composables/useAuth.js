/**
 * 用户认证相关的组合式函数
 */

export const useAuth = () => {
  const api = useApi()
  const router = useRouter()
  
  // 用户状态
  const user = ref(null)
  const isLoggedIn = computed(() => !!user.value)
  const isLoading = ref(false)

  // 认证 token
  const token = useCookie('auth-token', {
    default: () => null,
    maxAge: 60 * 60 * 24 * 7, // 7 天
    secure: true,
    sameSite: 'strict'
  })

  // 登录
  const login = async (credentials) => {
    isLoading.value = true
    
    try {
      const response = await api.post('/auth/login', credentials)

      if (response.success) {
        user.value = response.data.user
        token.value = response.data.token
        
        // 跳转到首页
        await router.push('/')
        
        return { success: true }
      } else {
        throw new Error(response.message || '登录失败')
      }
    } catch (error) {
      console.error('登录失败:', error)
      return {
        success: false,
        message: error instanceof Error ? error.message : '登录失败'
      }
    } finally {
      isLoading.value = false
    }
  }

  // 注册
  const register = async (data) => {
    isLoading.value = true
    
    try {
      const response = await api.post('/auth/register', data)

      if (response.success) {
        user.value = response.data.user
        token.value = response.data.token
        
        // 跳转到首页
        await router.push('/')
        
        return { success: true }
      } else {
        throw new Error(response.message || '注册失败')
      }
    } catch (error) {
      console.error('注册失败:', error)
      return {
        success: false,
        message: error instanceof Error ? error.message : '注册失败'
      }
    } finally {
      isLoading.value = false
    }
  }

  // 登出
  const logout = async () => {
    try {
      await api.post('/auth/logout')
    } catch (error) {
      console.error('登出请求失败:', error)
    } finally {
      // 清除本地状态
      user.value = null
      token.value = null
      
      // 跳转到登录页
      await router.push('/auth/login')
    }
  }

  // 获取当前用户信息
  const fetchUser = async () => {
    if (!token.value) {
      return
    }

    try {
      const response = await api.get('/auth/me')
      
      if (response.success) {
        user.value = response.data
      } else {
        // token 可能已过期，清除本地状态
        token.value = null
      }
    } catch (error) {
      console.error('获取用户信息失败:', error)
      token.value = null
    }
  }

  // 检查认证状态
  const checkAuth = async () => {
    if (token.value && !user.value) {
      await fetchUser()
    }
  }

  // 刷新 token
  const refreshToken = async () => {
    try {
      const response = await api.post('/auth/refresh')
      
      if (response.success) {
        token.value = response.data.token
        return true
      }
    } catch (error) {
      console.error('刷新 token 失败:', error)
      await logout()
    }
    
    return false
  }

  return {
    user: readonly(user),
    isLoggedIn,
    isLoading: readonly(isLoading),
    login,
    register,
    logout,
    fetchUser,
    checkAuth,
    refreshToken
  }
}
