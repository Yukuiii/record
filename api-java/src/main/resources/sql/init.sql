-- 创建用户表
CREATE TABLE IF NOT EXISTS "user" (
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
CREATE UNIQUE INDEX IF NOT EXISTS idx_user_user_id ON "user"(user_id);
CREATE UNIQUE INDEX IF NOT EXISTS idx_user_email ON "user"(email) WHERE email IS NOT NULL;
CREATE UNIQUE INDEX IF NOT EXISTS idx_user_phone ON "user"(phone) WHERE phone IS NOT NULL;
CREATE INDEX IF NOT EXISTS idx_user_status ON "user"(status);
CREATE INDEX IF NOT EXISTS idx_user_create_time ON "user"(create_time);

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
    BEFORE UPDATE ON "user" 
    FOR EACH ROW 
    EXECUTE FUNCTION update_updated_at_column();

-- 添加表注释
COMMENT ON TABLE "user" IS '用户表';
COMMENT ON COLUMN "user".id IS '主键ID';
COMMENT ON COLUMN "user".user_id IS '用户UUID';
COMMENT ON COLUMN "user".user_name IS '用户名';
COMMENT ON COLUMN "user".password IS '密码';
COMMENT ON COLUMN "user".email IS '邮箱';
COMMENT ON COLUMN "user".phone IS '手机号';
COMMENT ON COLUMN "user".avatar IS '头像URL';
COMMENT ON COLUMN "user".remark IS '备注';
COMMENT ON COLUMN "user".status IS '状态：1正常 0禁用 -1删除';
COMMENT ON COLUMN "user".create_time IS '创建时间（毫秒时间戳）';
COMMENT ON COLUMN "user".update_time IS '更新时间（毫秒时间戳）';
COMMENT ON COLUMN "user".create_by IS '创建人ID';
COMMENT ON COLUMN "user".update_by IS '更新人ID';