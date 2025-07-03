# 前端项目结构说明

## 目录结构

```
frontend/
├── assets/                   # 静态资源
│   ├── css/
│   │   └── main.css         # 全局样式文件
│   └── images/
│       └── .gitkeep         # 图片资源目录
├── components/               # Vue 组件
│   ├── ui/                  # 基础 UI 组件
│   │   ├── Button.vue       # 按钮组件
│   │   ├── Card.vue         # 卡片组件
│   │   └── Input.vue        # 输入框组件
│   ├── layout/              # 布局组件
│   │   └── Header.vue       # 头部导航组件
│   └── record/              # 记录相关组件
│       └── RecordForm.vue   # 记录表单组件
├── composables/              # 组合式函数
│   ├── useApi.js            # API 调用函数
│   ├── useAuth.js           # 用户认证函数
│   └── useRecords.js        # 记录管理函数
├── layouts/                  # 页面布局
│   ├── default.vue          # 默认布局
│   └── auth.vue             # 认证页面布局
├── middleware/               # 路由中间件
│   └── auth.js              # 认证中间件
├── pages/                    # 页面组件（路由）
│   ├── index.vue            # 首页
│   └── auth/
│       └── login.vue        # 登录页面
├── plugins/                  # 插件
│   └── api.client.js        # API 客户端插件
├── public/                   # 静态文件
│   ├── favicon.ico          # 网站图标
│   └── robots.txt           # 搜索引擎配置
├── server/                   # 服务端目录
│   └── tsconfig.json        # 服务端 TypeScript 配置
├── stores/                   # 状态管理
│   └── user.js              # 用户状态管理
├── types/                    # 类型定义
│   └── index.js             # 通用类型定义
├── app.vue                   # 主应用组件
├── nuxt.config.ts           # Nuxt.js 配置文件
├── package.json             # 项目配置和依赖
├── tsconfig.json            # TypeScript 配置
└── README.md                # 项目说明文档
```

## 技术栈

- **框架**: Nuxt.js 3.17.5
- **UI 库**: Vue.js 3.5.17
- **语言**: JavaScript (ES6+)
- **路由**: Vue Router 4.5.1
- **构建工具**: Vite
- **样式**: CSS3 + CSS Variables
