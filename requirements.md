# 个人记账应用需求分析文档

## 1. 项目概述

### 1.1 项目名称

个人记账管理系统

### 1.2 项目描述

基于 Web 的个人记账应用，帮助用户记录和管理个人财务状况，提供直观的收支统计和分析功能。

### 1.3 技术栈

- **前端**: Nuxt.js
- **构建工具**: Vite
- **后端**: Go + Gin 框架
- **数据库**: Supabase
- **目标平台**: Web 应用

## 2. 用户需求

### 2.1 目标用户

- 主要面向个人用户
- 需要管理个人财务的普通用户
- 希望了解个人消费习惯的用户

### 2.2 用户场景

- 日常收支记录
- 月度/年度财务回顾
- 多账户资金管理
- 消费分析和预算控制

## 3. 功能需求

### 3.1 用户管理模块

- **用户注册**
  - 邮箱注册
  - 手机号注册
  - 用户信息验证
- **用户登录**
  - 邮箱/手机号登录
  - 密码找回功能
  - 记住登录状态
- **用户信息管理**
  - 个人资料编辑
  - 密码修改
  - 账户安全设置

### 3.2 收支记录模块

- **收入记录**
  - 收入金额输入
  - 收入来源分类（工资、投资、其他等）
  - 收入时间记录
  - 备注信息
- **支出记录**
  - 支出金额输入
  - 支出分类（餐饮、交通、购物、娱乐等）
  - 支出时间记录
  - 备注信息
- **记录管理**
  - 查看历史记录
  - 编辑/删除记录
  - 记录搜索和筛选

### 3.3 统计分析模块

- **月度统计**
  - 月收入统计
  - 月支出统计
  - 月结余计算
  - 分类支出占比
- **年度统计**
  - 年收入统计
  - 年支出统计
  - 年结余计算
  - 月度趋势对比
- **数据可视化**
  - 收支趋势图表
  - 分类支出饼图
  - 月度对比柱状图

## 4. 非功能需求

### 4.1 性能需求

- 页面加载时间 < 3 秒
- 数据查询响应时间 < 0.5 秒
- 支持并发用户数: 500+

### 4.2 安全需求

- 用户密码加密存储
- 数据传输 HTTPS 加密
- 用户会话管理与 JWT 认证
- 防止 SQL 注入攻击
- 请求频率限制

### 4.3 兼容性需求

- 支持主流浏览器（Chrome、Firefox、Safari、Edge）
- 响应式设计，适配移动端
- 支持常见屏幕分辨率

### 4.4 可用性需求

- 界面简洁直观
- 操作流程简单
- 错误提示友好
- 支持键盘快捷操作

## 5. 数据库设计

### 5.1 数据模型设计

#### 5.1.1 用户模型 (users)

```go
type User struct {
    ID           uint      `gorm:"primaryKey" json:"id"`
    Email        string    `gorm:"uniqueIndex;size:255" json:"email"`
    Phone        string    `gorm:"uniqueIndex;size:20" json:"phone"`
    Password     string    `gorm:"size:255" json:"-"`
    Nickname     string    `gorm:"size:50" json:"nickname"`
    Avatar       string    `gorm:"size:255" json:"avatar"`
    Gender       string    `gorm:"size:10" json:"gender"`
    Birthday     time.Time `json:"birthday"`
    RegisterTime time.Time `gorm:"autoCreateTime" json:"register_time"`
    LastLogin    time.Time `json:"last_login"`
    Status       int       `gorm:"default:1" json:"status"` // 1: 正常, 0: 禁用
    CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
```

#### 5.1.2 分类模型 (categories)

```go
type Category struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    Name      string    `gorm:"size:50" json:"name"`
    Type      string    `gorm:"size:20" json:"type"` // income: 收入, expense: 支出
    Icon      string    `gorm:"size:50" json:"icon"`
    Color     string    `gorm:"size:20" json:"color"`
    IsDefault bool      `gorm:"default:false" json:"is_default"`
    CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
```

分类定义:

- 交通
- 餐饮
- 购物
- 住房
- 医疗
- 教育
- 旅游
- 人情
- 其他

#### 5.1.3 交易记录模型 (transactions)

```go
type Transaction struct {
    ID          uint      `gorm:"primaryKey" json:"id"`
    UserID      uint      `json:"user_id"`
    CategoryID  uint      `json:"category_id"`
    Amount      float64   `json:"amount"`
    Type        string    `gorm:"size:20" json:"type"` // income: 收入, expense: 支出
    Description string    `gorm:"size:255" json:"description"`
    RecordTime  time.Time `json:"record_time"`
    Location    string    `gorm:"size:255" json:"location"`
    ImageURL    string    `gorm:"size:255" json:"image_url"`
    Tags        string    `gorm:"size:255" json:"tags"` // 以逗号分隔的标签
    CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`

    // 关联
    Category    Category  `gorm:"foreignKey:CategoryID" json:"category"`
}
```

### 5.2 数据库索引设计

- **users 表**:
  - 主键索引: `id`
  - 唯一索引: `email`, `phone`
- **categories 表**:
  - 主键索引: `id`
  - 索引: `type`
  - 唯一索引: `name` (确保分类名称唯一)
- **transactions 表**:
  - 主键索引: `id`
  - 索引: `user_id`, `category_id`, `type`, `record_time`
  - 复合索引: `(user_id, record_time)`, `(user_id, category_id)`

### 5.3 数据关系

- 用户(users) 1:N 交易记录(transactions)
- 分类(categories) 1:N 交易记录(transactions)

## 6. API 接口设计

### 6.1 通用响应格式

```json
{
  "code": 200, // 状态码: 200成功, 400请求错误, 401未授权, 403禁止访问, 404未找到, 500服务器错误
  "message": "success", // 响应消息
  "data": {} // 响应数据
}
```

### 6.2 用户相关接口

#### 6.2.1 用户注册

- **请求**: `POST /api/v1/auth/register`
- **请求体**:

```json
{
  "email": "user@example.com",
  "phone": "13800138000",
  "password": "password123",
  "nickname": "用户昵称"
}
```

- **响应**:

```json
{
  "code": 200,
  "message": "注册成功",
  "data": {
    "user_id": 1,
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

#### 6.2.2 用户登录

- **请求**: `POST /api/v1/auth/login`
- **请求体**:

```json
{
  "account": "user@example.com", // 邮箱或手机号
  "password": "password123"
}
```

- **响应**:

```json
{
  "code": 200,
  "message": "登录成功",
  "data": {
    "user": {
      "id": 1,
      "email": "user@example.com",
      "phone": "13800138000",
      "nickname": "用户昵称",
      "avatar": "https://example.com/avatar.jpg"
    },
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

#### 6.2.3 获取用户信息

- **请求**: `GET /api/v1/user/profile`
- **请求头**: `Authorization: Bearer {token}`
- **响应**:

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "user": {
      "id": 1,
      "email": "user@example.com",
      "phone": "13800138000",
      "nickname": "用户昵称",
      "avatar": "https://example.com/avatar.jpg",
      "gender": "male",
      "birthday": "1990-01-01T00:00:00Z",
      "register_time": "2023-01-01T00:00:00Z",
      "last_login": "2023-06-01T00:00:00Z"
    }
  }
}
```

#### 6.2.4 更新用户信息

- **请求**: `PUT /api/v1/user/profile`
- **请求头**: `Authorization: Bearer {token}`
- **请求体**:

```json
{
  "nickname": "新昵称",
  "avatar": "https://example.com/new-avatar.jpg",
  "gender": "female",
  "birthday": "1992-01-01T00:00:00Z"
}
```

- **响应**:

```json
{
  "code": 200,
  "message": "更新成功",
  "data": {
    "user": {
      "id": 1,
      "nickname": "新昵称",
      "avatar": "https://example.com/new-avatar.jpg",
      "gender": "female",
      "birthday": "1992-01-01T00:00:00Z"
    }
  }
}
```

### 6.3 分类相关接口

#### 6.3.1 获取分类列表

- **请求**: `GET /api/v1/categories?type=expense`
- **请求头**: `Authorization: Bearer {token}`
- **参数**: `type` - 可选，收入(income)或支出(expense)
- **响应**:

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "categories": [
      {
        "id": 1,
        "name": "餐饮",
        "type": "expense",
        "icon": "food",
        "color": "#FF5722",
        "is_default": true
      },
      {
        "id": 2,
        "name": "交通",
        "type": "expense",
        "icon": "car",
        "color": "#2196F3",
        "is_default": true
      }
    ]
  }
}
```

#### 6.3.2 创建分类

- **请求**: `POST /api/v1/categories`
- **请求头**: `Authorization: Bearer {token}`
- **请求体**:

```json
{
  "name": "旅游",
  "type": "expense",
  "icon": "travel",
  "color": "#9C27B0"
}
```

- **响应**:

```json
{
  "code": 200,
  "message": "创建成功",
  "data": {
    "category": {
      "id": 10,
      "name": "旅游",
      "type": "expense",
      "icon": "travel",
      "color": "#9C27B0",
      "is_default": false
    }
  }
}
```

#### 6.3.3 更新分类

- **请求**: `PUT /api/v1/categories/{id}`
- **请求头**: `Authorization: Bearer {token}`
- **请求体**:

```json
{
  "name": "度假旅游",
  "icon": "vacation",
  "color": "#673AB7"
}
```

- **响应**:

```json
{
  "code": 200,
  "message": "更新成功",
  "data": {
    "category": {
      "id": 10,
      "name": "度假旅游",
      "type": "expense",
      "icon": "vacation",
      "color": "#673AB7",
      "is_default": false
    }
  }
}
```

#### 6.3.4 删除分类

- **请求**: `DELETE /api/v1/categories/{id}`
- **请求头**: `Authorization: Bearer {token}`
- **响应**:

```json
{
  "code": 200,
  "message": "删除成功",
  "data": null
}
```

- **注意**: 系统默认分类不可删除，已被交易记录使用的分类也不可删除

### 6.4 交易记录相关接口

#### 6.4.1 创建交易记录

- **请求**: `POST /api/v1/transactions`
- **请求头**: `Authorization: Bearer {token}`
- **请求体**:

```json
{
  "category_id": 1,
  "amount": 100.5,
  "type": "expense",
  "description": "午餐",
  "record_time": "2023-06-01T12:30:00Z",
  "location": "公司附近餐厅",
  "tags": "工作餐,午餐"
}
```

- **响应**:

```json
{
  "code": 200,
  "message": "创建成功",
  "data": {
    "transaction": {
      "id": 1,
      "category_id": 1,
      "amount": 100.5,
      "type": "expense",
      "description": "午餐",
      "record_time": "2023-06-01T12:30:00Z",
      "location": "公司附近餐厅",
      "tags": "工作餐,午餐",
      "category": {
        "id": 1,
        "name": "餐饮",
        "type": "expense",
        "icon": "food",
        "color": "#FF5722"
      }
    }
  }
}
```

#### 6.4.2 获取交易记录列表

- **请求**: `GET /api/v1/transactions?type=expense&start_date=2023-06-01&end_date=2023-06-30&category_id=1&page=1&page_size=20`
- **请求头**: `Authorization: Bearer {token}`
- **参数**:
  - `type`: 可选，收入(income)或支出(expense)
  - `start_date`: 可选，开始日期
  - `end_date`: 可选，结束日期
  - `category_id`: 可选，分类 ID
  - `page`: 可选，页码，默认 1
  - `page_size`: 可选，每页条数，默认 20
- **响应**:

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "total": 100,
    "page": 1,
    "page_size": 20,
    "transactions": [
      {
        "id": 1,
        "category_id": 1,
        "amount": 100.5,
        "type": "expense",
        "description": "午餐",
        "record_time": "2023-06-01T12:30:00Z",
        "tags": "工作餐,午餐",
        "category": {
          "id": 1,
          "name": "餐饮",
          "type": "expense",
          "icon": "food",
          "color": "#FF5722"
        }
      }
    ]
  }
}
```

#### 6.4.3 获取交易记录详情

- **请求**: `GET /api/v1/transactions/{id}`
- **请求头**: `Authorization: Bearer {token}`
- **响应**:

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "transaction": {
      "id": 1,
      "category_id": 1,
      "amount": 100.5,
      "type": "expense",
      "description": "午餐",
      "record_time": "2023-06-01T12:30:00Z",
      "location": "公司附近餐厅",
      "image_url": "https://example.com/receipt.jpg",
      "tags": "工作餐,午餐",
      "category": {
        "id": 1,
        "name": "餐饮",
        "type": "expense",
        "icon": "food",
        "color": "#FF5722"
      }
    }
  }
}
```

#### 6.4.4 更新交易记录

- **请求**: `PUT /api/v1/transactions/{id}`
- **请求头**: `Authorization: Bearer {token}`
- **请求体**:

```json
{
  "category_id": 2,
  "amount": 120.0,
  "description": "修改后的午餐",
  "record_time": "2023-06-01T12:45:00Z"
}
```

- **响应**:

```json
{
  "code": 200,
  "message": "更新成功",
  "data": {
    "transaction": {
      "id": 1,
      "category_id": 2,
      "amount": 120.0,
      "description": "修改后的午餐",
      "record_time": "2023-06-01T12:45:00Z"
    }
  }
}
```

#### 6.4.5 删除交易记录

- **请求**: `DELETE /api/v1/transactions/{id}`
- **请求头**: `Authorization: Bearer {token}`
- **响应**:

```json
{
  "code": 200,
  "message": "删除成功",
  "data": null
}
```

### 6.5 统计相关接口

#### 6.5.1 获取月度统计

- **请求**: `GET /api/v1/statistics/monthly?year=2023&month=6`
- **请求头**: `Authorization: Bearer {token}`
- **响应**:

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "year": 2023,
    "month": 6,
    "income_total": 8000.0,
    "expense_total": 5000.0,
    "balance": 3000.0,
    "expense_by_category": [
      {
        "category_id": 1,
        "category_name": "餐饮",
        "icon": "food",
        "color": "#FF5722",
        "amount": 2000.0,
        "percentage": 40
      },
      {
        "category_id": 2,
        "category_name": "交通",
        "icon": "car",
        "color": "#2196F3",
        "amount": 1000.0,
        "percentage": 20
      }
    ],
    "income_by_category": [
      {
        "category_id": 5,
        "category_name": "工资",
        "icon": "salary",
        "color": "#4CAF50",
        "amount": 7000.0,
        "percentage": 87.5
      },
      {
        "category_id": 6,
        "category_name": "投资",
        "icon": "investment",
        "color": "#FFC107",
        "amount": 1000.0,
        "percentage": 12.5
      }
    ],
    "daily_expenses": [
      {
        "date": "2023-06-01",
        "amount": 200.0
      },
      {
        "date": "2023-06-02",
        "amount": 150.0
      }
    ]
  }
}
```

#### 6.5.2 获取年度统计

- **请求**: `GET /api/v1/statistics/yearly?year=2023`
- **请求头**: `Authorization: Bearer {token}`
- **响应**:

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "year": 2023,
    "income_total": 96000.0,
    "expense_total": 60000.0,
    "balance": 36000.0,
    "monthly_summary": [
      {
        "month": 1,
        "income": 8000.0,
        "expense": 5000.0,
        "balance": 3000.0
      },
      {
        "month": 2,
        "income": 8000.0,
        "expense": 5500.0,
        "balance": 2500.0
      }
    ],
    "expense_by_category": [
      {
        "category_id": 1,
        "category_name": "餐饮",
        "icon": "food",
        "color": "#FF5722",
        "amount": 24000.0,
        "percentage": 40
      }
    ],
    "income_by_category": [
      {
        "category_id": 5,
        "category_name": "工资",
        "icon": "salary",
        "color": "#4CAF50",
        "amount": 84000.0,
        "percentage": 87.5
      }
    ]
  }
}
```

## 7. 开发计划

### 7.1 开发阶段

1. **第一阶段**: 用户管理和 Go+Gin 基础架构搭建
2. **第二阶段**: 账户管理和收支记录功能
3. **第三阶段**: 统计分析和数据可视化
4. **第四阶段**: 优化和测试

### 7.2 预估工期

- 总开发周期: 3-5 周
- 测试和优化: 1-2 周

## 8. 风险评估

### 8.1 技术风险

- Supabase 服务稳定性
- Go 与前端框架的数据交互
- 数据安全和隐私保护
- 团队对 Go 语言的熟悉程度

### 8.2 业务风险

- 用户需求变更
- 竞品功能对比
- 用户体验优化需求

## 9. 后续扩展

### 9.1 可能的功能扩展

- 数据导入导出
- 多币种支持
- 投资组合管理
- 家庭共享账本

### 9.2 技术扩展

- 移动端 App 开发
- 数据备份和恢复
- 第三方银行接口集成
- AI 智能分类推荐
- Go 微服务架构拆分
- 容器化部署与 CI/CD 流程

## 10. Go 后端技术实现细节

### 10.1 项目结构

```
project-root/
├── api/           # API接口定义
├── config/        # 配置文件
├── controllers/   # 控制器
├── docs/          # API文档(Swagger)
├── middleware/    # 中间件
├── models/        # 数据模型
├── repositories/  # 数据访问层
├── routes/        # 路由定义
├── services/      # 业务逻辑层
├── utils/         # 工具函数
├── main.go        # 应用入口
└── go.mod         # Go模块定义
```

### 10.2 核心依赖

- **gin**: Web 框架
- **gorm**: ORM 库
- **jwt-go**: JWT 认证
- **validator**: 请求验证
- **zap**: 日志管理
- **viper**: 配置管理
- **testify**: 单元测试

### 10.3 性能优化

- 使用 Go 协程处理并发请求
- 实现请求缓存机制
- 数据库连接池优化
- 使用适当的索引提升查询性能

### 10.4 安全措施

- 请求频率限制中间件
- CORS 安全配置
- 参数验证与清洁
- 错误处理与日志记录
- 敏感信息加密存储
