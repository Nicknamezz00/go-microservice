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

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/pkg/errors"

	"github.com/Nicknamezz00/go-microservice/api/proto"

	"github.com/Nicknamezz00/go-microservice/internal/app/details/services"
	"go.uber.org/zap"
)

type DetailsServer struct {
	logger  *zap.Logger
	service services.DetailsService
}

func NewDetailsServer(logger *zap.Logger, s services.DetailsService) (*DetailsServer, error) {
	return &DetailsServer{
		logger:  logger,
		service: s,
	}, nil
}

func (s *DetailsServer) Get(ctx context.Context, req *proto.GetDetailsRequest) (*proto.Detail, error) {
	v, err := s.service.Get(req.Id)
	if err != nil {
		return nil, errors.Wrap(err, "details grpc server get error")
	}
	createTime := timestamppb.New(v.CreatedTime)
	resp := &proto.Detail{
		Id:          uint64(v.ID),
		Name:        v.Name,
		Price:       v.Price,
		CreatedTime: createTime,
	}
	return resp, nil
}
