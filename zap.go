package go_common_aws_sqs

import (
	"encoding/json"
	zinc "github.com/yxw21/go-commons-zinc"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

type ZincWriter struct {
	zinc.Client
}

func (zincWriter *ZincWriter) Write(p []byte) (n int, err error) {
	var document = make(map[string]interface{})
	if err := json.Unmarshal(p, &document); err != nil {
		return 0, err
	}
	_, _, err = zincWriter.DocumentIndex(document)
	if err != nil {
		return 0, err
	}
	return len(p), nil
}

func NewZincWriter(index string) *ZincWriter {
	endpoint := os.Getenv("ZINC_ENDPOINT")
	username := os.Getenv("ZINC_USERNAME")
	password := os.Getenv("ZINC_PASSWORD")
	instance := &ZincWriter{}
	instance.SetEndpoint(endpoint)
	instance.SetAuth(username, password)
	instance.SetIndex(index)
	return instance
}

func NewLoggerWithZinc(zincWriter *ZincWriter) *zap.Logger {
	lowPriority := zap.LevelEnablerFunc(func(lv zapcore.Level) bool {
		return lv >= zapcore.DebugLevel
	})
	jsonEnc := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	stdCore := zapcore.NewCore(jsonEnc, zapcore.Lock(os.Stdout), lowPriority)
	syncer := zapcore.AddSync(zincWriter)
	redisCore := zapcore.NewCore(jsonEnc, syncer, lowPriority)
	core := zapcore.NewTee(stdCore, redisCore)
	return zap.New(core).WithOptions(
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.WarnLevel),
		zap.Fields(
			zap.String("time", time.Now().Format("2006-01-02 15:04:05")),
		),
	)
}

func NewLogger() *zap.Logger {
	lowPriority := zap.LevelEnablerFunc(func(lv zapcore.Level) bool {
		return lv >= zapcore.DebugLevel
	})
	jsonEnc := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	stdCore := zapcore.NewCore(jsonEnc, zapcore.Lock(os.Stdout), lowPriority)
	core := zapcore.NewTee(stdCore)
	return zap.New(core).WithOptions(
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.WarnLevel),
		zap.Fields(
			zap.String("time", time.Now().Format("2006-01-02 15:04:05")),
		),
	)
}
