#!/bin/bash

# Hugging Face Spaces 部署脚本

set -e

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# 打印带颜色的消息
print_message() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

print_step() {
    echo -e "${BLUE}[STEP]${NC} $1"
}

# 检查必要的工具
check_requirements() {
    print_step "检查部署要求..."
    
    if ! command -v git &> /dev/null; then
        print_error "Git 未安装，请先安装 Git"
        exit 1
    fi
    
    if ! command -v docker &> /dev/null; then
        print_warning "Docker 未安装，无法本地测试"
    fi
    
    print_message "要求检查完成"
}

# 准备部署文件
prepare_files() {
    print_step "准备 Hugging Face Spaces 部署文件..."
    
    # 复制 Hugging Face 专用的 Dockerfile
    if [ -f "Dockerfile.hf" ]; then
        cp Dockerfile.hf Dockerfile
        print_message "已使用 Hugging Face 专用 Dockerfile"
    else
        print_error "Dockerfile.hf 不存在"
        exit 1
    fi
    
    # 检查 README.md 是否包含 Hugging Face 元数据
    if ! grep -q "app_port: 7860" README.md; then
        print_warning "README.md 可能缺少 Hugging Face 元数据"
    fi
    
    print_message "文件准备完成"
}

# 本地测试
test_locally() {
    print_step "本地测试 Docker 镜像..."
    
    if ! command -v docker &> /dev/null; then
        print_warning "Docker 未安装，跳过本地测试"
        return
    fi
    
    # 构建镜像
    print_message "构建 Docker 镜像..."
    docker build -t record-api-hf:test .
    
    # 运行容器进行测试
    print_message "启动测试容器..."
    docker run -d --name record-api-test -p 7860:7860 \
        -e DATABASE_HOST="${DATABASE_HOST:-demo}" \
        -e DATABASE_PORT="${DATABASE_PORT:-6543}" \
        -e DATABASE_USER="${DATABASE_USER:-demo}" \
        -e DATABASE_PASSWORD="${DATABASE_PASSWORD:-demo}" \
        -e DATABASE_NAME="${DATABASE_NAME:-postgres}" \
        -e DATABASE_SSLMODE="${DATABASE_SSLMODE:-require}" \
        -e RECORD_JWT_SECRET="${RECORD_JWT_SECRET:-test-secret}" \
        record-api-hf:test
    
    # 等待容器启动
    sleep 5
    
    # 测试健康检查
    if curl -f http://localhost:7860/health > /dev/null 2>&1; then
        print_message "本地测试通过 ✅"
    else
        print_warning "本地测试失败，但可能是数据库连接问题"
    fi
    
    # 清理测试容器
    docker stop record-api-test > /dev/null 2>&1 || true
    docker rm record-api-test > /dev/null 2>&1 || true
    docker rmi record-api-hf:test > /dev/null 2>&1 || true
    
    print_message "本地测试完成"
}

# 创建 .gitignore（如果不存在）
create_gitignore() {
    if [ ! -f ".gitignore" ]; then
        print_step "创建 .gitignore 文件..."
        cat > .gitignore << EOF
# 构建产物
main
*.exe
*.exe~
*.dll
*.so
*.dylib

# 测试文件
*.test
*.prof

# 依赖目录
vendor/

# 日志文件
logs/
*.log

# 临时文件
*.tmp
*.temp
.DS_Store
Thumbs.db

# IDE 文件
.vscode/
.idea/
*.swp
*.swo
*~

# 环境文件
.env
.env.local
.env.*.local

# Docker 相关（保留 Dockerfile）
docker-compose*.yml
.dockerignore
Dockerfile.hf

# 配置文件
config/config.yaml
EOF
        print_message ".gitignore 文件已创建"
    fi
}

# 显示部署说明
show_deployment_guide() {
    print_step "Hugging Face Spaces 部署指南"
    echo ""
    echo "1. 创建 Hugging Face Spaces："
    echo "   - 访问 https://huggingface.co/new-space"
    echo "   - 选择 Docker SDK"
    echo "   - 设置 Space 名称和描述"
    echo ""
    echo "2. 克隆您的 Space 仓库："
    echo "   git clone https://huggingface.co/spaces/YOUR_USERNAME/YOUR_SPACE_NAME"
    echo ""
    echo "3. 复制文件到 Space 仓库："
    echo "   cp -r * /path/to/your/space/"
    echo ""
    echo "4. 配置环境变量（在 Hugging Face Spaces 设置中）："
    echo "   - DATABASE_HOST"
    echo "   - DATABASE_PORT"
    echo "   - DATABASE_USER"
    echo "   - DATABASE_PASSWORD"
    echo "   - DATABASE_NAME"
    echo "   - DATABASE_SSLMODE"
    echo "   - RECORD_JWT_SECRET"
    echo ""
    echo "5. 提交并推送："
    echo "   git add ."
    echo "   git commit -m \"Initial deployment\""
    echo "   git push"
    echo ""
    echo "6. 等待构建完成，访问您的 Space！"
    echo ""
    print_message "部署指南显示完成"
}

# 创建部署包
create_deployment_package() {
    print_step "创建部署包..."
    
    # 创建临时目录
    TEMP_DIR="hf-deployment-$(date +%Y%m%d-%H%M%S)"
    mkdir -p "$TEMP_DIR"
    
    # 复制必要文件
    cp -r \
        *.go \
        go.mod \
        go.sum \
        README.md \
        Dockerfile \
        api/ \
        config/ \
        controllers/ \
        database/ \
        middleware/ \
        models/ \
        repositories/ \
        services/ \
        utils/ \
        "$TEMP_DIR/" 2>/dev/null || true
    
    # 创建压缩包
    tar -czf "${TEMP_DIR}.tar.gz" "$TEMP_DIR"
    
    # 清理临时目录
    rm -rf "$TEMP_DIR"
    
    print_message "部署包已创建: ${TEMP_DIR}.tar.gz"
}

# 显示帮助信息
show_help() {
    echo "Hugging Face Spaces 部署脚本"
    echo ""
    echo "用法: $0 [选项]"
    echo ""
    echo "选项:"
    echo "  prepare   准备部署文件"
    echo "  test      本地测试 Docker 镜像"
    echo "  package   创建部署包"
    echo "  guide     显示部署指南"
    echo "  all       执行完整的部署准备流程"
    echo "  help      显示此帮助信息"
    echo ""
    echo "示例:"
    echo "  $0 all      # 完整的部署准备"
    echo "  $0 test     # 仅本地测试"
    echo "  $0 guide    # 显示部署指南"
}

# 主函数
main() {
    case "${1:-help}" in
        "prepare")
            check_requirements
            prepare_files
            create_gitignore
            ;;
        "test")
            check_requirements
            prepare_files
            test_locally
            ;;
        "package")
            check_requirements
            prepare_files
            create_deployment_package
            ;;
        "guide")
            show_deployment_guide
            ;;
        "all")
            check_requirements
            prepare_files
            create_gitignore
            test_locally
            create_deployment_package
            show_deployment_guide
            ;;
        "help"|*)
            show_help
            ;;
    esac
}

# 执行主函数
main "$@"
