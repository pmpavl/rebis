package rebis

import (
	"io"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	Info = "[INFO] "
	Set  = "[SET]  "
)

type logger struct {
	logOut   io.Writer
	logLevel string
}

func startLogger(path string, level int8) {
	logger, _ := zap.Config{
		Encoding:    "console",
		Level:       zap.NewAtomicLevelAt(zapcore.Level(level)),
		OutputPaths: []string{path},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "message",
			LevelKey:     "level",
			EncodeLevel:  zapcore.CapitalLevelEncoder,
			TimeKey:      "time",
			EncodeTime:   zapcore.RFC3339TimeEncoder,
			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}.Build()

	logger.Info("This is an INFO message")

}
