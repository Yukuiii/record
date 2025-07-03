#!/bin/bash

# 个人记账应用 Docker 构建脚本

set -e

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
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

# 检查 Docker 是否安装
check_docker() {
    if ! command -v docker &> /dev/null; then
        print_error "Docker 未安装，请先安装 Docker"
        exit 1
    fi
    
    if ! command -v docker-compose &> /dev/null; then
        print_error "Docker Compose 未安装，请先安装 Docker Compose"
        exit 1
    fi
}

# 检查环境变量文件
check_env_file() {
    if [ ! -f ".env" ]; then
        print_warning ".env 文件不存在，正在创建..."
        if [ -f ".env.example" ]; then
            cp .env.example .env
            print_message "已从 .env.example 创建 .env 文件，请编辑后重新运行"
            exit 1
        else
            print_error ".env.example 文件不存在"
            exit 1
        fi
    fi
}

# 构建镜像
build_image() {
    print_message "开始构建 Docker 镜像..."
    docker build -t record-api:latest .
    print_message "镜像构建完成"
}

# 运行开发环境
run_dev() {
    print_message "启动开发环境..."
    docker-compose up --build -d
    print_message "开发环境启动完成"
    print_message "访问地址: http://localhost:8080"
    print_message "健康检查: http://localhost:8080/health"
}

# 运行生产环境
run_prod() {
    print_message "启动生产环境..."
    docker-compose -f docker-compose.prod.yml up --build -d
    print_message "生产环境启动完成"
}

# 停止服务
stop_services() {
    print_message "停止服务..."
    docker-compose down
    docker-compose -f docker-compose.prod.yml down 2>/dev/null || true
    print_message "服务已停止"
}

# 查看日志
show_logs() {
    docker-compose logs -f record-api
}

# 清理
cleanup() {
    print_message "清理 Docker 资源..."
    docker-compose down --volumes --remove-orphans
    docker system prune -f
    print_message "清理完成"
}

# 显示帮助信息
show_help() {
    echo "个人记账应用 Docker 构建脚本"
    echo ""
    echo "用法: $0 [选项]"
    echo ""
    echo "选项:"
    echo "  build     构建 Docker 镜像"
    echo "  dev       启动开发环境"
    echo "  prod      启动生产环境"
    echo "  stop      停止所有服务"
    echo "  logs      查看应用日志"
    echo "  cleanup   清理 Docker 资源"
    echo "  help      显示此帮助信息"
    echo ""
    echo "示例:"
    echo "  $0 build     # 构建镜像"
    echo "  $0 dev       # 启动开发环境"
    echo "  $0 logs      # 查看日志"
    echo "  $0 stop      # 停止服务"
}

# 主函数
main() {
    case "${1:-help}" in
        "build")
            check_docker
            build_image
            ;;
        "dev")
            check_docker
            check_env_file
            run_dev
            ;;
        "prod")
            check_docker
            check_env_file
            run_prod
            ;;
        "stop")
            check_docker
            stop_services
            ;;
        "logs")
            check_docker
            show_logs
            ;;
        "cleanup")
            check_docker
            cleanup
            ;;
        "help"|*)
            show_help
            ;;
    esac
}

# 执行主函数
main "$@"
