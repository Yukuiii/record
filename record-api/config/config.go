package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

// Config 应用程序配置结构
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Port string
	Mode string
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// JWTConfig JWT配置
type JWTConfig struct {
	Secret     string
	ExpireTime int // 单位：小时
}

var config Config

// LoadConfig 加载配置文件
func LoadConfig() error {
	// 设置默认值
	setDefaults()

	// 设置配置文件信息
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	
	// 查找配置文件的路径
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	
	// 尝试创建默认配置文件
	if err := createDefaultConfigIfNotExists(); err != nil {
		return err
	}

	// 读取环境变量
	viper.AutomaticEnv()
	viper.SetEnvPrefix("RECORD")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// 绑定特定的环境变量
	viper.BindEnv("database.host", "DATABASE_HOST")
	viper.BindEnv("database.port", "DATABASE_PORT")
	viper.BindEnv("database.user", "DATABASE_USER")
	viper.BindEnv("database.password", "DATABASE_PASSWORD")
	viper.BindEnv("database.dbname", "DATABASE_NAME")
	viper.BindEnv("database.sslmode", "DATABASE_SSLMODE")

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	// 将配置映射到结构体
	if err := viper.Unmarshal(&config); err != nil {
		return err
	}

	return nil
}

// GetConfig 获取配置信息
func GetConfig() Config {
	return config
}

// 设置默认配置值
func setDefaults() {
	viper.SetDefault("server.port", "8080")
	viper.SetDefault("server.mode", "debug")
	
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
	viper.SetDefault("database.user", "postgres")
	viper.SetDefault("database.password", "postgres")
	viper.SetDefault("database.dbname", "postgres")
	viper.SetDefault("database.sslmode", "disable")
	
	viper.SetDefault("jwt.secret", "your-secret-key")
	viper.SetDefault("jwt.expiretime", 24) // 默认24小时
}

// 如果配置文件不存在，创建默认配置文件
func createDefaultConfigIfNotExists() error {
	// 检查配置文件是否存在
	if _, err := os.Stat("./config/config.yaml"); os.IsNotExist(err) {
		// 确保目录存在
		if err := os.MkdirAll("./config", 0755); err != nil {
			return err
		}
		
		// 创建默认配置文件
		f, err := os.Create(filepath.Join("./config", "config.yaml"))
		if err != nil {
			return err
		}
		defer f.Close()
		
		// 写入默认配置
		defaultConfig := `server:
  port: 8080
  mode: debug

database:
  host: localhost
  port: 5432
  user: postgres
  password: postgres
  dbname: postgres
  sslmode: disable

jwt:
  secret: your-secret-key
  expiretime: 24
`
		if _, err := f.WriteString(defaultConfig); err != nil {
			return err
		}
	}
	
	return nil
}
