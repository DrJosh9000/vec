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
	"math"
	"testing"
)

func TestF2Dot(t *testing.T) {
	if got, want := (F2{1, 1}.Dot(F2{1, 1})), 2.0; got != want {
		t.Errorf("dot: got %f want %f", got, want)
	}
}

func TestF2Normal(t *testing.T) {
	if got, want := (F2{0, 5}.Normal()), (F2{-5, 0}); got != want {
		t.Errorf("normal: got %v want %v", got, want)
	}
	if got, want := (F2{5, 0}.Normal()), (F2{0, 5}); got != want {
		t.Errorf("normal: got %v want %v", got, want)
	}
	if got, want := (F2{0, -5}.Normal()), (F2{5, 0}); got != want {
		t.Errorf("normal: got %v want %v", got, want)
	}
	if got, want := (F2{-5, 0}.Normal()), (F2{0, -5}); got != want {
		t.Errorf("normal: got %v want %v", got, want)
	}
}

func TestSegmentIntersect(t *testing.T) {
	tests := []struct {
		p, q, a, b F2
		want       bool
		wantT      float64
	}{
		{F2{0, 0}, F2{1, 1}, F2{0, 1}, F2{1, 0}, true, 0.5},
		{F2{0, 1}, F2{1, 0}, F2{0, 0}, F2{1, 1}, true, 0.5},
		{F2{0, 0}, F2{1, 1}, F2{0, 0}, F2{0, 1}, true, 0.0},
		{F2{0, 0}, F2{1, 1}, F2{1, 0}, F2{1, 1}, false, 0.0}, // Start of line is included, not end.
		{F2{-1000000, 0}, F2{1000000, 0}, F2{-10, -10}, F2{10, 10}, true, 0.5},
		{F2{-1000000, 0}, F2{1000000, 0}, F2{-1000000, -1}, F2{1000000, 1}, true, 0.5}, // Nearly parallel
		{F2{-1, 0}, F2{1, 0}, F2{-1, -1}, F2{1, -1}, false, 0.0},                       // Parallel
		{F2{-1, 0}, F2{1, 0}, F2{-1, 0.0000001}, F2{1, 0.0000001}, false, 0.0},         // Parallel
		{F2{-1, -1}, F2{1, 1}, F2{0, 3}, F2{3, 0}, false, 0.0},                         // a-b beyond p-q.
	}

	for i, test := range tests {
		gotT, got := SegmentIntersect(test.p, test.q, test.a, test.b)
		if got != test.want {
			t.Errorf("SegmentIntersect test #%d: got %t, want %t", i, got, test.want)
		}
		if math.Abs(gotT-test.wantT) > 0.00001 {
			t.Errorf("SegmentIntersect test #%d: got t=%f, want %f", i, gotT, test.wantT)
		}
	}
}
