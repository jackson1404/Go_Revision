package logger_configs

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.SugaredLogger
}

var logInstance *Logger

// InitializeLogger initializes the global logger for a given environment.
func InitializeLogger(env string) error {
	var encoderCfg zapcore.EncoderConfig
	var level zapcore.Level

	if env == "prod" {
		encoderCfg = zap.NewProductionEncoderConfig()
		encoderCfg.TimeKey = "ts"
		encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
		level = zapcore.InfoLevel
	} else {
		encoderCfg = zap.NewDevelopmentEncoderConfig()
		encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
		level = zapcore.DebugLevel
	}

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderCfg),
		zapcore.AddSync(os.Stdout),
		zap.NewAtomicLevelAt(level),
	)

	logger := zap.New(core)
	logInstance = &Logger{logger.Sugar()}

	logInstance.Infof("✅ Logger initialized successfully (%s mode)", env)
	return nil
}

func GetLogger() *Logger {
	if logInstance == nil {
		logInstance = &Logger{SugaredLogger: zap.NewNop().Sugar()}
	}
	return logInstance
}

// helper functions
func Info(args ...interface{})             { GetLogger().Info(args...) }
func Infof(t string, args ...interface{})  { GetLogger().Infof(t, args...) }
func Error(args ...interface{})            { GetLogger().Error(args...) }
func Errorf(t string, args ...interface{}) { GetLogger().Errorf(t, args...) }
func Errorw(msg string, kv ...interface{}) { GetLogger().Errorw(msg, kv...) }
func Debugf(t string, args ...interface{}) { GetLogger().Debugf(t, args...) }
func Warnf(t string, args ...interface{})  { GetLogger().Warnf(t, args...) }

func Sync() {
	if logInstance != nil {
		_ = logInstance.Sync()
	}
}
