package initialize

import (
	"go.uber.org/zap"
)

func InitLogger() {
	/*
		1. S()可以获取一个全局的suger，可以让我们设置一个全局的logger
		2. 日志是分级别的（debug info warn error fetal），Debug一旦过多容易占用内存，可以设置为Error，这样就只输出Error以上的日志
		3. S函数和L函数很有用
	*/
	/*
		return Config{
			Level:            NewAtomicLevelAt(DebugLevel),
			Development:      true,
			Encoding:         "console",
			EncoderConfig:    NewDevelopmentEncoderConfig(),
			OutputPaths:      []string{"stderr"},
			ErrorOutputPaths: []string{"stderr"},
		}
	*/
	loggerConfig := zap.NewDevelopmentConfig()
	loggerConfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	logger, _ := loggerConfig.Build()
	zap.ReplaceGlobals(logger)
}
