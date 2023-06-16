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
	"encoding/json"
	"flag"
	"fmt"
	"github.com/Nicknamezz00/go-microservice/internal/pkg/models"
	"github.com/Nicknamezz00/go-microservice/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine
var configFile = flag.String("f", "details.yml", "config file which viper loads")

func setup() {
	r = gin.New()
}

func TestDetailsController_Get(t *testing.T) {
	flag.Parse()
	setup()
	repo := new(mocks.DetailsRepository)
	repo.On("Get", mock.AnythingOfType("uint64")).Return(func(id uint64) (d *models.Detail) {
		return &models.Detail{ID: id}
	}, func(id uint64) error {
		return nil
	})
	c, err := CreateDetailsController(*configFile, repo)
	if err != nil {
		t.Fatalf("create detail controller error: %+v", err)
	}
	r.GET("/proto/:id", c.Get)
	tests := []struct {
		name     string
		id       uint64
		expected uint64
	}{
		{"1", 1, 1},
		{"2", 2, 2},
		{"3", 2, 3},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uri := fmt.Sprintf("/proto/%d", test.id)
			req := httptest.NewRequest("GET", uri, nil)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			resp := w.Result()
			defer func() {
				_ = resp.Body.Close()
			}()
			body, _ := ioutil.ReadAll(resp.Body)
			d := new(models.Detail)
			if err := json.Unmarshal(body, d); err != nil {
				t.Errorf("unmarshal response body error: %v", err)
			}
			assert.Equal(t, test.expected, d.ID)
		})
	}
}
