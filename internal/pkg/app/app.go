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

package app

import (
	"github.com/google/wire"
	"os"
	"os/signal"
	"syscall"

	"github.com/Nicknamezz00/go-microservice/internal/pkg/transports/grpc"
	"github.com/Nicknamezz00/go-microservice/internal/pkg/transports/http"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Application struct {
	name       string
	logger     *zap.Logger
	httpServer *http.Server
	grpcServer *grpc.Server
}

type Option func(app *Application) error

func HttpServerOption(hs *http.Server) Option {
	return func(a *Application) error {
		hs.Application(a.name)
		a.httpServer = hs
		return nil
	}
}

func GrpcServerOption(gs *grpc.Server) Option {
	return func(a *Application) error {
		gs.Application(a.name)
		a.grpcServer = gs
		return nil
	}
}

func NewApplication(name string, logger *zap.Logger, options ...Option) (*Application, error) {
	app := &Application{
		name:   name,
		logger: logger.With(zap.String("type", "Application")),
	}
	for _, option := range options {
		if err := option(app); err != nil {
			return nil, err
		}
	}
	return app, nil
}

func (a *Application) Start() error {
	if a.httpServer != nil {
		if err := a.httpServer.Start(); err != nil {
			return errors.Wrap(err, "http server start error")
		}
	}
	if a.grpcServer != nil {
		if err := a.grpcServer.Start(); err != nil {
			return errors.Wrap(err, "grpc server start error")
		}
	}
	return nil
}

func (a *Application) AwaitSignal() {
	c := make(chan os.Signal, 1)
	signal.Reset(syscall.SIGTERM, syscall.SIGINT)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	select {
	case s := <-c:
		a.logger.Info("receive signal", zap.String("signal", s.String()))
		if a.httpServer != nil {
			if err := a.httpServer.Stop(); err != nil {
				a.logger.Warn("stop http server error", zap.Error(err))
			}
		}
		if a.grpcServer != nil {
			if err := a.grpcServer.Stop(); err != nil {
				a.logger.Warn("stop grpc server error", zap.Error(err))
			}
		}
		os.Exit(0)
	}
}

var ProviderSet = wire.NewSet(NewApplication)
