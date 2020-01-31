package zlog

import (
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

type Setting struct {
	Level  Level              `mapstructure:"level"`
	Format Format             `mapstructure:"format"`
	Color  bool               `mapstructure:"color"`
	Out    *lumberjack.Logger `mapstructure:"out"`
}

type config struct {
	format Format
	encCfg zapcore.EncoderConfig
	writer zapcore.WriteSyncer
	level  zapcore.LevelEnabler
}

var DefaultEncoderConfig = zapcore.EncoderConfig{
	MessageKey:    "M",
	LevelKey:      "L",
	TimeKey:       "T",
	NameKey:       "N",
	CallerKey:     "C",
	StacktraceKey: "S",
	LineEnding:    zapcore.DefaultLineEnding,

	EncodeLevel:    zapcore.CapitalLevelEncoder,
	EncodeTime:     zapcore.RFC3339TimeEncoder,
	EncodeDuration: zapcore.StringDurationEncoder,
	EncodeCaller:   zapcore.ShortCallerEncoder,

	EncodeName: zapcore.FullNameEncoder,
}

func Default() *config {
	return &config{format: Json, encCfg: DefaultEncoderConfig, writer: os.Stdout, level: zapcore.DebugLevel}
}

func (c *config) Core() zapcore.Core {
	return zapcore.NewCore(c.format.Encoder(c.encCfg), c.writer, c.level)
}

func (c *config) withSetting(setting *Setting) *config {
	if setting.Format != 0 {
		c.format = setting.Format
	}
	if setting.Level != 0 {
		c.level = setting.Level.Zap()
	}
	if setting.Out != nil {
		c.writer = zapcore.AddSync(setting.Out)
	}
	if setting.Color {
		c.encCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}
	return c
}
