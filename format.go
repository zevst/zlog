package log

import "go.uber.org/zap/zapcore"

type Format int8
type fMapping map[Format]func(cfg zapcore.EncoderConfig) zapcore.Encoder

const (
	Console Format = iota - 1
	_
	Json
)

func (o Format) Encoder(cfg zapcore.EncoderConfig) zapcore.Encoder {
	return fMapping{
		Console: zapcore.NewConsoleEncoder,
		Json:    zapcore.NewJSONEncoder,
	}[o](cfg)
}
