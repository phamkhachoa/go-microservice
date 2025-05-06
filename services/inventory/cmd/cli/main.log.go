package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func main() {
	//sugar := zap.NewExample().Sugar()
	//
	//sugar.Infof("Hello %s, age: %d", "TipsGo", 40)
	//
	//// logger
	//logger := zap.NewExample()
	//logger.Info("Hello World", zap.String("name", "TipsGo"), zap.Int("age", 40))

	//logger := zap.NewExample()
	//logger.Info("Hello")
	//
	//// development
	//logger, _ = zap.NewDevelopment()
	//logger.Info("Hello development")
	//
	//// production
	//logger, _ = zap.NewProduction()
	//logger.Info("Hello production")

	// 3.
	encoder := getEncoderLog()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCaller())

	logger.Info("Info log", zap.Int("line", 1))
	logger.Error("Error log", zap.Int("line", 2))
}

// format log
func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()

	// 1745342189.31297 -> 2025-04-23T00:16:29.311+0700
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	encodeConfig.TimeKey = "time"
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encodeConfig)
}

func getWriterSync() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncFile, syncConsole)
}
