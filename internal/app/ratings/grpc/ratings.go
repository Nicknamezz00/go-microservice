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

	"github.com/Nicknamezz00/go-microservice/api/proto"
	"github.com/Nicknamezz00/go-microservice/internal/app/ratings/services"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type RatingsServer struct {
	logger  *zap.Logger
	service services.RatingsService
}

func NewRatingsServer(logger *zap.Logger, service services.RatingsService) (*RatingsServer, error) {
	return &RatingsServer{
		logger:  logger,
		service: service,
	}, nil
}

func (s *RatingsServer) Get(ctx context.Context, req *proto.GetRatingsRequest) (*proto.Rating, error) {
	r, err := s.service.Get(req.ProductID)
	if err != nil {
		return nil, errors.Wrap(err, "ratings grpc server get rating error")
	}
	updateTime := timestamppb.New(r.UpdatedTime)
	resp := &proto.Rating{
		Id:          uint64(r.ID),
		ProductID:   r.ProductID,
		Score:       r.Score,
		UpdatedTime: updateTime,
	}
	return resp, nil
}
