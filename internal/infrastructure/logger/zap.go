package logger

import (
	"go-mini-admin/config"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger interface {
	Debug(args ...any)
	Info(args ...any)
	Warn(args ...any)
	Error(args ...any)
	Fatal(args ...any)
	Debugf(template string, args ...any)
	Infof(template string, args ...any)
	Warnf(template string, args ...any)
	Errorf(template string, args ...any)
	Fatalf(template string, args ...any)
}

type zapLogger struct {
	sugar *zap.SugaredLogger
}

func New(cfg *config.LogConfig) (Logger, error) {
	// 初始化日志目录文件
	logDir := filepath.Dir(cfg.Filename)
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return nil, err
	}

	// 日志等级
	var level zapcore.Level
	switch cfg.Level {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	case "fatal":
		level = zapcore.FatalLevel
	default:
		level = zapcore.InfoLevel
	}

	// encoder config
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	fileWriter := &lumberjack.Logger{
		Filename:   cfg.Filename,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
		Compress:   cfg.Compress,
	}

	// Multi-writer: file + console
	core := zapcore.NewTee(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			zapcore.AddSync(fileWriter),
			level,
		),
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			zapcore.AddSync(os.Stdout),
			level,
		),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	return &zapLogger{sugar: logger.Sugar()}, nil
}

func (l *zapLogger) Debug(args ...any) { l.sugar.Debug(args...) }
func (l *zapLogger) Info(args ...any)  { l.sugar.Info(args...) }
func (l *zapLogger) Warn(args ...any)  { l.sugar.Warn(args...) }
func (l *zapLogger) Error(args ...any) { l.sugar.Error(args...) }
func (l *zapLogger) Fatal(args ...any) { l.sugar.Fatal(args...) }

func (l *zapLogger) Debugf(template string, args ...any) { l.sugar.Debugf(template, args...) }
func (l *zapLogger) Infof(template string, args ...any)  { l.sugar.Infof(template, args...) }
func (l *zapLogger) Warnf(template string, args ...any)  { l.sugar.Warnf(template, args...) }
func (l *zapLogger) Errorf(template string, args ...any) { l.sugar.Errorf(template, args...) }
func (l *zapLogger) Fatalf(template string, args ...any) { l.sugar.Fatalf(template, args...) }

// Package-level convenience functions using the global logger
var defaultLogger Logger

func Init(cfg *config.LogConfig) error {
	var err error
	defaultLogger, err = New(cfg)
	return err
}
func Default() Logger {
	return defaultLogger
}
func Debug(args ...any) { defaultLogger.Debug(args...) }
func Info(args ...any)  { defaultLogger.Info(args...) }
func Warn(args ...any)  { defaultLogger.Warn(args...) }
func Error(args ...any) { defaultLogger.Error(args...) }
func Fatal(args ...any) { defaultLogger.Fatal(args...) }

func Debugf(template string, args ...any) { defaultLogger.Debugf(template, args...) }
func Infof(template string, args ...any)  { defaultLogger.Infof(template, args...) }
func Warnf(template string, args ...any)  { defaultLogger.Warnf(template, args...) }
func Errorf(template string, args ...any) { defaultLogger.Errorf(template, args...) }
func Fatalf(template string, args ...any) { defaultLogger.Fatalf(template, args...) }
