package logger

import (
	"go.uber.org/zap/zaptest/observer"
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Zaplog struct {
	*zap.Logger
}

var (
	instance *Zaplog
	once     sync.Once
	logs     *observer.ObservedLogs
	obs      zapcore.Core
	log      *zap.Logger
)

// initLogger initialise Logger instance only once
func initLogger() {
	once.Do(func() {
		config := zap.NewProductionEncoderConfig()
		config.EncodeTime = zapcore.ISO8601TimeEncoder
		consoleEncoder := zapcore.NewConsoleEncoder(config)

		defaultLogLevel := zapcore.InfoLevel
		if len(os.Getenv("DEV")) > 0 {
			defaultLogLevel = zapcore.DebugLevel
		}

		obs, logs = observer.New(defaultLogLevel)

		core := zapcore.NewTee(
			zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
			obs,
		)

		log = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
		instance = &Zaplog{log}
	})
}

// Log is invoking Zap Logger function
func Log() *Zaplog {
	initLogger()
	return instance
}

// InfoErr is invoking Zap Logger function with error message and fields and log level Info
func (z *Zaplog) InfoErr(e error, fields ...zapcore.Field) *Zaplog {
	z.Logger.Info(e.Error(), fields...)
	return z
}

// DebugErr is invoking Zap Logger function with error message and fields and log level Debug
func (z *Zaplog) DebugErr(e error, fields ...zapcore.Field) *Zaplog {
	z.Logger.Debug(e.Error(), fields...)
	return z
}

// WarnErr is invoking Zap Logger function with error message and fields and log level Warn
func (z *Zaplog) WarnErr(e error, fields ...zapcore.Field) *Zaplog {
	z.Logger.Warn(e.Error(), fields...)
	return z
}

// ErrorErr is invoking Zap Logger function with error message and fields and log level Error
func (z *Zaplog) ErrorErr(e error, fields ...zapcore.Field) *Zaplog {
	z.Logger.Error(e.Error(), fields...)
	return z
}

// PanicErr is invoking Zap Logger function with error message and fields and log level Panic
func (z *Zaplog) PanicErr(e error, fields ...zapcore.Field) *Zaplog {
	z.Logger.Panic(e.Error(), fields...)
	return z
}

// GetObservedLogs returns observed logs for testing purposes
func GetObservedLogs() *observer.ObservedLogs {
	initLogger()
	return logs
}

// ErrorLevel returns zapcore error level
func ErrorLevel() zapcore.Level {
	return zapcore.ErrorLevel
}
