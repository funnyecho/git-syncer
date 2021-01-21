package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.SugaredLogger

func init() {
	errLevelEncoderConfig := zap.NewProductionEncoderConfig()
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
			zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig()),
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
	if verbose < VerboseError {
		return
	}

	logger.Error(keyvals...)
}

func Info(keyvals ...interface{}) {
	if verbose < VerboseInfo {
		return
	}

	logger.Info(keyvals...)
}

func Debug(keyvals ...interface{}) {
	if verbose < VerboseDebug {
		return
	}

	logger.Debug(keyvals...)
}

func Errorw(msg string, keyvals ...interface{}) {
	if verbose < VerboseError {
		return
	}

	logger.Errorw(msg, keyvals...)
}

func Infow(msg string, keyvals ...interface{}) {
	if verbose < VerboseInfo {
		return
	}

	logger.Infow(msg, keyvals...)
}

func Debugw(msg string, keyvals ...interface{}) {
	if verbose < VerboseDebug {
		return
	}

	logger.Debugw(msg, keyvals...)
}

func Errore(msg string, err error, keyvals ...interface{}) {
	if verbose < VerboseError {
		return
	}

	logger.Errorw(msg, append(keyvals, "err", err)...)
}
