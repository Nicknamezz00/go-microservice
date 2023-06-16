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

package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/google/wire"

	"github.com/Nicknamezz00/go-microservice/internal/pkg/utils/netutil"
	"github.com/pkg/errors"

	"github.com/Nicknamezz00/go-microservice/internal/pkg/transports/http/middleware/ginpromethues"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/gin-contrib/pprof"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/opentracing-contrib/go-gin/ginhttp"
	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Options struct {
	Port int
	Mode string
}

func NewOptions(v *viper.Viper) (*Options, error) {
	var (
		err error
		o   = new(Options)
	)
	if err = v.UnmarshalKey("http", o); err != nil {
		return nil, err
	}
	return o, err
}

type Server struct {
	o          *Options
	app        string
	host       string
	port       int
	logger     *zap.Logger
	router     *gin.Engine
	httpServer http.Server
	consulCli  *consulapi.Client
}

type InitControllers func(r *gin.Engine)

// NewRouter New router, basically configuring gin
func NewRouter(o *Options, logger *zap.Logger, init InitControllers, tracer opentracing.Tracer) *gin.Engine {
	gin.SetMode(o.Mode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(ginzap.RecoveryWithZap(logger, true))
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginpromethues.New(r).Middleware())
	r.Use(ginhttp.Middleware(tracer))

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	pprof.Register(r)

	init(r)
	return r
}

// NewServer returns a server
func NewServer(o *Options, logger *zap.Logger, router *gin.Engine, consulCli *consulapi.Client) (*Server, error) {
	var s = &Server{
		o:         o,
		logger:    logger.With(zap.String("type", "http.Server")),
		router:    router,
		consulCli: consulCli,
	}
	return s, nil
}

func (s *Server) Application(name string) {
	s.app = name
}

func (s *Server) register() error {
	addr := fmt.Sprintf("%s:%d", s.host, s.port)
	check := &consulapi.AgentServiceCheck{
		Interval:                       "10s",
		TCP:                            addr,
		DeregisterCriticalServiceAfter: "60m",
	}
	id := fmt.Sprintf("%s[%s:%d]", s.app, s.host, s.port)
	svcReg := &consulapi.AgentServiceRegistration{
		ID:                id,
		Name:              s.app,
		Tags:              []string{"http"},
		Port:              s.port,
		Address:           s.host,
		EnableTagOverride: true,
		Check:             check,
		Checks:            nil,
	}
	if err := s.consulCli.Agent().ServiceRegister(svcReg); err != nil {
		return errors.Wrap(err, "register http service error")
	}
	s.logger.Info("register http service success", zap.String("id", id))
	return nil
}

func (s *Server) deregister() error {
	id := fmt.Sprintf("%s[%s:%d]", s.app, s.host, s.port)

	err := s.consulCli.Agent().ServiceDeregister(id)
	if err != nil {
		return errors.Wrapf(err, "deregister http service error with [id = %s]", id)
	}
	s.logger.Info("deregister http service success ", zap.String("service", id))

	return nil
}

func (s *Server) Start() error {
	s.port = s.o.Port
	if s.port == 0 {
		s.port = netutil.GetAvailablePort()
	}
	s.host = netutil.GetLocalIPv4()
	if s.host == "" {
		return errors.New("get local ipv4 error")
	}

	addr := fmt.Sprintf("%s:%d", s.host, s.port)
	s.logger.Info("http server starting ... ", zap.String("addr", addr))

	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Fatal("start http server error", zap.Error(err))
			return
		}
	}()
	if err := s.register(); err != nil {
		return errors.Wrap(err, "register http server error")
	}
	return nil
}

func (s *Server) Stop() error {
	s.logger.Info("stopping http server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.deregister(); err != nil {
		return errors.Wrap(err, "deregister http server error")
	}
	if err := s.httpServer.Shutdown(ctx); err != nil {
		return errors.Wrap(err, "shutdown http server error")
	}
	return nil
}

var ProviderSet = wire.NewSet(NewServer, NewRouter, NewOptions)
