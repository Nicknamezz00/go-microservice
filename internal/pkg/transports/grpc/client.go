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
	"context"
	"fmt"
	"github.com/Nicknamezz00/go-microservice/internal/pkg/consul"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type ClientOptions struct {
	Wait            time.Duration
	Tag             string
	GrpcDialOptions []grpc.DialOption
}

func NewClientOptions(v *viper.Viper, tracer opentracing.Tracer) (*ClientOptions, error) {
	var (
		o   = new(ClientOptions)
		err error
	)
	if err = v.UnmarshalKey("grpc.client", o); err != nil {
		return nil, err
	}
	grpc_prometheus.EnableClientHandlingTimeHistogram()
	o.GrpcDialOptions = append(o.GrpcDialOptions,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(
			grpc_prometheus.UnaryClientInterceptor,
			otgrpc.OpenTracingClientInterceptor(tracer)),
		),
		grpc.WithStreamInterceptor(grpc_middleware.ChainStreamClient(
			grpc_prometheus.StreamClientInterceptor,
			otgrpc.OpenTracingStreamClientInterceptor(tracer)),
		),
	)
	return o, nil
}

type ClientOptional func(o *ClientOptions)

func WithTimeout(d time.Duration) ClientOptional {
	return func(o *ClientOptions) {
		o.Wait = d
	}
}

func WithTag(tag string) ClientOptional {
	return func(o *ClientOptions) {
		o.Tag = tag
	}
}

func WithGrpcDialOptions(options ...grpc.DialOption) ClientOptional {
	return func(o *ClientOptions) {
		o.GrpcDialOptions = append(o.GrpcDialOptions, options...)
	}
}

type Client struct {
	consulOptions *consul.Options
	o             *ClientOptions
}

func NewClient(consulOptions *consul.Options, o *ClientOptions) (*Client, error) {
	return &Client{
		consulOptions: consulOptions,
		o:             o,
	}, nil
}

func (c *Client) Dial(service string, options ...ClientOptional) (*grpc.ClientConn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	co := &ClientOptions{
		Wait:            c.o.Wait,
		Tag:             c.o.Tag,
		GrpcDialOptions: c.o.GrpcDialOptions,
	}
	for _, option := range options {
		option(co)
	}
	url := fmt.Sprintf("consul://%s/%s?wait=%s&tag=%s", c.consulOptions.Addr, service, co.Wait, co.Tag)
	conn, err := grpc.DialContext(ctx, url, co.GrpcDialOptions...)
	if err != nil {
		return nil, errors.Wrap(err, "grpc dial context error")
	}

	return conn, nil
}
