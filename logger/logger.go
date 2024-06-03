package logger

import (
	"CallTheRoll/config"
	"errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
)

// LoggerInit 初始化日志
func LoggerInit() error {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // 或者自定义时间编码器
	encoder := zapcore.NewConsoleEncoder(encoderConfig)   // 使用带颜色的控制台编码器

	// 获取输出目标
	outputs, err := getWriter(config.G().Log.Output)
	if err != nil {
		return err
	}

	// 解析日志级别
	level, err := zapcore.ParseLevel(config.G().Log.Level)
	if err != nil {
		return errors.New("logf: " + err.Error())
	}

	// 构建核心logger
	core := zapcore.NewCore(encoder, outputs, zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l >= level
	}))

	// 构建logger实例
	logger := zap.New(core)

	// 替换 zap 库默认的全局logger
	zap.ReplaceGlobals(logger)

	return nil
}

func getWriter(ss []string) (zapcore.WriteSyncer, error) {
	var writers []zapcore.WriteSyncer
	for _, target := range ss {
		switch target {
		case "stdout":
			writers = append(writers, zapcore.Lock(os.Stdout))
		case "stderr":
			writers = append(writers, zapcore.Lock(os.Stderr))
		default:
			// 确保目录存在
			if err := os.MkdirAll(filepath.Dir(target), 0755); err != nil {
				return nil, errors.New("logf: " + err.Error())
			}
			file, err := os.OpenFile(target, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
			if err != nil {
				return nil, errors.New("logf: " + err.Error())
			}
			writers = append(writers, zapcore.AddSync(file))
		}
	}
	return zapcore.NewMultiWriteSyncer(writers...), nil
}
