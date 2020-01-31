## Zlog is a wrapper for [uber-go/zap](https://godoc.org/go.uber.org/zap)

##### Install
```shell script
go get -u github.com/zevst/zlog
```

By default, log displays messages in **STDOUT**.
You can override current logger using the config.

###### Example
```yaml
# example config for zlog.MultiLogger
stdout:
  level: -1
  format: -1
  color: true
kibana:
  level: 1
  out:
    filename: filepath.log
```
```go
var config *zlog.Setting // OR zlog.MultiLogger
_ = viper.ReadInConfig()
_ = viper.Unmarshal(&config)
zlog.Start(config.Core())
defer zlog.End()
```

**Formats**
```
-1          Console
 1  Default Json
```

**Levels**
```
-1          DebugLevel
 1  Default InfoLevel
 2          WarnLevel
 3          ErrorLevel
 4          DPanicLevel
 5          PanicLevel
 6          FatalLevel
```