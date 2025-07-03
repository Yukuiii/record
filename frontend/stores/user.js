/**
 * 用户状态管理 (Pinia Store)
 * 注意：这个文件需要在安装 Pinia 后才能正常使用
 */

// 当前使用 Nuxt 3 内置的状态管理，如果需要 Pinia，需要安装 @pinia/nuxt 模块

// 示例：使用 Nuxt 3 内置状态管理
export const useUserStore = () => {
  // 用户信息
  const user = useState("user", () => null);

  // 用户偏好设置
  const preferences = useState("user.preferences", () => ({
    theme: "light",
    language: "zh-CN",
    currency: "CNY",
  }));

  // 设置用户信息
  const setUser = (userData) => {
    user.value = userData;
  };

  // 清除用户信息
  const clearUser = () => {
    user.value = null;
  };

  // 更新用户偏好
  const updatePreferences = (newPreferences) => {
    preferences.value = {
      ...preferences.value,
      ...newPreferences,
    };
  };

  // 计算属性
  const isLoggedIn = computed(() => !!user.value);
  const userName = computed(() => user.value?.name || "");
  const userEmail = computed(() => user.value?.email || "");

  return {
    // 状态
    user: readonly(user),
    preferences: readonly(preferences),

    // 计算属性
    isLoggedIn,
    userName,
    userEmail,

    // 方法
    setUser,
    clearUser,
    updatePreferences,
  };
};

// 如果使用 Pinia，可以这样定义：
/*
import { defineStore } from 'pinia'

export const useUserStore = defineStore('user', () => {
  // 状态
  const user = ref(null)
  const preferences = ref({
    theme: 'light',
    language: 'zh-CN',
    currency: 'CNY'
  })
  
  // 计算属性
  const isLoggedIn = computed(() => !!user.value)
  const userName = computed(() => user.value?.name || '')
  const userEmail = computed(() => user.value?.email || '')
  
  // 方法
  const setUser = (userData) => {
    user.value = userData
  }
  
  const clearUser = () => {
    user.value = null
  }
  
  const updatePreferences = (newPreferences) => {
    preferences.value = {
      ...preferences.value,
      ...newPreferences
    }
  }
  
  return {
    user,
    preferences,
    isLoggedIn,
    userName,
    userEmail,
    setUser,
    clearUser,
    updatePreferences
  }
})
*/
