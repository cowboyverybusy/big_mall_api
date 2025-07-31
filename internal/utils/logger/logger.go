package logger

import (
	"big_mall_api/configs"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
	"time"
)

var GlobalLogger *logrus.Logger

// DefaultConfig 默认配置
func DefaultConfig() *configs.LogConfig {
	return &configs.LogConfig{
		Level:  "info",
		Format: "text",
		Output: "stdout",
	}
}

// Init 根据配置创建日志实例
func Init(cfg *configs.LogConfig) (*logrus.Logger, error) {
	if cfg == nil {
		cfg = DefaultConfig()
	}

	logger := logrus.New()

	// 设置日志级别
	level, err := logrus.ParseLevel(cfg.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	logger.SetLevel(level)

	// 设置日志格式
	switch cfg.Format {
	case "json":
		logger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
	default:
		logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
		})
	}

	// 设置输出位置
	var output io.Writer
	switch cfg.Output {
	case "stderr":
		output = os.Stderr
	case "file":
		if cfg.Dir == "" {
			return nil, &ConfigError{Message: "log directory is required when output is file"}
		}
		if cfg.FilePrefix == "" {
			return nil, &ConfigError{Message: "file prefix is required when output is file"}
		}

		// 创建带日期后缀的文件名
		now := time.Now()
		dateStr := now.Format("2006-01-02")

		// 组合成带日期的文件名: prefix_2006-01-02.log
		filename := fmt.Sprintf("%s_%s.log", cfg.FilePrefix, dateStr)
		fullPath := filepath.Join(cfg.Dir, filename)

		// 创建目录（如果不存在）
		if err := os.MkdirAll(cfg.Dir, 0755); err != nil {
			return nil, &ConfigError{Message: "failed to create log directory: " + err.Error()}
		}

		// 打开或创建日志文件
		file, err := os.OpenFile(fullPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return nil, &ConfigError{Message: "failed to open log file: " + err.Error()}
		}
		output = file
	default:
		output = os.Stdout
	}
	logger.SetOutput(output)

	GlobalLogger = logger
	return logger, nil
}

// WithFields 添加结构化字段
func WithFields(fields map[string]interface{}) *logrus.Entry {
	return GlobalLogger.WithFields(fields)
}

// Info 信息级别日志
func Info(args ...interface{}) {
	GlobalLogger.Info(args...)
}

// Infof 格式化信息级别日志
func Infof(format string, args ...interface{}) {
	GlobalLogger.Infof(format, args...)
}

// Warn 警告级别日志
func Warn(args ...interface{}) {
	GlobalLogger.Warn(args...)
}

// Warnf 格式化警告级别日志
func Warnf(format string, args ...interface{}) {
	GlobalLogger.Warnf(format, args...)
}

// Error 错误级别日志
func Error(args ...interface{}) {
	GlobalLogger.Error(args...)
}

// Errorf 格式化错误级别日志
func Errorf(format string, args ...interface{}) {
	GlobalLogger.Errorf(format, args...)
}

// Fatal 致命错误级别日志（会退出程序）
func Fatal(args ...interface{}) {
	GlobalLogger.Fatal(args...)
}

// Fatalf 格式化致命错误级别日志（会退出程序）
func Fatalf(format string, args ...interface{}) {
	GlobalLogger.Fatalf(format, args...)
}

// Panic 恐慌级别日志（会触发panic）
func Panic(args ...interface{}) {
	GlobalLogger.Panic(args...)
}

// Debug 调试级别日志
func Debug(args ...interface{}) {
	GlobalLogger.Debug(args...)
}

// Debugf 格式化调试级别日志
func Debugf(format string, args ...interface{}) {
	GlobalLogger.Debugf(format, args...)
}

// ConfigError 配置错误
type ConfigError struct {
	Message string
}

func (e *ConfigError) Error() string {
	return "logger config error: " + e.Message
}
