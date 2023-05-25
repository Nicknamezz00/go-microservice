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
	"net/http"
	"strconv"

	"github.com/Nicknamezz00/go-microservice/internal/pkg/models"
	"github.com/pkg/errors"

	"github.com/Nicknamezz00/go-microservice/internal/app/details/repositories"

	"github.com/Nicknamezz00/go-microservice/internal/app/details/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DetailsController struct {
	logger  *zap.Logger
	service services.DetailsService
}

func NewDetailsController(logger *zap.Logger, s services.DetailsService) *DetailsController {
	return &DetailsController{
		logger:  logger,
		service: s,
	}
}

func (dc *DetailsController) Get(c *gin.Context) {
	ID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	d, err := dc.service.Get(ID)
	if err != nil {
		dc.logger.Error("get detail by id error", zap.Error(err))
		c.String(http.StatusInternalServerError, "%+v", err)
		return
	}
	c.JSON(http.StatusOK, d)
}

type DefaultDetailsService struct {
	logger     *zap.Logger
	Repository repositories.DetailsRepository
}

func (s *DefaultDetailsService) Get(ID uint64) (p *models.Detail, err error) {
	if p, err = s.Repository.Get(ID); err != nil {
		return nil, errors.Wrap(err, "detail service get detail error")
	}
	return
}
