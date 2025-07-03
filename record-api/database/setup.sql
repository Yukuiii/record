-- 个人记账应用完整数据库初始化脚本
-- 包含表结构、索引、触发器和默认数据

-- ================================
-- 1. 创建表结构
-- ================================

-- 创建用户表
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    phone VARCHAR(20) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    nickname VARCHAR(50) NOT NULL,
    avatar VARCHAR(255) DEFAULT '',
    gender VARCHAR(10) DEFAULT '',
    birthday TIMESTAMP DEFAULT NULL,
    register_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    last_login TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status INTEGER DEFAULT 1,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建分类表
CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    type VARCHAR(20) NOT NULL CHECK (type IN ('income', 'expense')),
    icon VARCHAR(50) NOT NULL,
    color VARCHAR(20) NOT NULL,
    is_default BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建交易记录表
CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    category_id INTEGER NOT NULL REFERENCES categories(id) ON DELETE RESTRICT,
    amount DECIMAL(15,2) NOT NULL CHECK (amount > 0),
    type VARCHAR(20) NOT NULL CHECK (type IN ('income', 'expense')),
    description VARCHAR(255) DEFAULT '',
    record_time TIMESTAMP NOT NULL,
    location VARCHAR(255) DEFAULT '',
    image_url VARCHAR(255) DEFAULT '',
    tags VARCHAR(255) DEFAULT '',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- ================================
-- 2. 创建索引
-- ================================

CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
CREATE INDEX IF NOT EXISTS idx_users_phone ON users(phone);
CREATE INDEX IF NOT EXISTS idx_transactions_user_id ON transactions(user_id);
CREATE INDEX IF NOT EXISTS idx_transactions_category_id ON transactions(category_id);
CREATE INDEX IF NOT EXISTS idx_transactions_type ON transactions(type);
CREATE INDEX IF NOT EXISTS idx_transactions_record_time ON transactions(record_time);
CREATE INDEX IF NOT EXISTS idx_transactions_user_time ON transactions(user_id, record_time);

-- ================================
-- 3. 创建触发器
-- ================================

-- 创建更新时间触发器函数
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

-- 为各表创建更新时间触发器
DROP TRIGGER IF EXISTS update_users_updated_at ON users;
CREATE TRIGGER update_users_updated_at
    BEFORE UPDATE ON users
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

DROP TRIGGER IF EXISTS update_categories_updated_at ON categories;
CREATE TRIGGER update_categories_updated_at
    BEFORE UPDATE ON categories
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

DROP TRIGGER IF EXISTS update_transactions_updated_at ON transactions;
CREATE TRIGGER update_transactions_updated_at
    BEFORE UPDATE ON transactions
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- ================================
-- 4. 插入默认分类数据
-- ================================

-- 插入默认收入分类
INSERT INTO categories (name, type, icon, color, is_default) VALUES
('工资', 'income', 'salary', '#4CAF50', true),
('奖金', 'income', 'bonus', '#8BC34A', true),
('投资', 'income', 'investment', '#CDDC39', true),
('报销', 'income', 'reimburse', '#FFC107', true),
('其他收入', 'income', 'other_income', '#FF9800', true)
ON CONFLICT DO NOTHING;

-- 插入默认支出分类
INSERT INTO categories (name, type, icon, color, is_default) VALUES
('餐饮', 'expense', 'food', '#F44336', true),
('交通', 'expense', 'transport', '#E91E63', true),
('购物', 'expense', 'shopping', '#9C27B0', true),
('娱乐', 'expense', 'entertainment', '#673AB7', true),
('居家', 'expense', 'home', '#3F51B5', true),
('通讯', 'expense', 'communication', '#2196F3', true),
('医疗', 'expense', 'medical', '#00BCD4', true),
('教育', 'expense', 'education', '#009688', true),
('其他支出', 'expense', 'other_expense', '#FF5722', true)
ON CONFLICT DO NOTHING;

-- ================================
-- 5. 验证数据
-- ================================

-- 查看创建的表
SELECT 'Tables created:' as status;
SELECT table_name FROM information_schema.tables 
WHERE table_schema = 'public' AND table_name IN ('users', 'categories', 'transactions');

-- 查看默认分类数据
SELECT 'Default categories:' as status;
SELECT COUNT(*) as category_count, type FROM categories WHERE is_default = true GROUP BY type;
