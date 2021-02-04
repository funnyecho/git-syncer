package log

import (
	"os"

	"github.com/funnyecho/git-syncer/variable"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.SugaredLogger

func init() {
	errLevelEncoderConfig := zap.NewProductionEncoderConfig()
	errLevelEncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	errLevelEncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	core := zapcore.NewTee(
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(errLevelEncoderConfig),
			zapcore.Lock(os.Stderr),
			zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
				return lvl >= zapcore.ErrorLevel
			}),
		),
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(errLevelEncoderConfig),
			zapcore.Lock(os.Stdout),
			zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
				return lvl < zapcore.ErrorLevel
			}),
		),
	)

	logger = zap.New(core).Sugar()
	defer logger.Sync()
}

func Error(keyvals ...interface{}) {
	if !variable.UseVerboseError() {
		return
	}

	logger.Error(keyvals...)
}

func Info(keyvals ...interface{}) {
	if !variable.UseVerboseInfo() {
		return
	}

	logger.Info(keyvals...)
}

func Debug(keyvals ...interface{}) {
	if !variable.UseVerboseDebug() {
		return
	}

	logger.Debug(keyvals...)
}

func Errorw(msg string, keyvals ...interface{}) {
	if !variable.UseVerboseError() {
		return
	}

	logger.Errorw(msg, keyvals...)
}

func Infow(msg string, keyvals ...interface{}) {
	if !variable.UseVerboseInfo() {
		return
	}

	logger.Infow(msg, keyvals...)
}

func Debugw(msg string, keyvals ...interface{}) {
	if !variable.UseVerboseDebug() {
		return
	}

	logger.Debugw(msg, keyvals...)
}

func Errore(msg string, err error, keyvals ...interface{}) {
	if !variable.UseVerboseError() {
		return
	}

	logger.Errorw(msg, append(keyvals, "err", err)...)
}
