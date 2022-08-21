package utils

import (
	"encoding/json"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func InitLogger() {
	var logFile = "/var/logs/goapi/goapi.log"
	rotator, err := rotatelogs.New(
		logFile,
		rotatelogs.WithRotationTime(time.Second*5))
	if err != nil {
		panic(err)
	}

	encoderConfig := map[string]string{
		"levelEncoder": "capital",
		"timeKey":      "date",
		"timeEncoder":  "iso8601",
	}
	data, _ := json.Marshal(encoderConfig)
	var encCfg zapcore.EncoderConfig
	if err := json.Unmarshal(data, &encCfg); err != nil {
		panic(err)
	}

	w := zapcore.AddSync(rotator)
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encCfg),
		w,
		zap.InfoLevel)
	Logger = zap.New(core)
}
