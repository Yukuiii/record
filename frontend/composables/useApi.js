/**
 * API 调用的组合式函数
 */

export const useApi = () => {
  const config = useRuntimeConfig()
  const baseURL = config.public.apiBase

  // 创建请求实例
  const createRequest = async (endpoint, options = {}) => {
    const url = `${baseURL}${endpoint}`
    
    // 默认请求头
    const defaultHeaders = {
      'Content-Type': 'application/json',
    }

    // 获取认证 token（如果存在）
    const token = useCookie('auth-token')
    if (token.value) {
      defaultHeaders['Authorization'] = `Bearer ${token.value}`
    }

    const requestOptions = {
      ...options,
      headers: {
        ...defaultHeaders,
        ...options.headers,
      },
    }

    try {
      const response = await fetch(url, requestOptions)
      
      if (!response.ok) {
        const errorData = await response.json().catch(() => ({}))
        throw new Error(errorData.message || `HTTP ${response.status}`)
      }

      const data = await response.json()
      return data
    } catch (error) {
      console.error('API 请求失败:', error)
      throw error
    }
  }

  // GET 请求
  const get = (endpoint, params) => {
    const url = params ? `${endpoint}?${new URLSearchParams(params)}` : endpoint
    return createRequest(url, { method: 'GET' })
  }

  // POST 请求
  const post = (endpoint, data) => {
    return createRequest(endpoint, {
      method: 'POST',
      body: data ? JSON.stringify(data) : undefined,
    })
  }

  // PUT 请求
  const put = (endpoint, data) => {
    return createRequest(endpoint, {
      method: 'PUT',
      body: data ? JSON.stringify(data) : undefined,
    })
  }

  // DELETE 请求
  const del = (endpoint) => {
    return createRequest(endpoint, { method: 'DELETE' })
  }

  return {
    get,
    post,
    put,
    delete: del,
  }
}
