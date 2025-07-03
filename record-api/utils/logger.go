package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/sakura/record-api/config"
)

// Logger 日志记录器
type Logger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
	debugLogger *log.Logger
}

var logger *Logger

// InitLogger 初始化日志记录器
func InitLogger() error {
	// 确保日志目录存在
	logDir := "./logs"
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return err
	}

	// 创建日志文件
	today := time.Now().Format("2006-01-02")
	
	// 信息日志文件
	infoFile, err := os.OpenFile(
		filepath.Join(logDir, fmt.Sprintf("info_%s.log", today)),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666,
	)
	if err != nil {
		return err
	}

	// 错误日志文件
	errorFile, err := os.OpenFile(
		filepath.Join(logDir, fmt.Sprintf("error_%s.log", today)),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666,
	)
	if err != nil {
		return err
	}

	// 调试日志文件
	debugFile, err := os.OpenFile(
		filepath.Join(logDir, fmt.Sprintf("debug_%s.log", today)),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666,
	)
	if err != nil {
		return err
	}

	// 创建日志记录器
	logger = &Logger{
		infoLogger:  log.New(infoFile, "[INFO] ", log.LstdFlags|log.Lshortfile),
		errorLogger: log.New(errorFile, "[ERROR] ", log.LstdFlags|log.Lshortfile),
		debugLogger: log.New(debugFile, "[DEBUG] ", log.LstdFlags|log.Lshortfile),
	}

	return nil
}

// GetLogger 获取日志记录器实例
func GetLogger() *Logger {
	if logger == nil {
		// 如果日志记录器未初始化，使用默认的标准输出
		logger = &Logger{
			infoLogger:  log.New(os.Stdout, "[INFO] ", log.LstdFlags|log.Lshortfile),
			errorLogger: log.New(os.Stderr, "[ERROR] ", log.LstdFlags|log.Lshortfile),
			debugLogger: log.New(os.Stdout, "[DEBUG] ", log.LstdFlags|log.Lshortfile),
		}
	}
	return logger
}

// Info 记录信息日志
func (l *Logger) Info(v ...interface{}) {
	l.infoLogger.Println(v...)
}

// Infof 格式化记录信息日志
func (l *Logger) Infof(format string, v ...interface{}) {
	l.infoLogger.Printf(format, v...)
}

// Error 记录错误日志
func (l *Logger) Error(v ...interface{}) {
	l.errorLogger.Println(v...)
}

// Errorf 格式化记录错误日志
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.errorLogger.Printf(format, v...)
}

// Debug 记录调试日志
func (l *Logger) Debug(v ...interface{}) {
	// 只在调试模式下记录调试日志
	if config.GetConfig().Server.Mode == "debug" {
		l.debugLogger.Println(v...)
	}
}

// Debugf 格式化记录调试日志
func (l *Logger) Debugf(format string, v ...interface{}) {
	// 只在调试模式下记录调试日志
	if config.GetConfig().Server.Mode == "debug" {
		l.debugLogger.Printf(format, v...)
	}
}

// 全局日志函数

// LogInfo 全局信息日志
func LogInfo(v ...interface{}) {
	GetLogger().Info(v...)
}

// LogInfof 全局格式化信息日志
func LogInfof(format string, v ...interface{}) {
	GetLogger().Infof(format, v...)
}

// LogError 全局错误日志
func LogError(v ...interface{}) {
	GetLogger().Error(v...)
}

// LogErrorf 全局格式化错误日志
func LogErrorf(format string, v ...interface{}) {
	GetLogger().Errorf(format, v...)
}

// LogDebug 全局调试日志
func LogDebug(v ...interface{}) {
	GetLogger().Debug(v...)
}

// LogDebugf 全局格式化调试日志
func LogDebugf(format string, v ...interface{}) {
	GetLogger().Debugf(format, v...)
}
