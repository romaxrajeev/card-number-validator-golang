package pkg

import "go.uber.org/zap"

func Logger(logFilePath string, enableConsole bool) *zap.Logger {
	config := zap.NewProductionConfig()
	if !enableConsole {
		config.OutputPaths = []string{logFilePath}
	} else {
		config.OutputPaths = []string{"stdout", logFilePath}
	}
	logger, err := config.Build(zap.AddCaller())
	if err != nil {
		panic(err)
	}
	return logger
}

func DefaultLogger() *zap.Logger {
	return Logger("./logs/logs.log", true)
}
