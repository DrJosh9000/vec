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

import "testing"

func TestOverlap(t *testing.T) {
	tests := []struct {
		r, s Rect
		want bool
	}{
		{r: Rect{I2{}, I2{}}, s: Rect{I2{}, I2{}}, want: false},

		{r: Rect{I2{3, 5}, I2{8, 10}}, s: Rect{I2{8, 10}, I2{13, 15}}, want: false},
		{r: Rect{I2{8, 5}, I2{13, 10}}, s: Rect{I2{8, 10}, I2{13, 15}}, want: false},
		{r: Rect{I2{13, 5}, I2{17, 10}}, s: Rect{I2{8, 10}, I2{13, 15}}, want: false},
		{r: Rect{I2{3, 10}, I2{8, 15}}, s: Rect{I2{8, 10}, I2{13, 15}}, want: false},
		{r: Rect{I2{8, 10}, I2{13, 15}}, s: Rect{I2{8, 10}, I2{13, 15}}, want: true},
		{r: Rect{I2{13, 10}, I2{17, 15}}, s: Rect{I2{8, 10}, I2{13, 15}}, want: false},
		{r: Rect{I2{3, 15}, I2{8, 20}}, s: Rect{I2{8, 10}, I2{13, 15}}, want: false},
		{r: Rect{I2{8, 15}, I2{13, 20}}, s: Rect{I2{8, 10}, I2{13, 15}}, want: false},
		{r: Rect{I2{13, 15}, I2{17, 20}}, s: Rect{I2{8, 10}, I2{13, 15}}, want: false},

		{r: Rect{I2{3, 5}, I2{8, 10}}, s: Rect{I2{5, 7}, I2{10, 12}}, want: true},
		{r: Rect{I2{3, 5}, I2{12, 10}}, s: Rect{I2{5, 7}, I2{10, 12}}, want: true},
		{r: Rect{I2{8, 5}, I2{12, 10}}, s: Rect{I2{5, 7}, I2{10, 12}}, want: true},
		{r: Rect{I2{8, 5}, I2{12, 15}}, s: Rect{I2{5, 7}, I2{10, 12}}, want: true},
		{r: Rect{I2{8, 10}, I2{12, 15}}, s: Rect{I2{5, 7}, I2{10, 12}}, want: true},
		{r: Rect{I2{3, 10}, I2{12, 15}}, s: Rect{I2{5, 7}, I2{10, 12}}, want: true},
		{r: Rect{I2{3, 10}, I2{8, 15}}, s: Rect{I2{5, 7}, I2{10, 12}}, want: true},
		{r: Rect{I2{3, 5}, I2{8, 15}}, s: Rect{I2{5, 7}, I2{10, 12}}, want: true},

		{r: Rect{I2{3, 5}, I2{12, 10}}, s: Rect{I2{5, 7}, I2{10, 12}}, want: true},
		{r: Rect{I2{8, 5}, I2{12, 15}}, s: Rect{I2{5, 7}, I2{10, 12}}, want: true},
		{r: Rect{I2{3, 10}, I2{12, 15}}, s: Rect{I2{5, 7}, I2{10, 12}}, want: true},
		{r: Rect{I2{3, 5}, I2{8, 15}}, s: Rect{I2{5, 7}, I2{10, 12}}, want: true},

		{r: Rect{I2{3, 2}, I2{12, 7}}, s: Rect{I2{5, 7}, I2{10, 12}}, want: false},
		{r: Rect{I2{10, 5}, I2{15, 15}}, s: Rect{I2{5, 7}, I2{10, 12}}, want: false},
		{r: Rect{I2{3, 12}, I2{12, 17}}, s: Rect{I2{5, 7}, I2{10, 12}}, want: false},
		{r: Rect{I2{0, 5}, I2{5, 15}}, s: Rect{I2{5, 7}, I2{10, 12}}, want: false},
	}
	for i, test := range tests {
		if got, want := test.r.Overlaps(test.s), test.want; got != want {
			t.Errorf("Test #%d: got r.Overlaps(s) = %t, want %t", i, got, want)
		}
		if got, want := test.s.Overlaps(test.r), test.want; got != want {
			t.Errorf("Test #%d: got s.Overlaps(r) = %t, want %t", i, got, want)
		}
	}
}
