/*
Copyright © 2023 Michael Fero

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package log

import (
	"fmt"

	"github.com/mikefero/tpl/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type LoggerCommandType int

const (
	LoggerCommandTypeMigrate LoggerCommandType = iota
)

func (l LoggerCommandType) String() string {
	return [...]string{
		"migrate",
	}[l]
}

func NewLogger(config *config.Config, commandType LoggerCommandType) (*zap.Logger, error) {
	zapLoggerLevel, err := zapcore.ParseLevel(config.Logger.Level)
	if err != nil {
		return nil, fmt.Errorf("unable to parse log level: %w", err)
	}

	// Add daily log rotator for zap logger
	logger := &lumberjack.Logger{
		Filename:   config.Logger.Filename,
		MaxSize:    0, // unlimited
		MaxBackups: config.Logger.Retention,
		MaxAge:     config.Logger.Retention,
		Compress:   true,
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(logger),
		zapLoggerLevel,
	).With([]zapcore.Field{
		zap.String("command", commandType.String()),
	})
	zapLogger := zap.New(core)

	return zapLogger, nil
}
