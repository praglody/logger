package logger

import (
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func Init(logfile string) error {
	// 创建目录（如果不存在）
	dir := filepath.Dir(logfile)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	oFile, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	rsiWriter := zapcore.AddSync(oFile)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	logger = zap.New(zapcore.NewCore(encoder, rsiWriter, zapcore.DebugLevel), zap.AddCaller(), zap.AddCallerSkip(1))
	return nil
}

func GetLogger(logfile string) (newLogger *zap.Logger, err error) {
	// 创建目录（如果不存在）
	dir := filepath.Dir(logfile)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return nil, err
	}

	oFile, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	rsiWriter := zapcore.AddSync(oFile)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	newLogger = zap.New(zapcore.NewCore(encoder, rsiWriter, zapcore.DebugLevel), zap.AddCaller())
	return
}

func Debug(args ...interface{}) {
	logger.Sugar().Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	logger.Sugar().Debugf(template, args...)
}

func Info(args ...interface{}) {
	logger.Sugar().Info(args...)
}

func Infof(template string, args ...interface{}) {
	logger.Sugar().Infof(template, args...)
}

func Warn(args ...interface{}) {
	logger.Sugar().Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	logger.Sugar().Warnf(template, args...)
}

func Error(args ...interface{}) {
	logger.Sugar().Error(args...)
}

func Errof(template string, args ...interface{}) {
	logger.Sugar().Errorf(template, args...)
}

func Panic(args ...interface{}) {
	logger.Sugar().Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	logger.Sugar().Panicf(template, args...)
}

func Fatal(args ...interface{}) {
	logger.Sugar().Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	logger.Sugar().Fatalf(template, args...)
}

func Sync() {
	logger.Sugar().Sync()
}
