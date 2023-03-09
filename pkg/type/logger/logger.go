package logger

import (
	"fmt"
	"os"
	"strings"

	"github.com/evgeniy-dammer/clean-architecture/pkg/type/context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type options struct {
	Level        zapcore.Level
	IsProduction bool
}

type Logger struct {
	logger *zap.Logger
}

var opt *options

func init() {
	opt = &options{IsProduction: false}

	if strings.ToLower(strings.TrimSpace(os.Getenv("IS_PRODUCTION"))) == "true" {
		opt.IsProduction = true
	}

	switch strings.ToUpper(strings.TrimSpace(os.Getenv("LOG_LEVEL"))) {
	case "ERR", "ERROR":
		opt.Level = zapcore.ErrorLevel
	case "WARN", "WARNING":
		opt.Level = zapcore.WarnLevel
	case "INFO":
		opt.Level = zapcore.InfoLevel
	case "DEBUG":
		opt.Level = zapcore.DebugLevel
	case "FATAL":
		opt.Level = zapcore.FatalLevel
	default:
		opt.Level = zapcore.InfoLevel
	}
}

func newLogger() (*Logger, error) {
	var config zap.Config

	if opt.IsProduction {
		config = zap.NewProductionConfig()
	} else {
		config = zap.NewDevelopmentConfig()
	}

	config.Level = zap.NewAtomicLevelAt(opt.Level)

	newLogger, err := config.Build(zap.AddCallerSkip(2))
	if err != nil {
		return nil, err
	}

	newLogger.Info("Set LOG_LEVEL", zap.Stringer("level", opt.Level))

	log = &Logger{logger: newLogger}

	return log, nil
}

func New() (*Logger, error) {
	return newLogger()
}

func (l *Logger) getContextFields(ctx context.Context) []zap.Field {
	return []zap.Field{zap.String("id", ctx.ID())}
}

func (l *Logger) Debug(msg string, fields ...zap.Field) {
	l.logger.Debug(msg, fields...)
}

func (l *Logger) DebugWithContext(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, l.getContextFields(ctx)...)
	l.Debug(msg, fields...)
}

func (l *Logger) Info(msg string, fields ...zap.Field) {
	l.logger.Info(msg, fields...)
}

func (l *Logger) InfoWithContext(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, l.getContextFields(ctx)...)
	l.Info(msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...zap.Field) {
	l.logger.Warn(msg, fields...)
}

func (l *Logger) WarnWithContext(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, l.getContextFields(ctx)...)
	l.Warn(msg, fields...)
}

func (l *Logger) Error(msg interface{}, fields ...zap.Field) {
	if msg == nil {
		return
	}

	switch value := msg.(type) {
	case string:
		l.logger.Error(value, fields...)
	case error:
		l.logger.Error(value.Error(), fields...)
	case fmt.Stringer:
		l.logger.Error(value.String(), fields...)
	default:
		l.logger.Error(fmt.Sprintf("%v", value), fields...)
	}
}

func (l *Logger) ErrorWithContext(ctx context.Context, err error, fields ...zap.Field) error {
	fields = append(fields, l.getContextFields(ctx)...)

	l.Error(err, fields...)

	return err
}

func (l *Logger) Fatal(msg interface{}, fields ...zap.Field) {
	if msg == nil {
		return
	}

	switch msg.(type) {
	case string:
		if v, ok := msg.(string); ok {
			l.logger.Fatal(v, fields...)
		}
	case error:
		if v, ok := msg.(error); ok {
			l.logger.Fatal(v.Error(), fields...)
		}
	case fmt.Stringer:
		if v, ok := msg.(fmt.Stringer); ok {
			l.logger.Fatal(v.String(), fields...)
		}
	default:
		l.logger.Fatal(fmt.Sprintf("%v", msg), fields...)
	}
}

func (l *Logger) FatalWithContext(ctx context.Context, err error, fields ...zap.Field) error {
	fields = append(fields, l.getContextFields(ctx)...)

	l.Fatal(err, fields...)

	return err
}
