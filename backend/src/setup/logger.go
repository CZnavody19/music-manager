package setup

import (
	"context"
	"os"
	"sort"
	"time"

	"github.com/CZnavody19/music-manager/src/config"
	"go.elastic.co/ecszap"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLogger(configuration config.Config) {
	cores := []zapcore.Core{}
	if configuration.LoggingConfig.EnableDebugLogger {
		lowCore, highCore := setupDebugLogger(configuration.LoggingConfig)
		cores = append(cores, lowCore)
		cores = append(cores, highCore)
	}
	if configuration.LoggingConfig.EnableFileLogger {
		cores = append(cores, setupFileLogger(configuration.LoggingConfig))
	}
	core := zapcore.NewTee(cores...)
	logger := zap.New(core, buildOptions(configuration)...)
	zap.ReplaceGlobals(logger)
}

func setupFileLogger(logConfig config.LoggingConfig) zapcore.Core {
	var level zapcore.Level
	err := level.Set(logConfig.FileLogLevel)
	if err != nil {
		// We want to panic here, since logger is important
		panic(err)
	}

	encoderConfig := ecszap.NewDefaultEncoderConfig()
	if logConfig.FileLogOutput == "stdout" {
		sink, _, err := zap.Open(logConfig.FileLogOutput)
		if err != nil {
			// We want to panic here, since logger is important
			panic(err)
		}
		return ecszap.NewCore(encoderConfig, sink, level)
	} else {
		// rolling file appender
		w := zapcore.AddSync(&lumberjack.Logger{
			Filename:   logConfig.FileLogOutput,
			MaxSize:    100, // megabytes
			MaxBackups: 2,
			MaxAge:     2, // days
		})

		return ecszap.NewCore(encoderConfig, w, level)
	}
}

func setupDebugLogger(logConfig config.LoggingConfig) (zapcore.Core, zapcore.Core) {
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	errorPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl > zap.ErrorLevel
	})

	priority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl <= zap.ErrorLevel
	})

	consoleWriter := zapcore.Lock(os.Stdout)
	consoleError := zapcore.Lock(os.Stderr)

	return zapcore.NewCore(consoleEncoder, consoleWriter, priority), zapcore.NewCore(consoleEncoder, consoleError, errorPriority)
}

func buildOptions(cfg config.Config) []zap.Option {
	opts := []zap.Option{}

	opts = append(opts, zap.Context(func(ctx context.Context) []zap.Field {
		fields := []zap.Field{}
		spanCtx := trace.SpanContextFromContext(ctx)
		if spanCtx.IsSampled() && spanCtx.HasTraceID() {
			traceId := spanCtx.TraceID().String()
			fields = append(fields, zap.String("trace.id", traceId))
			fields = append(fields, zap.String("traceID", traceId))
		}
		if spanCtx.IsSampled() && spanCtx.HasSpanID() {
			spanId := spanCtx.SpanID().String()
			fields = append(fields, zap.String("span.id", spanId))
			fields = append(fields, zap.String("spanID", spanId))
		}

		return fields
	}))

	if cfg.LoggingConfig.LoggerConfig.Development {
		opts = append(opts, zap.Development())
	}

	if !cfg.LoggingConfig.LoggerConfig.DisableCaller {
		opts = append(opts, zap.AddCaller())
	}

	stackLevel := zap.ErrorLevel
	if cfg.LoggingConfig.LoggerConfig.Development {
		stackLevel = zap.WarnLevel
	}
	if !cfg.LoggingConfig.LoggerConfig.DisableStacktrace {
		opts = append(opts, zap.AddStacktrace(stackLevel))
	}

	if cfg.LoggingConfig.LoggerConfig.SamplingEnable {
		scfg := cfg.LoggingConfig.LoggerConfig.Sampling
		opts = append(opts, zap.WrapCore(func(core zapcore.Core) zapcore.Core {
			var samplerOpts []zapcore.SamplerOption
			return zapcore.NewSamplerWithOptions(
				core,
				time.Second,
				scfg.Initial,
				scfg.Thereafter,
				samplerOpts...,
			)
		}))
	}

	if len(cfg.LoggingConfig.LoggerConfig.InitialFields) > 0 {
		fs := make([]zap.Field, 0, len(cfg.LoggingConfig.LoggerConfig.InitialFields))
		keys := make([]string, 0, len(cfg.LoggingConfig.LoggerConfig.InitialFields))
		for k := range cfg.LoggingConfig.LoggerConfig.InitialFields {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			fs = append(fs, zap.Any(k, cfg.LoggingConfig.LoggerConfig.InitialFields[k]))
		}
		opts = append(opts, zap.Fields(fs...))
	}

	serviceFields := []zapcore.Field{
		zap.String("name", "test"),
		zap.String("environment", "test"),
	}

	opts = append(opts, zap.Fields(serviceFields...))

	return opts
}
