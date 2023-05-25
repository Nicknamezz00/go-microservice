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

package jaeger

import (
	"github.com/google/wire"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-lib/metrics/prometheus"
	"go.uber.org/zap"
)

func NewConfiguration(v *viper.Viper, logger *zap.Logger) (*config.Configuration, error) {
	var (
		err error
		c   = new(config.Configuration)
	)
	if err = v.UnmarshalKey("jaeger", c); err != nil {
		return nil, errors.Wrap(err, "unmarshal jaeger configuration error")
	}
	logger.Info("load jaeger configuration succuess")
	return c, nil
}

func NewJaeger(c *config.Configuration) (opentracing.Tracer, error) {
	metricsFactory := prometheus.New()
	tracer, _, err := c.NewTracer(config.Metrics(metricsFactory))
	if err != nil {
		return nil, errors.Wrap(err, "new jaeger tracer error")
	}
	return tracer, nil
}

var ProviderSet = wire.NewSet(NewJaeger, NewConfiguration)
