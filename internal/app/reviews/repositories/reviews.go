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

package repositories

import (
	"github.com/Nicknamezz00/go-microservice/internal/pkg/models"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ReviewsRepository interface {
	Query(productID uint64) (reviews []*models.Review, err error)
}

type MySQLReviewsRepository struct {
	logger *zap.Logger
	db     *gorm.DB
}

func NewMySQLReviewsRepository(logger *zap.Logger, db *gorm.DB) ReviewsRepository {
	return &MySQLReviewsRepository{
		logger: logger.With(zap.String("type", "ReviewsRepository")),
		db:     db,
	}
}

func (s *MySQLReviewsRepository) Query(productID uint64) (reviews []*models.Review, err error) {
	if err = s.db.Table("reviews").Where("product_id = ?", productID).Find(&reviews).Error; err != nil {
		return nil, errors.Wrapf(err, "get reviews error [product_id = %d]", productID)
	}
	return
}