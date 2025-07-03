# Docker 部署指南

本文档介绍如何使用 Docker 部署个人记账应用后端。

## 文件说明

- `Dockerfile` - 多阶段构建的 Docker 镜像定义
- `docker-compose.yml` - 开发环境 Docker Compose 配置
- `docker-compose.prod.yml` - 生产环境 Docker Compose 配置
- `.dockerignore` - Docker 构建忽略文件

## 快速开始

### 1. 准备环境变量

复制环境变量示例文件：
```bash
cp .env.example .env
```

编辑 `.env` 文件，填入您的 Supabase 数据库信息：
```bash
DATABASE_HOST=aws-0-ap-southeast-1.pooler.supabase.com
DATABASE_PORT=6543
DATABASE_USER=postgres.your-project-ref
DATABASE_PASSWORD=your-supabase-password
DATABASE_NAME=postgres
DATABASE_SSLMODE=require
RECORD_JWT_SECRET=your-jwt-secret-key
```

### 2. 构建和运行（开发环境）

```bash
# 构建并启动
docker-compose up --build

# 后台运行
docker-compose up -d --build

# 查看日志
docker-compose logs -f record-api

# 停止服务
docker-compose down
```

### 3. 生产环境部署

```bash
# 使用生产环境配置
docker-compose -f docker-compose.prod.yml up -d --build

# 查看状态
docker-compose -f docker-compose.prod.yml ps

# 查看日志
docker-compose -f docker-compose.prod.yml logs -f
```

## 单独使用 Docker

### 构建镜像

```bash
# 构建镜像
docker build -t record-api:latest .

# 查看镜像
docker images | grep record-api
```

### 运行容器

```bash
# 运行容器
docker run -d \
  --name record-api \
  -p 8080:8080 \
  -e DATABASE_HOST=your-supabase-host \
  -e DATABASE_PORT=6543 \
  -e DATABASE_USER=your-user \
  -e DATABASE_PASSWORD=your-password \
  -e DATABASE_NAME=postgres \
  -e DATABASE_SSLMODE=require \
  -e RECORD_JWT_SECRET=your-jwt-secret \
  -v $(pwd)/logs:/app/logs \
  record-api:latest

# 查看容器状态
docker ps

# 查看日志
docker logs -f record-api

# 停止容器
docker stop record-api

# 删除容器
docker rm record-api
```

## 健康检查

容器启动后，可以通过以下方式检查服务状态：

```bash
# 健康检查
curl http://localhost:8080/health

# 简单 ping
curl http://localhost:8080/ping

# 查看 Docker 健康状态
docker ps --format "table {{.Names}}\t{{.Status}}"
```

## 日志管理

### 查看日志

```bash
# Docker Compose 日志
docker-compose logs -f record-api

# 容器日志
docker logs -f record-api

# 应用日志文件
tail -f logs/info_$(date +%Y-%m-%d).log
tail -f logs/error_$(date +%Y-%m-%d).log
```

### 日志轮转

建议在生产环境中配置日志轮转：

```bash
# 添加到 crontab
0 0 * * * find /path/to/logs -name "*.log" -mtime +7 -delete
```

## 环境变量说明

| 变量名 | 说明 | 默认值 | 必需 |
|--------|------|--------|------|
| `DATABASE_HOST` | 数据库主机 | - | ✅ |
| `DATABASE_PORT` | 数据库端口 | 6543 | ❌ |
| `DATABASE_USER` | 数据库用户 | - | ✅ |
| `DATABASE_PASSWORD` | 数据库密码 | - | ✅ |
| `DATABASE_NAME` | 数据库名 | postgres | ❌ |
| `DATABASE_SSLMODE` | SSL 模式 | require | ❌ |
| `RECORD_SERVER_PORT` | 服务端口 | 8080 | ❌ |
| `RECORD_SERVER_MODE` | 运行模式 | debug | ❌ |
| `RECORD_JWT_SECRET` | JWT 密钥 | - | ✅ |
| `RECORD_JWT_EXPIRETIME` | JWT 过期时间(小时) | 24 | ❌ |

## 故障排除

### 常见问题

1. **容器启动失败**
   ```bash
   # 查看详细错误
   docker-compose logs record-api
   ```

2. **数据库连接失败**
   - 检查环境变量是否正确
   - 确认 Supabase 数据库可访问
   - 验证网络连接

3. **端口冲突**
   ```bash
   # 修改端口映射
   ports:
     - "8081:8080"  # 使用不同的主机端口
   ```

4. **权限问题**
   ```bash
   # 检查日志目录权限
   chmod 755 logs/
   ```

### 调试模式

```bash
# 进入容器调试
docker exec -it record-api sh

# 查看容器内文件
docker exec record-api ls -la /app

# 查看环境变量
docker exec record-api env | grep DATABASE
```

## 性能优化

### 生产环境建议

1. **资源限制**：在 `docker-compose.prod.yml` 中已配置
2. **健康检查**：自动重启不健康的容器
3. **日志管理**：定期清理旧日志文件
4. **监控**：集成 Prometheus/Grafana 监控

### 镜像优化

- 使用多阶段构建减小镜像大小
- 使用 Alpine Linux 基础镜像
- 只复制必要的文件
- 使用非 root 用户运行
