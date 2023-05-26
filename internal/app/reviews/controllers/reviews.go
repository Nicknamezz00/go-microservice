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

package controllers

import (
	"github.com/Nicknamezz00/go-microservice/internal/app/reviews/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type ReviewsController struct {
	logger  *zap.Logger
	service services.ReviewsService
}

func NewReviewsController(logger *zap.Logger, service services.ReviewsService) *ReviewsController {
	return &ReviewsController{logger: logger, service: service}
}

func (rc *ReviewsController) Query(c *gin.Context) {
	id, err := strconv.ParseUint(c.Query("productID"), 10, 64)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	reviews, err := rc.service.Query(id)
	if err != nil {
		rc.logger.Error("reviews controller query review by product_id error", zap.Error(err))
		c.String(http.StatusInternalServerError, "%+v", err)
		return
	}
	c.JSON(http.StatusOK, reviews)
}
