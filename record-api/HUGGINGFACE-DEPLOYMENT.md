# Hugging Face Spaces 部署指南

本指南将帮助您将个人记账应用后端部署到 Hugging Face Spaces。

## 🚀 快速部署

### 1. 准备部署文件

```bash
# 使用部署脚本准备所有文件
chmod +x deploy-hf.sh
./deploy-hf.sh all
```

### 2. 创建 Hugging Face Space

1. 访问 [Hugging Face Spaces](https://huggingface.co/new-space)
2. 填写以下信息：
   - **Space name**: `personal-record-api`
   - **License**: `MIT`
   - **SDK**: 选择 `Docker`
   - **Hardware**: `CPU basic` (免费)
3. 点击 **Create Space**

### 3. 克隆 Space 仓库

```bash
# 替换为您的用户名和 Space 名称
git clone https://huggingface.co/spaces/YOUR_USERNAME/personal-record-api
cd personal-record-api
```

### 4. 复制项目文件

```bash
# 从您的项目目录复制文件
cp /path/to/your/record-api/* ./
```

### 5. 配置环境变量

在 Hugging Face Spaces 的设置页面中添加以下环境变量：

| 变量名 | 值 | 说明 |
|--------|-----|------|
| `DATABASE_HOST` | `aws-0-ap-southeast-1.pooler.supabase.com` | Supabase 主机 |
| `DATABASE_PORT` | `6543` | Supabase 端口 |
| `DATABASE_USER` | `postgres.your-ref` | Supabase 用户名 |
| `DATABASE_PASSWORD` | `your-password` | Supabase 密码 |
| `DATABASE_NAME` | `postgres` | 数据库名 |
| `DATABASE_SSLMODE` | `require` | SSL 模式 |
| `RECORD_JWT_SECRET` | `your-strong-secret` | JWT 密钥 |
| `RECORD_SERVER_MODE` | `release` | 运行模式 |

### 6. 提交并部署

```bash
git add .
git commit -m "Initial deployment to Hugging Face Spaces"
git push
```

## 📋 部署要求

### Hugging Face Spaces 特殊要求

1. **端口**: 必须使用 7860 端口
2. **Dockerfile**: 必须在根目录
3. **README.md**: 必须包含 YAML 元数据
4. **资源限制**: 免费版有 CPU 和内存限制

### 项目文件结构

```
your-space/
├── README.md              # 包含 HF 元数据
├── Dockerfile             # HF 专用 Dockerfile
├── main.go               # Go 主程序
├── go.mod                # Go 模块文件
├── go.sum                # Go 依赖锁定
├── api/                  # API 定义
├── config/               # 配置管理
├── controllers/          # 控制器
├── database/             # 数据库
├── middleware/           # 中间件
├── models/               # 数据模型
├── repositories/         # 数据访问层
├── services/             # 业务逻辑层
└── utils/                # 工具函数
```

## 🔧 配置说明

### README.md 元数据

```yaml
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
```

### Dockerfile 关键配置

- 使用多阶段构建优化镜像大小
- 监听 7860 端口（HF 要求）
- 使用非 root 用户运行
- 包含健康检查

## 🧪 本地测试

在部署前，建议先本地测试：

```bash
# 构建镜像
docker build -f Dockerfile.hf -t record-api-hf .

# 运行测试
docker run -p 7860:7860 \
  -e DATABASE_HOST=your-host \
  -e DATABASE_PASSWORD=your-password \
  -e RECORD_JWT_SECRET=test-secret \
  record-api-hf

# 测试接口
curl http://localhost:7860/health
```

## 📊 监控和调试

### 查看构建日志

1. 进入您的 Space 页面
2. 点击 **Logs** 标签
3. 查看构建和运行日志

### 常见问题

1. **构建失败**
   - 检查 Dockerfile 语法
   - 确认所有文件都已提交

2. **应用无法启动**
   - 检查环境变量配置
   - 查看运行日志

3. **数据库连接失败**
   - 验证 Supabase 配置
   - 检查网络连接

### 健康检查

访问以下端点检查应用状态：

- `https://your-username-personal-record-api.hf.space/health`
- `https://your-username-personal-record-api.hf.space/ping`

## 🔄 更新部署

```bash
# 修改代码后重新部署
git add .
git commit -m "Update: description of changes"
git push
```

## 💡 优化建议

### 性能优化

1. **镜像优化**
   - 使用 Alpine Linux
   - 多阶段构建
   - 最小化依赖

2. **应用优化**
   - 启用 release 模式
   - 配置适当的日志级别
   - 使用连接池

### 安全建议

1. **环境变量**
   - 使用强密码
   - 定期轮换密钥
   - 不在代码中硬编码敏感信息

2. **网络安全**
   - 启用 HTTPS（HF 自动提供）
   - 配置 CORS
   - 实施请求限流

## 📞 支持

如果遇到问题：

1. 查看 [Hugging Face Spaces 文档](https://huggingface.co/docs/hub/spaces)
2. 检查项目的 GitHub Issues
3. 联系开发团队

## 🎉 部署完成

部署成功后，您的 API 将在以下地址可用：

`https://your-username-personal-record-api.hf.space`

现在您可以：
- 测试 API 接口
- 集成前端应用
- 分享给其他用户使用
