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
	"flag"
	"testing"

	"github.com/stretchr/testify/assert"
)

var configFile = flag.String("f", "../../../../configs/details.yml", "set config file which viper will loading.")

func TestDetailsRepository_Get(t *testing.T) {
	flag.Parse()

	sto, err := CreateDetailsRepository(*configFile)
	if err != nil {
		t.Fatalf("create product Repository error,%+v", err)
	}

	tests := []struct {
		name     string
		id       uint64
		expected bool
	}{
		{"id=1", 1, true},
		{"id=2", 2, true},
		{"id=3", 3, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := sto.Get(test.id)

			if test.expected {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
