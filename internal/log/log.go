package log

import (
	"context"
	"fmt"
	"github.com/funnyecho/git-syncer/internal/scopex"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
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

func Error(ctx context.Context, keyvals ...interface{})  {
	verbose := scopex.UseVerbose(ctx)
	if verbose < scopex.VerboseError {
		return
	}

	logger.Error(keyvals...)
}

func Info(ctx context.Context, keyvals ...interface{}) {
	verbose := scopex.UseVerbose(ctx)
	if verbose < scopex.VerboseInfo {
		return
	}

	logger.Info(keyvals...)
}

func Debug(ctx context.Context, keyvals ...interface{}) {
	verbose := scopex.UseVerbose(ctx)
	if verbose < scopex.VerboseDebug {
		return
	}

	logger.Debug(keyvals...)
}

func Errorw(ctx context.Context, msg string, keyvals ...interface{}) {
	verbose := scopex.UseVerbose(ctx)
	if verbose < scopex.VerboseError {
		return
	}

	logger.Errorw(msg, keyvals...)
}

func Infow(ctx context.Context, msg string, keyvals ...interface{}) {
	verbose := scopex.UseVerbose(ctx)
	if verbose < scopex.VerboseInfo {
		return
	}

	logger.Infow(msg, keyvals...)
}

func Debugw(ctx context.Context, msg string, keyvals ...interface{}) {
	verbose := scopex.UseVerbose(ctx)
	if verbose < scopex.VerboseDebug {
		return
	}

	logger.Debugw(msg, keyvals...)
}

func Errore(ctx context.Context, err error, keyvals ...interface{}) {
	verbose := scopex.UseVerbose(ctx)
	if verbose < scopex.VerboseError {
		return
	}

	logger.Errorw(fmt.Sprintf("%v", err), keyvals...)
}
