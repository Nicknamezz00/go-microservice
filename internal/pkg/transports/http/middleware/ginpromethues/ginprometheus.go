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

package ginpromethues

import (
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	httpHistogram = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "http_server",
		Name:      "requests_seconds",
		Help:      "Histogram of http handlers' response latency (seconds).",
	}, []string{"method", "code", "uri"})
)

func init() {
	prometheus.MustRegister(httpHistogram)
}

type handlerPath struct {
	sync.Map
}

func (h *handlerPath) get(handler string) string {
	v, ok := h.Load(handler)
	if !ok {
		return ""
	}
	return v.(string)
}

func (h *handlerPath) set(ri gin.RouteInfo) {
	h.Store(ri.Handler, ri.Path)
}

type GinPrometheus struct {
	engine         *gin.Engine
	handlerPathMap *handlerPath
	updated        bool
	ignored        map[string]bool
}

// Option configurable
type Option func(*GinPrometheus)

// Ignore add paths to ignore
func Ignore(paths ...string) Option {
	return func(gp *GinPrometheus) {
		for _, path := range paths {
			gp.ignored[path] = true
		}
	}
}

const (
	metricsPath = "/metrics"
	faviconPath = "/favicon.ico"
)

func New(e *gin.Engine, options ...Option) *GinPrometheus {
	if e == nil {
		return nil
	}
	gp := &GinPrometheus{
		engine:         e,
		handlerPathMap: &handlerPath{},
		ignored: map[string]bool{
			metricsPath: true,
			faviconPath: true,
		},
	}
	for _, o := range options {
		o(gp)
	}
	return gp
}

func (gp *GinPrometheus) updatePath() {
	gp.updated = true
	for _, router := range gp.engine.Routes() {
		gp.handlerPathMap.set(router)
	}
}

// Middleware returns a gin prometheus middleware
func (gp *GinPrometheus) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !gp.updated {
			gp.updatePath()
		}
		// filter ignores
		if gp.ignored[c.Request.URL.String()] {
			c.Next()
			return
		}

		start := time.Now()
		c.Next()
		httpHistogram.WithLabelValues(
			c.Request.Method,
			strconv.Itoa(c.Writer.Status()),
			gp.handlerPathMap.get(c.HandlerName()),
		).Observe(time.Since(start).Seconds())
	}
}
