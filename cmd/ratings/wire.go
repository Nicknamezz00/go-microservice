//go:build wireinject

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

package main

import (
	"github.com/Nicknamezz00/go-microservice/internal/app/ratings"
	"github.com/Nicknamezz00/go-microservice/internal/app/ratings/controllers"
	"github.com/Nicknamezz00/go-microservice/internal/app/ratings/grpc"
	"github.com/Nicknamezz00/go-microservice/internal/app/ratings/repositories"
	"github.com/Nicknamezz00/go-microservice/internal/app/ratings/services"
	"github.com/Nicknamezz00/go-microservice/internal/pkg/app"
	"github.com/Nicknamezz00/go-microservice/internal/pkg/config"
	"github.com/Nicknamezz00/go-microservice/internal/pkg/consul"
	"github.com/Nicknamezz00/go-microservice/internal/pkg/database"
	"github.com/Nicknamezz00/go-microservice/internal/pkg/jaeger"
	"github.com/Nicknamezz00/go-microservice/internal/pkg/log"
	grpcserver "github.com/Nicknamezz00/go-microservice/internal/pkg/transports/grpc"
	"github.com/Nicknamezz00/go-microservice/internal/pkg/transports/http"
	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	services.ProviderSet,
	repositories.ProviderSet,
	consul.ProviderSet,
	jaeger.ProviderSet,
	http.ProviderSet,
	grpc.ProviderSet,
	ratings.ProviderSet,
	controllers.ProviderSet,
	grpcserver.ProviderSet,
)

func CreateApp(f string) (*app.Application, error) {
	panic(wire.Build(providerSet))
}
