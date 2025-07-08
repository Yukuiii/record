-- 创建用户表
CREATE TABLE IF NOT EXISTS "users" (
    -- 主键ID（自增）
    id BIGSERIAL PRIMARY KEY,
    
    -- 业务UUID（唯一标识）
    user_id UUID NOT NULL DEFAULT gen_random_uuid(),
    
    -- 用户基本信息
    user_name VARCHAR(50) NOT NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE,
    phone VARCHAR(20) UNIQUE,
    avatar VARCHAR(500),
    remark TEXT,
    
    -- 用户状态（1:正常 0:禁用 -1:删除）
    status INTEGER NOT NULL DEFAULT 1,
    
    -- 审计字段（继承自BaseEntity）
    create_time BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()) * 1000,
    update_time BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()) * 1000,
    create_by BIGINT,
    update_by BIGINT,
    
    -- 创建时间戳（PostgreSQL标准）
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- 创建索引
CREATE UNIQUE INDEX IF NOT EXISTS idx_user_user_id ON "users"(user_id);
CREATE UNIQUE INDEX IF NOT EXISTS idx_user_email ON "users"(email) WHERE email IS NOT NULL;
CREATE UNIQUE INDEX IF NOT EXISTS idx_user_phone ON "users"(phone) WHERE phone IS NOT NULL;
CREATE INDEX IF NOT EXISTS idx_user_status ON "users"(status);
CREATE INDEX IF NOT EXISTS idx_user_create_time ON "users"(create_time);

-- 创建更新时间触发器
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    NEW.update_time = EXTRACT(EPOCH FROM NOW()) * 1000;
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_user_updated_at 
    BEFORE UPDATE ON "users" 
    FOR EACH ROW 
    EXECUTE FUNCTION update_updated_at_column();

-- 添加表注释
COMMENT ON TABLE "users" IS '用户表';
COMMENT ON COLUMN "users".id IS '主键ID';
COMMENT ON COLUMN "users".user_id IS '用户UUID';
COMMENT ON COLUMN "users".user_name IS '用户名';
COMMENT ON COLUMN "users".password IS '密码';
COMMENT ON COLUMN "users".email IS '邮箱';
COMMENT ON COLUMN "users".phone IS '手机号';
COMMENT ON COLUMN "users".avatar IS '头像URL';
COMMENT ON COLUMN "users".remark IS '备注';
COMMENT ON COLUMN "users".status IS '状态：1正常 0禁用 -1删除';
COMMENT ON COLUMN "users".create_time IS '创建时间（毫秒时间戳）';
COMMENT ON COLUMN "users".update_time IS '更新时间（毫秒时间戳）';
COMMENT ON COLUMN "users".create_by IS '创建人ID';
COMMENT ON COLUMN "users".update_by IS '更新人ID';


-- 创建交易记录表
CREATE TABLE IF NOT EXISTS transaction (
    -- 主键ID，自增
    id BIGSERIAL PRIMARY KEY,
    
    -- 用户ID，关联用户表
    user_id BIGINT NOT NULL,
    
    -- 交易类型：income-收入，expense-支出
    type VARCHAR(20) NOT NULL CHECK (type IN ('income', 'expense')),
    
    -- 交易分类ID，关联分类表
    category_id BIGINT,
    
    -- 交易金额，使用DECIMAL确保精度
    amount DECIMAL(15,2) NOT NULL CHECK (amount > 0),
    
    -- 交易描述信息
    description TEXT,
    
    -- 交易日期时间戳
    transaction_date BIGINT NOT NULL,
    
    -- 支付方式
    payment_method VARCHAR(50),
    
    -- 交易状态
    status VARCHAR(20) DEFAULT 'completed' CHECK (status IN ('pending', 'completed', 'cancelled')),
    
    -- 交易地点
    location VARCHAR(255),
    
    -- 货币类型
    currency VARCHAR(10) DEFAULT 'CNY',
    
    -- 交易标签，多个标签用逗号分隔
    tags TEXT,
    
    -- 交易备注信息
    remark TEXT,
    
    -- 继承自BaseEntity的字段
    create_time BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()) * 1000,
    update_time BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()) * 1000,
    create_by BIGINT,
    update_by BIGINT
);

-- 创建索引优化查询性能
CREATE INDEX idx_transaction_user_id ON transaction(user_id);
CREATE INDEX idx_transaction_type ON transaction(type);
CREATE INDEX idx_transaction_category_id ON transaction(category_id);
CREATE INDEX idx_transaction_date ON transaction(transaction_date);
CREATE INDEX idx_transaction_status ON transaction(status);
CREATE INDEX idx_transaction_create_time ON transaction(create_time);

-- 创建复合索引
CREATE INDEX idx_transaction_user_type_date ON transaction(user_id, type, transaction_date DESC);

-- 添加外键约束（如果有用户表的话）
-- ALTER TABLE transaction ADD CONSTRAINT fk_transaction_user 
--     FOREIGN KEY (user_id) REFERENCES "users"(id) ON DELETE CASCADE;

-- 创建更新时间自动更新的触发器
CREATE OR REPLACE FUNCTION update_transaction_updated_time()
RETURNS TRIGGER AS $$
BEGIN
    NEW.update_time = EXTRACT(EPOCH FROM NOW()) * 1000;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_transaction_time
    BEFORE UPDATE ON transaction
    FOR EACH ROW
    EXECUTE FUNCTION update_transaction_updated_time();

-- 添加表注释
COMMENT ON TABLE transaction IS '交易记录表';
COMMENT ON COLUMN transaction.id IS '交易记录主键ID';
COMMENT ON COLUMN transaction.user_id IS '用户ID，关联用户表';
COMMENT ON COLUMN transaction.type IS '交易类型：income-收入，expense-支出';
COMMENT ON COLUMN transaction.category_id IS '交易分类ID，关联分类表';
COMMENT ON COLUMN transaction.amount IS '交易金额';
COMMENT ON COLUMN transaction.description IS '交易描述信息';
COMMENT ON COLUMN transaction.transaction_date IS '交易日期时间戳';
COMMENT ON COLUMN transaction.payment_method IS '支付方式：现金、银行卡、支付宝、微信等';
COMMENT ON COLUMN transaction.status IS '交易状态：pending-待处理，completed-已完成，cancelled-已取消';
COMMENT ON COLUMN transaction.location IS '交易地点';
COMMENT ON COLUMN transaction.currency IS '货币类型：CNY、USD、EUR等';
COMMENT ON COLUMN transaction.tags IS '交易标签，多个标签用逗号分隔';
COMMENT ON COLUMN transaction.remark IS '交易备注信息';
