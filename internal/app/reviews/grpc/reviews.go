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
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/Nicknamezz00/go-microservice/api/proto"
	"github.com/Nicknamezz00/go-microservice/internal/app/reviews/services"
	"go.uber.org/zap"
)

type ReviewsServer struct {
	logger  *zap.Logger
	service services.ReviewsService
}

func NewReviewsServer(logger *zap.Logger, service services.ReviewsService) (*ReviewsServer, error) {
	return &ReviewsServer{
		logger:  logger,
		service: service,
	}, nil
}

func (s *ReviewsServer) Query(ctx context.Context, req *proto.QueryReviewsRequest) (*proto.QueryReviewsResponse, error) {
	reviews, err := s.service.Query(req.ProductID)
	if err != nil {
		return nil, errors.Wrap(err, "reviews grpc server query reviews error")
	}
	resp := &proto.QueryReviewsResponse{
		Reviews: make([]*proto.Review, 0, len(reviews)),
	}
	for _, review := range reviews {
		createdTime := timestamppb.New(review.CreatedTime)
		r := &proto.Review{
			Id:          uint64(review.ID),
			ProductID:   review.ProductID,
			Message:     review.Message,
			CreatedTime: createdTime,
		}
		resp.Reviews = append(resp.Reviews, r)
	}
	return resp, nil
}
