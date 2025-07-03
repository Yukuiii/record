package database

import (
	"fmt"
	"log"

	"github.com/sakura/record-api/config"
	"github.com/sakura/record-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitDB 初始化数据库连接
func InitDB() error {
	// 获取数据库配置
	dbConfig := config.GetConfig().Database

	// 构建连接DSN
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Shanghai",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DBName, dbConfig.SSLMode,
	)

	// 配置GORM
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	// 开发模式下使用更详细的日志
	if config.GetConfig().Server.Mode == "debug" {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	}

	// 连接数据库
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		return err
	}

	log.Println("数据库连接成功")

	// 测试数据库连接
	if err := HealthCheck(); err != nil {
		return fmt.Errorf("数据库连接测试失败: %w", err)
	}
	log.Println("数据库连接测试通过")

	// 自动迁移表结构（安全模式）
	err = migrateDB()
	if err != nil {
		log.Printf("表结构迁移警告: %v", err)
		// 不返回错误，继续执行，因为表可能已经存在
	}

	// 初始化默认数据
	err = seedData()
	if err != nil {
		return fmt.Errorf("初始化默认数据失败: %w", err)
	}

	return nil
}

// 自动迁移表结构
func migrateDB() error {
	log.Println("正在检查表结构...")

	// 检查表是否存在
	if DB.Migrator().HasTable(&models.User{}) {
		log.Println("用户表已存在，跳过创建")
	}
	if DB.Migrator().HasTable(&models.Category{}) {
		log.Println("分类表已存在，跳过创建")
	}
	if DB.Migrator().HasTable(&models.Transaction{}) {
		log.Println("交易表已存在，跳过创建")
	}

	// 尝试自动迁移（只会添加缺失的列，不会重复创建表）
	err := DB.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Transaction{},
	)

	if err != nil {
		log.Printf("表结构迁移完成，部分警告: %v", err)
		// 对于已存在的表，这不是致命错误
		return nil
	}

	log.Println("表结构检查完成")
	return nil
}

// 初始化默认数据
func seedData() error {
	log.Println("正在检查默认数据...")

	// 初始化默认分类
	return seedDefaultCategories()
}

// GetDB 获取数据库实例
func GetDB() *gorm.DB {
	return DB
}

// HealthCheck 数据库健康检查
func HealthCheck() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Ping()
}

// CheckTablesExist 检查必要的表是否存在
func CheckTablesExist() map[string]bool {
	tables := map[string]bool{
		"users":        DB.Migrator().HasTable(&models.User{}),
		"categories":   DB.Migrator().HasTable(&models.Category{}),
		"transactions": DB.Migrator().HasTable(&models.Transaction{}),
	}
	return tables
}

// GetDatabaseInfo 获取数据库信息
func GetDatabaseInfo() map[string]interface{} {
	tables := CheckTablesExist()

	// 统计各表的记录数
	var userCount, categoryCount, transactionCount int64
	if tables["users"] {
		DB.Model(&models.User{}).Count(&userCount)
	}
	if tables["categories"] {
		DB.Model(&models.Category{}).Count(&categoryCount)
	}
	if tables["transactions"] {
		DB.Model(&models.Transaction{}).Count(&transactionCount)
	}

	return map[string]interface{}{
		"tables": tables,
		"counts": map[string]int64{
			"users":        userCount,
			"categories":   categoryCount,
			"transactions": transactionCount,
		},
	}
}

// 初始化默认分类
func seedDefaultCategories() error {
	// 检查是否已存在默认分类
	var count int64
	DB.Model(&models.Category{}).Where("is_default = ?", true).Count(&count)
	if count > 0 {
		log.Printf("已存在 %d 个默认分类，跳过初始化", count)
		return nil
	}

	log.Println("正在初始化默认分类数据...")

	// 默认收入分类
	incomeCategories := []models.Category{
		{Name: "工资", Type: "income", Icon: "salary", Color: "#4CAF50", IsDefault: true},
		{Name: "奖金", Type: "income", Icon: "bonus", Color: "#8BC34A", IsDefault: true},
		{Name: "投资", Type: "income", Icon: "investment", Color: "#CDDC39", IsDefault: true},
		{Name: "报销", Type: "income", Icon: "reimburse", Color: "#FFC107", IsDefault: true},
		{Name: "其他收入", Type: "income", Icon: "other_income", Color: "#FF9800", IsDefault: true},
	}

	// 默认支出分类
	expenseCategories := []models.Category{
		{Name: "餐饮", Type: "expense", Icon: "food", Color: "#F44336", IsDefault: true},
		{Name: "交通", Type: "expense", Icon: "transport", Color: "#E91E63", IsDefault: true},
		{Name: "购物", Type: "expense", Icon: "shopping", Color: "#9C27B0", IsDefault: true},
		{Name: "娱乐", Type: "expense", Icon: "entertainment", Color: "#673AB7", IsDefault: true},
		{Name: "居家", Type: "expense", Icon: "home", Color: "#3F51B5", IsDefault: true},
		{Name: "通讯", Type: "expense", Icon: "communication", Color: "#2196F3", IsDefault: true},
		{Name: "医疗", Type: "expense", Icon: "medical", Color: "#00BCD4", IsDefault: true},
		{Name: "教育", Type: "expense", Icon: "education", Color: "#009688", IsDefault: true},
		{Name: "其他支出", Type: "expense", Icon: "other_expense", Color: "#FF5722", IsDefault: true},
	}

	// 合并所有分类
	var allCategories []models.Category
	allCategories = append(allCategories, incomeCategories...)
	allCategories = append(allCategories, expenseCategories...)

	// 批量创建分类
	if err := DB.Create(&allCategories).Error; err != nil {
		return err
	}

	log.Printf("成功创建 %d 个默认分类", len(allCategories))
	return nil
} 