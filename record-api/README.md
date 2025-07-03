---
title: Personal Record API
emoji: 💰
colorFrom: blue
colorTo: green
sdk: docker
pinned: false
license: mit
app_port: 7860
---

# Personal Record API

个人记账应用后端 API，基于 Go 语言开发，提供完整的记账功能。

## 功能特性

- 🔐 用户认证（注册/登录）
- 📊 分类管理（收入/支出分类）
- 💰 交易记录（增删改查）
- 📈 统计分析（月度/年度统计）
- 🔍 健康检查和监控

## API 接口

### 认证相关

- `POST /api/v1/auth/register` - 用户注册
- `POST /api/v1/auth/login` - 用户登录

### 用户管理

- `GET /api/v1/user/profile` - 获取用户资料
- `PUT /api/v1/user/profile` - 更新用户资料

### 分类管理

- `GET /api/v1/categories` - 获取分类列表
- `POST /api/v1/categories` - 创建分类
- `PUT /api/v1/categories/:id` - 更新分类
- `DELETE /api/v1/categories/:id` - 删除分类

### 交易记录

- `GET /api/v1/transactions` - 获取交易记录
- `POST /api/v1/transactions` - 创建交易记录
- `GET /api/v1/transactions/:id` - 获取单个交易记录
- `PUT /api/v1/transactions/:id` - 更新交易记录
- `DELETE /api/v1/transactions/:id` - 删除交易记录

### 统计分析

- `GET /api/v1/statistics/monthly` - 月度统计
- `GET /api/v1/statistics/yearly` - 年度统计

### 系统接口

- `GET /health` - 健康检查
- `GET /ping` - 简单 ping 测试

## 技术栈

- **后端**: Go + Gin 框架
- **数据库**: PostgreSQL (Supabase)
- **认证**: JWT
- **部署**: Docker

## 环境变量

需要配置以下环境变量：

- `DATABASE_HOST` - 数据库主机
- `DATABASE_PORT` - 数据库端口
- `DATABASE_USER` - 数据库用户
- `DATABASE_PASSWORD` - 数据库密码
- `DATABASE_NAME` - 数据库名
- `DATABASE_SSLMODE` - SSL 模式
- `RECORD_JWT_SECRET` - JWT 密钥

## 使用方法

1. 访问 API 文档查看接口详情
2. 使用 POST /api/v1/auth/register 注册用户
3. 使用 POST /api/v1/auth/login 登录获取 token
4. 在请求头中添加 Authorization: Bearer <token>
5. 调用其他需要认证的接口

## 开发者

个人记账应用开发团队
