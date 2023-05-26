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
	"github.com/Nicknamezz00/go-microservice/internal/app/details/repositories"
	"github.com/Nicknamezz00/go-microservice/internal/pkg/models"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type DetailsService interface {
	Get(ID uint64) (*models.Detail, error)
}

type DefaultDetailsService struct {
	logger     *zap.Logger
	Repository repositories.DetailsRepository
}

func NewDetailsService(logger *zap.Logger, repo repositories.DetailsRepository) DetailsService {
	return &DefaultDetailsService{
		logger:     logger.With(zap.String("type", "DefaultDetailsService")),
		Repository: repo,
	}
}

func (s *DefaultDetailsService) Get(ID uint64) (d *models.Detail, err error) {
	if d, err = s.Repository.Get(ID); err != nil {
		return nil, errors.Wrap(err, "default details service get error")
	}
	return
}
