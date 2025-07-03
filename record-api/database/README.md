# 数据库初始化文件

本目录包含了个人记账应用的数据库初始化SQL脚本，适用于Supabase PostgreSQL数据库。

## 文件说明

### 1. `setup.sql` - 完整初始化脚本（推荐）
包含完整的数据库初始化，包括：
- 表结构创建
- 索引创建
- 触发器创建
- 默认分类数据插入

### 2. `init.sql` - 表结构初始化
仅包含表结构、索引和触发器的创建，不包含默认数据。

### 3. `seed_categories.sql` - 默认分类数据
仅包含默认分类数据的插入脚本。

## 使用方法

### 方法一：在Supabase Dashboard中执行（推荐）

1. 登录 [Supabase Dashboard](https://app.supabase.com)
2. 选择您的项目
3. 进入 **SQL Editor**
4. 复制 `setup.sql` 的内容并粘贴到编辑器中
5. 点击 **Run** 执行脚本

### 方法二：使用psql命令行工具

```bash
# 连接到Supabase数据库
psql "postgresql://postgres.your-ref:your-password@aws-0-ap-southeast-1.pooler.supabase.com:6543/postgres"

# 执行初始化脚本
\i database/setup.sql
```

### 方法三：分步执行

如果需要分步执行，可以按以下顺序：

1. 先执行表结构初始化：
```sql
\i database/init.sql
```

2. 再执行默认数据插入：
```sql
\i database/seed_categories.sql
```

## 表结构说明

### users 表 - 用户信息
- `id`: 主键，自增
- `email`: 邮箱，唯一
- `phone`: 手机号，唯一
- `password`: 密码（加密存储）
- `nickname`: 昵称
- `avatar`: 头像URL
- `gender`: 性别
- `birthday`: 生日
- `register_time`: 注册时间
- `last_login`: 最后登录时间
- `status`: 状态（1:正常，0:禁用）

### categories 表 - 分类信息
- `id`: 主键，自增
- `name`: 分类名称
- `type`: 分类类型（income:收入，expense:支出）
- `icon`: 图标名称
- `color`: 颜色代码
- `is_default`: 是否为默认分类

### transactions 表 - 交易记录
- `id`: 主键，自增
- `user_id`: 用户ID（外键）
- `category_id`: 分类ID（外键）
- `amount`: 金额
- `type`: 交易类型（income:收入，expense:支出）
- `description`: 描述
- `record_time`: 记录时间
- `location`: 地点
- `image_url`: 图片URL
- `tags`: 标签

## 默认分类

### 收入分类
- 工资、奖金、投资、报销、其他收入

### 支出分类
- 餐饮、交通、购物、娱乐、居家、通讯、医疗、教育、其他支出

## 注意事项

1. 执行脚本前请确保已连接到正确的Supabase数据库
2. 脚本使用了 `IF NOT EXISTS` 和 `ON CONFLICT DO NOTHING`，可以安全地重复执行
3. 所有表都包含了 `created_at` 和 `updated_at` 字段，并设置了自动更新触发器
4. 建议在生产环境使用前先在测试环境验证脚本
