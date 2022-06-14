package logger

import "go.uber.org/zap"

func NewLogger() *zap.SugaredLogger {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	sugar := logger.Sugar()

	defer func(logger *zap.SugaredLogger) {
		err = logger.Sync()
		if err != nil {
			sugar.Infof("%s", err)
		}
	}(sugar)

	return sugar
}
