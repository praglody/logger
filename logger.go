package logger

import (
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type LogSyncFunc func()

var logDir = "./logs"
var l *zap.Logger

func Init(logname string) error {

	logWriter, err := getWriter(logname)
	if err != nil {
		return err
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	l = zap.New(zapcore.NewCore(encoder, logWriter, zapcore.DebugLevel), zap.AddCaller(), zap.AddCallerSkip(1))
	return nil
}

func GetLogger(logname string) (newLogger *zap.Logger, callerSkip int, err error) {
	logWriter, err := getWriter(logname)
	if err != nil {
		return nil, err
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	newLogger = zap.New(zapcore.NewCore(encoder, logWriter, zapcore.DebugLevel), zap.AddCaller(), zap.AddCallerSkip(callerSkip))
	return
}

func getWriter(logname string) (zapcore.WriteSyncer, error) {
	// 创建目录（如果不存在）
	// dir := filepath.Dir(logDir)
	// println(dir)
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		return nil, err
	}

	logfile := logname + ".log"
	logpath := filepath.Join(logDir, logfile)
	// oFile, err := os.OpenFile(logpath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// if err != nil {
	// 	return err
	// }

	lumberJackLogger := &lumberjack.Logger{
		Filename:   logpath, // 文件位置
		MaxSize:    100,     // 进行切割之前,日志文件的最大大小(MB为单位)
		MaxAge:     90,      // 保留旧文件的最大天数
		MaxBackups: 1024,    // 保留旧文件的最大个数
		Compress:   false,   // 是否压缩/归档旧文件
	}

	return zapcore.AddSync(lumberJackLogger), nil
}

func SetLogDir(dir string) {
	logDir = dir
}

func Debug(args ...interface{}) {
	l.Sugar().Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	l.Sugar().Debugf(template, args...)
}

func Info(args ...interface{}) {
	l.Sugar().Info(args...)
}

func Infof(template string, args ...interface{}) {
	l.Sugar().Infof(template, args...)
}

func Warn(args ...interface{}) {
	l.Sugar().Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	l.Sugar().Warnf(template, args...)
}

func Error(args ...interface{}) {
	l.Sugar().Error(args...)
}

func Errorf(template string, args ...interface{}) {
	l.Sugar().Errorf(template, args...)
}

func Panic(args ...interface{}) {
	l.Sugar().Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	l.Sugar().Panicf(template, args...)
}

func Fatal(args ...interface{}) {
	l.Sugar().Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	l.Sugar().Fatalf(template, args...)
}

func Sync() {
	l.Sugar().Sync()
}
