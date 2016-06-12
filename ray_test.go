// Copyright 2016 Josh Deprez
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vec

import (
	"log"
	"math/rand"
	"testing"
)

func TestCellsTouchingSegmentSpecialCases(t *testing.T) {
	tests := []struct {
		start, end I2
		want       int
	}{
		{start: I2{0, 0}, end: I2{0, 0}, want: 1},
		{start: I2{0, 0}, end: I2{16, 0}, want: 2},
		{start: I2{0, 0}, end: I2{-1, 0}, want: 2},
		{start: I2{0, 0}, end: I2{0, -1}, want: 2},
		{start: I2{0, 0}, end: I2{-15, 0}, want: 2},
		{start: I2{0, 0}, end: I2{0, -15}, want: 2},
		{start: I2{0, 0}, end: I2{-16, 0}, want: 2},
		{start: I2{0, 0}, end: I2{0, -16}, want: 2},
		{start: I2{0, 0}, end: I2{15, -16}, want: 2},
		{start: I2{8, 8}, end: I2{9, 9}, want: 1},
		{start: I2{0, 0}, end: I2{15, 15}, want: 1},
		{start: I2{0, 0}, end: I2{159, 15}, want: 10},
	}

	for i, test := range tests {
		got := 0
		touch := func(I2) bool {
			got++
			return true
		}
		CellsTouchingSegment(I2{16, 16}, test.start, test.end, touch)
		if got != test.want {
			t.Errorf("Test %d: got %d want %d", i, got, test.want)
		}
	}
}

func TestCellsTouchingSegmentHalts(t *testing.T) {
	for i := 0; i < 1000; i++ {
		start := I2{rand.Intn(1000) - 500, rand.Intn(1000) - 500}
		end := I2{rand.Intn(1000) - 500, rand.Intn(1000) - 500}
		test.Logf("test %d: %v-%v", i, start, end)
		CellsTouchingSegment(I2{16, 16}, start, end, func(I2) bool { return true })
		test.Logf("test %d pass", i)
	}
}
