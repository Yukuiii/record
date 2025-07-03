/**
 * 类型定义文件 (JavaScript 版本)
 * 这里定义项目中使用的数据结构和常量
 */

// 用户相关类型
export const UserType = {
  // 用户对象结构示例
  createUser: (data = {}) => ({
    id: data.id || '',
    email: data.email || '',
    name: data.name || '',
    avatar: data.avatar || '',
    isActive: data.isActive || true,
    createdAt: data.createdAt || new Date().toISOString(),
    updatedAt: data.updatedAt || new Date().toISOString(),
    lastLoginAt: data.lastLoginAt || null
  })
}

// 记录相关类型
export const RecordType = {
  // 记录类型枚举
  TYPES: {
    INCOME: 'income',
    EXPENSE: 'expense'
  },
  
  // 创建记录对象
  createRecord: (data = {}) => ({
    id: data.id || '',
    userId: data.userId || '',
    type: data.type || RecordType.TYPES.EXPENSE,
    amount: data.amount || 0,
    category: data.category || '',
    description: data.description || '',
    recordDate: data.recordDate || new Date().toISOString().split('T')[0],
    createdAt: data.createdAt || new Date().toISOString(),
    updatedAt: data.updatedAt || new Date().toISOString()
  })
}

// 分类相关类型
export const CategoryType = {
  // 分类类型枚举
  TYPES: {
    INCOME: 'income',
    EXPENSE: 'expense'
  },
  
  // 创建分类对象
  createCategory: (data = {}) => ({
    id: data.id || '',
    name: data.name || '',
    type: data.type || CategoryType.TYPES.EXPENSE,
    description: data.description || '',
    icon: data.icon || '',
    color: data.color || '#3b82f6',
    sortOrder: data.sortOrder || 0,
    isActive: data.isActive || true,
    createdAt: data.createdAt || new Date().toISOString(),
    updatedAt: data.updatedAt || new Date().toISOString()
  })
}

// 预算相关类型
export const BudgetType = {
  // 预算周期枚举
  PERIODS: {
    MONTHLY: 'monthly',
    YEARLY: 'yearly',
    CUSTOM: 'custom'
  },
  
  // 预算状态枚举
  STATUS: {
    ACTIVE: 'active',
    INACTIVE: 'inactive',
    COMPLETED: 'completed'
  },
  
  // 创建预算对象
  createBudget: (data = {}) => ({
    id: data.id || '',
    userId: data.userId || '',
    category: data.category || '',
    amount: data.amount || 0,
    period: data.period || BudgetType.PERIODS.MONTHLY,
    startDate: data.startDate || new Date().toISOString().split('T')[0],
    endDate: data.endDate || '',
    description: data.description || '',
    status: data.status || BudgetType.STATUS.ACTIVE,
    createdAt: data.createdAt || new Date().toISOString(),
    updatedAt: data.updatedAt || new Date().toISOString()
  })
}

// API 响应类型
export const ApiResponseType = {
  // 创建成功响应
  createSuccessResponse: (data, message = '') => ({
    success: true,
    data,
    message
  }),
  
  // 创建错误响应
  createErrorResponse: (message, code = null) => ({
    success: false,
    message,
    code
  })
}

// 表单验证规则
export const ValidationRules = {
  // 邮箱验证
  email: (value) => {
    if (!value) return '请输入邮箱地址'
    if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(value)) {
      return '请输入有效的邮箱地址'
    }
    return ''
  },
  
  // 密码验证
  password: (value) => {
    if (!value) return '请输入密码'
    if (value.length < 6) return '密码至少需要6个字符'
    return ''
  },
  
  // 必填验证
  required: (value, fieldName = '此字段') => {
    if (!value || (typeof value === 'string' && !value.trim())) {
      return `${fieldName}不能为空`
    }
    return ''
  },
  
  // 数字验证
  number: (value, fieldName = '此字段') => {
    if (value === '' || value === null || value === undefined) {
      return `${fieldName}不能为空`
    }
    if (isNaN(Number(value))) {
      return `${fieldName}必须是有效数字`
    }
    return ''
  },
  
  // 正数验证
  positiveNumber: (value, fieldName = '此字段') => {
    const numberError = ValidationRules.number(value, fieldName)
    if (numberError) return numberError
    
    if (Number(value) <= 0) {
      return `${fieldName}必须大于0`
    }
    return ''
  }
}

// 常用常量
export const Constants = {
  // 分页默认值
  PAGINATION: {
    DEFAULT_PAGE: 1,
    DEFAULT_LIMIT: 10,
    MAX_LIMIT: 100
  },
  
  // 日期格式
  DATE_FORMATS: {
    DATE: 'YYYY-MM-DD',
    DATETIME: 'YYYY-MM-DD HH:mm:ss',
    TIME: 'HH:mm:ss'
  },
  
  // 货币符号
  CURRENCY: {
    CNY: '¥',
    USD: '$',
    EUR: '€'
  }
}
