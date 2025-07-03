/**
 * 记录管理相关的组合式函数
 */

export const useRecords = () => {
  const api = useApi()
  
  // 记录列表状态
  const records = ref([])
  const isLoading = ref(false)
  const total = ref(0)
  const currentPage = ref(1)
  const pageSize = ref(10)

  // 获取记录列表
  const fetchRecords = async (filters = {}) => {
    isLoading.value = true
    
    try {
      const params = {
        page: currentPage.value,
        limit: pageSize.value,
        ...filters
      }
      
      const response = await api.get('/records', params)

      if (response.success) {
        records.value = response.data.records
        total.value = response.data.total
        currentPage.value = response.data.page
      }
    } catch (error) {
      console.error('获取记录失败:', error)
      throw error
    } finally {
      isLoading.value = false
    }
  }

  // 创建记录
  const createRecord = async (data) => {
    try {
      const response = await api.post('/records', data)
      
      if (response.success) {
        // 添加到列表开头
        records.value.unshift(response.data)
        total.value += 1
        return { success: true, data: response.data }
      } else {
        throw new Error(response.message || '创建记录失败')
      }
    } catch (error) {
      console.error('创建记录失败:', error)
      return {
        success: false,
        message: error instanceof Error ? error.message : '创建记录失败'
      }
    }
  }

  // 更新记录
  const updateRecord = async (id, data) => {
    try {
      const response = await api.put(`/records/${id}`, data)
      
      if (response.success) {
        // 更新列表中的记录
        const index = records.value.findIndex(r => r.id === id)
        if (index !== -1) {
          records.value[index] = response.data
        }
        return { success: true, data: response.data }
      } else {
        throw new Error(response.message || '更新记录失败')
      }
    } catch (error) {
      console.error('更新记录失败:', error)
      return {
        success: false,
        message: error instanceof Error ? error.message : '更新记录失败'
      }
    }
  }

  // 删除记录
  const deleteRecord = async (id) => {
    try {
      const response = await api.delete(`/records/${id}`)
      
      if (response.success) {
        // 从列表中移除
        const index = records.value.findIndex(r => r.id === id)
        if (index !== -1) {
          records.value.splice(index, 1)
          total.value -= 1
        }
        return { success: true }
      } else {
        throw new Error(response.message || '删除记录失败')
      }
    } catch (error) {
      console.error('删除记录失败:', error)
      return {
        success: false,
        message: error instanceof Error ? error.message : '删除记录失败'
      }
    }
  }

  // 获取单条记录
  const getRecord = async (id) => {
    try {
      const response = await api.get(`/records/${id}`)
      
      if (response.success) {
        return { success: true, data: response.data }
      } else {
        throw new Error(response.message || '获取记录失败')
      }
    } catch (error) {
      console.error('获取记录失败:', error)
      return {
        success: false,
        message: error instanceof Error ? error.message : '获取记录失败'
      }
    }
  }

  // 计算统计信息
  const getStatistics = computed(() => {
    const income = records.value
      .filter(r => r.type === 'income')
      .reduce((sum, r) => sum + r.amount, 0)
    
    const expense = records.value
      .filter(r => r.type === 'expense')
      .reduce((sum, r) => sum + r.amount, 0)
    
    return {
      income,
      expense,
      balance: income - expense,
      total: records.value.length
    }
  })

  // 分页相关
  const totalPages = computed(() => Math.ceil(total.value / pageSize.value))
  
  const goToPage = (page) => {
    if (page >= 1 && page <= totalPages.value) {
      currentPage.value = page
      fetchRecords()
    }
  }

  const nextPage = () => {
    if (currentPage.value < totalPages.value) {
      goToPage(currentPage.value + 1)
    }
  }

  const prevPage = () => {
    if (currentPage.value > 1) {
      goToPage(currentPage.value - 1)
    }
  }

  return {
    records: readonly(records),
    isLoading: readonly(isLoading),
    total: readonly(total),
    currentPage: readonly(currentPage),
    pageSize: readonly(pageSize),
    totalPages,
    statistics: getStatistics,
    fetchRecords,
    createRecord,
    updateRecord,
    deleteRecord,
    getRecord,
    goToPage,
    nextPage,
    prevPage
  }
}
