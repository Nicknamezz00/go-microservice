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

package services

import (
	"context"
	"github.com/Nicknamezz00/go-microservice/api/proto"
	"github.com/Nicknamezz00/go-microservice/internal/pkg/models"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type ProductsService interface {
	Get(c context.Context, ID uint64) (*models.Product, error)
}

type DefaultProductsService struct {
	logger     *zap.Logger
	detailsSvc proto.DetailsClient
	ratingsSvc proto.RatingsClient
	reviewsSvc proto.ReviewsClient
}

func NewProductService(logger *zap.Logger, detailsSvc proto.DetailsClient, ratingsSvc proto.RatingsClient, reviewsSvc proto.ReviewsClient) ProductsService {
	return &DefaultProductsService{
		logger:     logger.With(zap.String("type", "DefaultProductsService")),
		detailsSvc: detailsSvc,
		ratingsSvc: ratingsSvc,
		reviewsSvc: reviewsSvc,
	}
}

func (s *DefaultProductsService) Get(c context.Context, ID uint64) (p *models.Product, err error) {
	var detail *models.Detail
	{
		req := &proto.GetDetailsRequest{Id: ID}
		d, err := s.detailsSvc.Get(c, req)
		if err != nil {
			return nil, errors.Wrap(err, "get product detail error")
		}
		detail = &models.Detail{
			ID:          d.Id,
			Name:        d.Name,
			Price:       d.Price,
			CreatedTime: d.CreatedTime.AsTime(),
		}
	}
	var rating *models.Rating
	{
		req := &proto.GetRatingsRequest{ProductID: ID}
		r, err := s.ratingsSvc.Get(c, req)
		if err != nil {
			return nil, errors.Wrap(err, "get product rating error")
		}
		rating = &models.Rating{
			ID:          r.Id,
			ProductID:   r.ProductID,
			Score:       r.Score,
			UpdatedTime: r.UpdatedTime.AsTime(),
		}
	}
	var reviews []*models.Review
	{
		req := &proto.QueryReviewsRequest{ProductID: ID}
		resp, err := s.reviewsSvc.Query(c, req)
		if err != nil {
			return nil, errors.Wrap(err, "get product reviews error")
		}
		reviews = make([]*models.Review, 0, len(resp.Reviews))
		for _, item := range resp.Reviews {
			r := &models.Review{
				ID:          item.Id,
				ProductID:   item.ProductID,
				Message:     item.Message,
				CreatedTime: item.CreatedTime.AsTime(),
			}
			reviews = append(reviews, r)
		}
	}
	return &models.Product{
		Detail:  detail,
		Rating:  rating,
		Reviews: reviews,
	}, nil
}
