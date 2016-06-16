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

func divDown(n, d int) (q int) {
	if n >= 0 {
		return n / d
	}
	return (n - d + 1) / d
}

func cell(v, cellSize I2) (w I2) {
	w.X = divDown(v.X, cellSize.X)
	w.Y = divDown(v.Y, cellSize.Y)
	return
}

// CellsTouchingSegment calls touch for every rectangular cell that
// the line segment (start-end) overlaps, stopping at end or when touch
// returns false.
func CellsTouchingSegment(cellSize, start, end I2, touch func(cell I2) bool) bool {
	p, q := cell(start, cellSize), cell(end, cellSize)
	s := q.Sub(p).Sgn()
	// Special case: axis-aligned movement.
	if s.X == 0 || s.Y == 0 {
		for {
			if !touch(p) {
				return false
			}
			if p == q {
				return true
			}
			p = p.Add(s)
		}
	}
	// General case.
	v := end.Sub(start).EMul(s).F2()
	d := cellSize.F2()
	d.X, d.Y = d.X/v.X, d.Y/v.Y
	//t := p.Add(s).EMul(cellSize).Sub(start).EMul(s).F2()
	t := F2{}
	if s.X > 0 {
		t.X = float64((p.X+1)*cellSize.X - start.X)
	} else {
		t.X = float64(start.X - p.X*cellSize.X)
	}
	if s.Y > 0 {
		t.Y = float64((p.Y+1)*cellSize.Y - start.Y)
	} else {
		t.Y = float64(start.Y - p.Y*cellSize.Y)
	}
	t.X, t.Y = t.X/v.X, t.Y/v.Y
	//log.Printf("CellsTouchingSegment general case: v, t, d = %v, %v, %v", v, t, d)
	for {
		if !touch(p) {
			return false
		}
		if p == q {
			return true
		}
		if t.X > 1 && t.Y > 1 {
			//log.Printf("CellsTouchingSegment general case: got to p = %v (q = %v) but t > 1 (t = %v)", p, q, t)
			return true
		}
		if t.X < t.Y {
			t.X += d.X
			p.X += s.X
		} else {
			t.Y += d.Y
			p.Y += s.Y
		}
		//log.Printf("CellsTouchingSegment general case: p, t = %v, %v", p, t)
	}
}
