/*
 * MIT License
 *
 * Copyright (c) 2023 Runze Wu
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 *
 */

package log

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

type Options struct {
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Level      string
	Stdout     bool
}

func NewOptions(v *viper.Viper) (*Options, error) {
	var (
		err error
		o   = new(Options)
	)
	if err = v.UnmarshalKey("log", o); err != nil {
		return nil, err
	}
	return o, nil
}

// New use `Options` to init zap logger
func New(o *Options) (*zap.Logger, error) {
	var (
		err   error
		level = zap.NewAtomicLevel()
	)

	if err = level.UnmarshalText([]byte(o.Level)); err != nil {
		return nil, err
	}

	ws := zapcore.AddSync(&lumberjack.Logger{
		Filename:   o.Filename,
		MaxSize:    o.MaxSize,
		MaxAge:     o.MaxAge,
		MaxBackups: o.MaxBackups,
	})
	cw := zapcore.Lock(os.Stdout)

	cores := make([]zapcore.Core, 0, 2)
	enc := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	cores = append(cores, zapcore.NewCore(enc, ws, level))

	if o.Stdout {
		consoleEnc := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		cores = append(cores, zapcore.NewCore(consoleEnc, cw, level))
	}

	core := zapcore.NewTee(cores...)
	logger := zap.New(core)
	zap.ReplaceGlobals(logger)

	return logger, nil
}

var ProviderSet = wire.NewSet(New, NewOptions)
