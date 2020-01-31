package zlog

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type MultiLogger map[string]*Setting

func (m MultiLogger) Core() zapcore.Core {
	var cores []zapcore.Core
	for _, logger := range m {
		cores = append(cores, logger.Core())
	}
	return zapcore.NewTee(cores...)
}

func (s *Setting) Core() zapcore.Core {
	c := Default()
	WithSetting(s)(c)
	return c.Core()
}

type Logger struct {
	zapLog *zap.Logger
}

var logger = new(Logger)

func init() {
	core := zapcore.NewCore(zapcore.NewJSONEncoder(DefaultEncoderConfig), os.Stdout, zapcore.DebugLevel)
	logger.zapLog = zap.New(core).WithOptions(zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zap.DPanicLevel))
}

func Start(core zapcore.Core, opts ...zap.Option) {
	logger.zapLog = zap.New(core, opts...)
}

func Get() *Logger {
	return logger
}

func End() {
	_ = logger.zapLog.Sync()
}

func Debug(msg string, fields ...zap.Field)  { logger.zapLog.Debug(msg, fields...) }
func Info(msg string, fields ...zap.Field)   { logger.zapLog.Info(msg, fields...) }
func Warn(msg string, fields ...zap.Field)   { logger.zapLog.Warn(msg, fields...) }
func Error(msg string, fields ...zap.Field)  { logger.zapLog.Error(msg, fields...) }
func DPanic(msg string, fields ...zap.Field) { logger.zapLog.DPanic(msg, fields...) }
func Panic(msg string, fields ...zap.Field)  { logger.zapLog.Panic(msg, fields...) }
func Fatal(msg string, fields ...zap.Field)  { logger.zapLog.Fatal(msg, fields...) }

func (l *Logger) Error(msg string)                       { l.Printf("ERROR: %s", msg) }
func (l *Logger) Infof(msg string, args ...interface{})  { l.Printf(msg, args...) }
func (l *Logger) Print(v ...interface{})                 { l.zapLog.Debug(fmt.Sprint(v...)) }
func (l *Logger) Printf(format string, v ...interface{}) { l.zapLog.Debug(fmt.Sprintf(format, v...)) }
func (l *Logger) Println(v ...interface{})               { l.zapLog.Debug(fmt.Sprint(v...)) }
func (l *Logger) Write(p []byte) (n int, err error)      { l.zapLog.Debug(string(p)); return len(p), nil }

func (l *Logger) WithOptions(opts ...zap.Option) *Logger {
	l.zapLog.WithOptions(opts...)
	return l
}
