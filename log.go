package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type MultiLogger map[string]*Logger

type Logger struct {
	zapLog *zap.Logger
}

var logger *Logger

func init() {
	core := zapcore.NewCore(zapcore.NewJSONEncoder(DefaultEncoderConfig), os.Stdout, zapcore.DebugLevel)
	logger.zapLog = zap.New(core).WithOptions(zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zap.DPanicLevel))
}

func New(opts ...Option) *Logger {
	c := Default()
	for _, opt := range opts {
		opt(c)
	}
	logger.zapLog = zap.New(c.Core())
	return logger
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
