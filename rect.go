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

type Rect struct {
	UL, DR I2
}

func NewRect(x0, y0, x1, y1 int) Rect {
	return Rect{UL: I2{X: x0, Y: y0}, DR: I2{X: x1, Y: y1}}
}

func (r Rect) C() (x0, y0, x1, y1 int) {
	return r.UL.X, r.UL.Y, r.DR.X, r.DR.Y
}

func (r Rect) Contains(p I2) bool {
	return p.X >= r.UL.X && p.X < r.DR.X && p.Y >= r.UL.Y && p.Y < r.DR.Y
}

func (r Rect) Overlaps(s Rect) bool {
	return s.DR.X > r.UL.X && s.UL.X < r.DR.X && s.DR.Y > r.UL.Y && s.UL.Y < r.DR.Y
}

func (r Rect) Translate(p I2) Rect {
	return Rect{UL: r.UL.Add(p), DR: r.DR.Add(p)}
}

func (r Rect) Size() I2 {
	return r.DR.Sub(r.UL)
}

func (r Rect) Reposition(ul I2) Rect {
	return Rect{UL: ul, DR: ul.Add(r.Size())}
}

func (r Rect) Resize(sz I2) Rect {
	return Rect{UL: r.UL, DR: r.UL.Add(sz)}
}
