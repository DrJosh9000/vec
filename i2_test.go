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

func TestSegmentIntersectI(t *testing.T) {
	tests := []struct {
		p, q, a, b I2
		wantP      I2
		want       bool
	}{
		{I2{0, 0}, I2{2, 2}, I2{0, 2}, I2{2, 0}, I2{1, 1}, true},
		{I2{0, 2}, I2{2, 0}, I2{0, 0}, I2{2, 2}, I2{1, 1}, true},
		{I2{0, 0}, I2{1, 1}, I2{0, 0}, I2{0, 1}, I2{0, 0}, true},
		{I2{0, 0}, I2{1, 1}, I2{1, 0}, I2{1, 1}, I2{1, 1}, true}, // Both ends of line included.
		{I2{-1000000, 0}, I2{1000000, 0}, I2{-10, -10}, I2{10, 10}, I2{0, 0}, true},
		{I2{-1000000, 0}, I2{1000000, 0}, I2{-1000000, -1}, I2{1000000, 1}, I2{0, 0}, true}, // Nearly parallel
		{I2{-1, 0}, I2{1, 0}, I2{-1, -1}, I2{1, -1}, I2{}, false},                           // Parallel
		{I2{-1, -1}, I2{1, 1}, I2{0, 3}, I2{3, 0}, I2{}, false},
		{I2{0, 32}, I2{32, 32}, I2{64, 0}, I2{0, 64}, I2{32, 32}, true},
		{I2{32, 64}, I2{32, 32}, I2{64, 0}, I2{0, 64}, I2{32, 32}, true},
		{I2{0, 32}, I2{32, 32}, I2{16, 32}, I2{16, 64}, I2{16, 32}, true},
		{I2{32, 64}, I2{32, 32}, I2{32, 48}, I2{0, 48}, I2{32, 48}, true}, // a-b beyond p-q.
	}

	for i, test := range tests {
		gotP, got := SegmentIntersectI(test.p, test.q, test.a, test.b)
		if got != test.want {
			t.Errorf("SegmentIntersectI test #%d: got %t, want %t", i, got, test.want)
		}
		if gotP != test.wantP {
			t.Errorf("SegmentIntersectI test #%d: got point %v, want %v", i, gotP, test.wantP)
		}
	}
}

func TestSegmentNearestPoint(t *testing.T) {
	tests := []struct {
		u, v, p I2
		q       I2
		d       int64
	}{
		{I2{0, 0}, I2{0, 0}, I2{1, 1}, I2{0, 0}, 2},                             // u == v
		{I2{0, 0}, I2{4, 0}, I2{1, 1}, I2{1, 0}, 1},                             // p above line
		{I2{0, 0}, I2{4, 0}, I2{1, -1}, I2{1, 0}, 1},                            // p below line
		{I2{0, 0}, I2{4, 0}, I2{-1, 0}, I2{0, 0}, 1},                            // p left of line
		{I2{0, 0}, I2{4, 0}, I2{5, 0}, I2{4, 0}, 1},                             // p right of line
		{I2{-4, -4}, I2{4, 4}, I2{4, -4}, I2{0, 0}, 32},                         // p below right
		{I2{-4, -4}, I2{4, 4}, I2{-4, 4}, I2{0, 0}, 32},                         // p above left
		{I2{4, -4}, I2{12, 4}, I2{12, -4}, I2{8, 0}, 32},                        // p below right
		{I2{4, -4}, I2{12, 4}, I2{4, 4}, I2{8, 0}, 32},                          // p above left
		{I2{-1000000, 1000000}, I2{1000000, -1000000}, I2{7, 7}, I2{0, 0}, 98},  // big line
		{I2{-1000000, 1000000}, I2{1000000, -1000000}, I2{7, -7}, I2{7, -7}, 0}, // along big line
	}
	for i, test := range tests {
		q, d := SegmentNearestPoint(test.u, test.v, test.p)
		if got, want := q, test.q; got != want {
			t.Errorf("SegmentNearestPoint test #%d: got %v, want %v", i, got, want)
		}
		if got, want := d, test.d; got != want {
			t.Errorf("SegmentNearestPoint test #%d: got %d, want %d", i, got, want)
		}
	}
}
