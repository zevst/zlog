```shell script
go get -u github.com/zevst/zlog
```

By default, log displays messages in **STDOUT**.
You can override current logger using the config.
###### Example
```go
var config *zlog.Setting // OR zlog.MultiLogger
_ = viper.ReadInConfig()
_ = viper.Unmarshal(&config)
zlog.Start(config.Core())
defer zlog.End()
```
