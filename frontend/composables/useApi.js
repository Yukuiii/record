/**
 * API 调用的组合式函数
 */

import { ElMessage } from "element-plus";

export const useApi = () => {
  const config = useRuntimeConfig();
  const baseURL = config.public.apiBase;

  // 创建请求实例
  const createRequest = async (endpoint, options = {}) => {
    const url = `${baseURL}${endpoint}`;

    // 默认请求头
    const defaultHeaders = {
      "Content-Type": "application/json",
    };

    // 获取认证 token（如果存在）
    const token = useCookie("auth-token");
    if (token.value) {
      defaultHeaders["Authorization"] = `Bearer ${token.value}`;
    }

    const requestOptions = {
      ...options,
      headers: {
        ...defaultHeaders,
        ...options.headers,
      },
    };

    try {
      const response = await fetch(url, requestOptions);

      if (!response.ok) {
        const errorData = await response.json().catch(() => ({}));
        const errorMessage = errorData.message || `HTTP ${response.status}`;

        // 自动显示错误 toast
        ElMessage.error(errorMessage);
        throw new Error(errorMessage);
      }

      const data = await response.json();
      return data;
    } catch (error) {
      // 网络错误等其他异常
      if (error.name === "TypeError" && error.message.includes("fetch")) {
        // 网络连接错误
        if (import.meta.client) {
          ElMessage.error("网络请求失败，请检查网络连接");
        }
      } else if (!error.message.includes("HTTP")) {
        // 其他未处理的错误
        if (import.meta.client) {
          ElMessage.error(error.message || "请求失败");
        }
      }
      throw error;
    }
  };

  // GET 请求
  const get = (endpoint, params) => {
    const url = params
      ? `${endpoint}?${new URLSearchParams(params)}`
      : endpoint;
    return createRequest(url, { method: "GET" });
  };

  // POST 请求
  const post = (endpoint, data) => {
    return createRequest(endpoint, {
      method: "POST",
      body: data ? JSON.stringify(data) : undefined,
    });
  };

  // PUT 请求
  const put = (endpoint, data) => {
    return createRequest(endpoint, {
      method: "PUT",
      body: data ? JSON.stringify(data) : undefined,
    });
  };

  // DELETE 请求
  const del = (endpoint) => {
    return createRequest(endpoint, { method: "DELETE" });
  };

  return {
    get,
    post,
    put,
    delete: del,
  };
};
