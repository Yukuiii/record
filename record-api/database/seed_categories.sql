-- 默认分类数据初始化脚本
-- 插入默认收入分类和支出分类

-- 清空现有分类数据（可选，如果需要重新初始化）
-- DELETE FROM categories WHERE is_default = true;

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
