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

package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/Nicknamezz00/go-microservice/internal/pkg/utils/netutil"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type ServerOptions struct {
	Port int
}

func NewServerOptions(v *viper.Viper) (*ServerOptions, error) {
	o := new(ServerOptions)
	if err := v.UnmarshalKey("grpc", o); err != nil {
		return nil, err
	}
	return o, nil
}

type Server struct {
	o         *ServerOptions
	app       string
	host      string
	port      int
	logger    *zap.Logger
	server    *grpc.Server
	consulCli *consulapi.Client
}

type InitServers func(s *grpc.Server)

func NewServer(o *ServerOptions, logger *zap.Logger, init InitServers, consulCli *consulapi.Client, tracer opentracing.Tracer) (*Server, error) {
	var gs *grpc.Server
	logger = logger.With(zap.String("type", "grpc"))

	grpc_prometheus.EnableHandlingTimeHistogram()
	gs = grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_prometheus.StreamServerInterceptor,
			grpc_zap.StreamServerInterceptor(logger),
			grpc_recovery.StreamServerInterceptor(),
			otgrpc.OpenTracingStreamServerInterceptor(tracer),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
			grpc_zap.UnaryServerInterceptor(logger),
			grpc_recovery.UnaryServerInterceptor(),
			otgrpc.OpenTracingServerInterceptor(tracer),
		)),
	)
	init(gs)

	return &Server{
		o:         o,
		logger:    logger.With(zap.String("type", "grpc.Server")),
		server:    gs,
		consulCli: consulCli,
	}, nil
}

func (s *Server) Application(name string) {
	s.app = name
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
	s.logger.Info("grpc server starting ...", zap.String("addr", addr))

	go func() {
		listen, err := net.Listen("tcp", addr)
		if err != nil {
			log.Fatalf("failed to listen to: %v", err)
		}
		if err := s.server.Serve(listen); err != nil {
			s.logger.Fatal("failed to serve: %v", zap.Error(err))
		}
	}()
	if err := s.register(); err != nil {
		return errors.Wrap(err, "register grpc server error")
	}
	return nil
}

func (s *Server) register() error {
	addr := fmt.Sprintf("%s:%d", s.host, s.port)
	for key, _ := range s.server.GetServiceInfo() {
		check := &consulapi.AgentServiceCheck{
			Interval:                       "10s",
			DeregisterCriticalServiceAfter: "60m",
			TCP:                            addr,
		}
		id := fmt.Sprintf("%s[%s:%d]", key, s.host, s.port)
		svc := &consulapi.AgentServiceRegistration{
			ID:                id,
			Name:              key,
			Tags:              []string{"grpc"},
			Port:              s.port,
			Address:           s.host,
			EnableTagOverride: true,
			Check:             check,
			Checks:            nil,
		}
		err := s.consulCli.Agent().ServiceRegister(svc)
		if err != nil {
			return errors.Wrap(err, "register grpc service error")
		}
		s.logger.Info("register grpc service success ", zap.String("id", id))
	}
	return nil
}

func (s *Server) deregister() error {
	for key, _ := range s.server.GetServiceInfo() {
		id := fmt.Sprintf("%s[%s:%d]", key, s.host, s.port)
		err := s.consulCli.Agent().ServiceDeregister(id)
		if err != nil {
			return errors.Wrapf(err, "deregister service error with [id = %s]", id)
		}
		s.logger.Info("deregister service success ", zap.String("id", id))
	}
	return nil
}

func (s *Server) Stop() error {
	s.logger.Info("grpc server stopping ...")
	if err := s.deregister(); err != nil {
		return errors.Wrap(err, "deregister grpc server error")
	}
	s.server.GracefulStop()
	return nil
}
