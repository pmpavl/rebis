package rebis

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func runLogger(c *cache, path string, level int8) (err error) {
	if path != DefaultLoggerPath {
		path += "logs" + strconv.Itoa(int(time.Now().Unix())) + ".log"
	}
	c.logger, err = zap.Config{
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
	if err != nil {
		return errors.New(fmt.Sprintf("wrong definition of zap logger %s", err.Error()))
	}

	c.logger.Info(
		"START LOGGER",
		zap.String("log path", path),
		zap.String("log level", zap.NewAtomicLevelAt(zapcore.Level(level)).String()),
	)
	return nil
}
