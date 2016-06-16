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

// I2 is a pair of integers, (X,Y).
type I2 struct{ X, Y int }

// NewI2 is a convenience function for creating an I2.
func NewI2(x, y int) I2 { return I2{x, y} }

// Div returns (n%d, n/d) as an I2.
func Div(n, d int) I2 { return I2{n % d, n / d} }

// C returns the components of v (X,Y).
func (v I2) C() (int, int) { return v.X, v.Y }

// C64 returns the components of v (X,Y) as int64s.
func (v I2) C64() (int64, int64) { return int64(v.X), int64(v.Y) }

// F2 casts this I2 to an F2, which may potentially cause loss of precision.
func (v I2) F2() F2 { return F2{float64(v.X), float64(v.Y)} }

// Add returns v + w.
func (v I2) Add(w I2) I2 { return I2{v.X + w.X, v.Y + w.Y} }

// Sub returns v - w.
func (v I2) Sub(w I2) I2 { return I2{v.X - w.X, v.Y - w.Y} }

// Mul returns the scalar product k * v.
func (v I2) Mul(k int) I2 { return I2{v.X * k, v.Y * k} }

// Div integer-divides both components by k.
func (v I2) Div(k int) I2 { return I2{v.X / k, v.Y / k} }

// MulDiv does both scalar multiplication and division, and tries to avoid overflow on 32-bit.
func (v I2) MulDiv(n, d int64) I2 {
	vx, vy := v.C64()
	vx *= n
	vy *= n
	vx /= d
	vy /= d
	return I2{int(vx), int(vy)}
}

// EMul returns the element-wise product of v and w.
func (v I2) EMul(w I2) I2 { return I2{v.X * w.X, v.Y * w.Y} }

// EDiv returns the element-wise quotient of v and w.
func (v I2) EDiv(w I2) I2 { return I2{v.X / w.X, v.Y / w.Y} }

// Mod returns the remainder of both components divided by k.
func (v I2) Mod(k int) I2 { return I2{v.X % k, v.Y % k} }

// Sgn returns a "unit-ish" vector (each component is normalised).
func (v I2) Sgn() I2 { return I2{Sgn(v.X), Sgn(v.Y)} }

// Area returns the product of X and Y.
func (v I2) Area() int { return v.X * v.Y }

// ClampLo returns v, but with components clamped below by components of e.
func (v I2) ClampLo(e I2) I2 {
	if v.X < e.X {
		v.X = e.X
	}
	if v.Y < e.Y {
		v.Y = e.Y
	}
	return v
}

// ClampHi returns v, but with components clamped above by components of e.
func (v I2) ClampHi(e I2) I2 {
	if v.X >= e.X {
		v.X = e.X
	}
	if v.Y >= e.Y {
		v.Y = e.Y
	}
	return v
}

// Dot returns the dot product, v dot w.
func (v I2) Dot(w I2) int64 { return int64(v.X)*int64(w.X) + int64(v.Y)*int64(w.Y) }

// Normal returns a vector perpendicular to v of the same length.
func (v I2) Normal() I2 { return I2{-v.Y, v.X} }

// Cmul returns the complex product v * w, where the X components are treated as real
// and the Y components as imaginary.
func (v I2) Cmul(w I2) I2 { return I2{w.X*v.X - w.Y*v.Y, w.Y*v.X + w.X*v.Y} }

// Swap switches x and y components.
func (v I2) Swap() I2 { return I2{v.Y, v.X} }

// InRect tests if v is in the rectangle ul-dr.
func (v I2) InRect(ul, dr I2) bool {
	return v.X >= ul.X && v.X <= dr.X && v.Y >= ul.Y && v.Y <= dr.Y
}

// LineIntersectI finds the intersection of the lines (infinite) through p,q and a,b,
// or returns false if they are parallel.
func LineIntersectI(p, q, a, b I2) (I2, bool) {
	dx1, dy1, dx2, dy2 := p.X-q.X, p.Y-q.Y, a.X-b.X, a.Y-b.Y
	det := dx2*dy1 - dx1*dy2
	if det == 0 {
		return I2{}, false
	}
	pq := (q.X*p.Y - p.X*q.Y)
	ab := (b.X*a.Y - a.X*b.Y)
	return I2{pq*dx2 - ab*dx1, pq*dy2 - ab*dy1}.Div(det), true
}

// Abs is the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Sgn is the sign of X (-1, 0, or 1).
func Sgn(x int) int {
	switch {
	case x < 0:
		return -1
	case x > 0:
		return 1
	default:
		return 0
	}
}

// SegmentIntersectI tests for the intersection of the line segments p-q, a-b. If
// there is an intersection an approximation of the point of intersection will
// be returned.
func SegmentIntersectI(p, q, a, b I2) (I2, bool) {
	dx1, dy1, dx2, dy2 := p.X-q.X, p.Y-q.Y, a.X-b.X, a.Y-b.Y
	det := dy2*dx1 - dx2*dy1
	if det == 0 {
		return I2{}, false
	}
	apy := a.Y - p.Y
	t := (a.X*(b.Y-p.Y) - b.X*apy + p.X*dy2) * Sgn(det)
	if t < 0 || t > Abs(det) {
		return I2{}, false
	}
	t = (p.X*(a.Y-q.Y) - q.X*apy - a.X*dy1) * Sgn(det)
	if t < 0 || t > Abs(det) {
		return I2{}, false
	}
	return b.Sub(a).MulDiv(int64(t), int64(Abs(det))).Add(a), true
}

// SignedArea2 returns double the signed area of the triangle abc.
func SignedArea2(a, b, c I2) int64 {
	ax, ay := a.C64()
	bx, by := b.C64()
	cx, cy := c.C64()
	return ay*(cx-bx) + by*(ax-cx) + cy*(bx-ax)
}

// LineNearestPoint locates the point on the line passing through uv that is closest to p,
// and returns the point and the square of the distance.
func LineNearestPoint(u, v, p I2) (I2, int64) {
	if u == v {
		p = p.Sub(u)
		return u, p.Dot(p)
	}
	n := v.Sub(u).Normal()
	// Enforce 64 bit here due to overflow.
	n2 := n.Dot(n)
	sa := SignedArea2(u, v, p)
	return p.Sub(n.MulDiv(sa, n2)), (sa * sa) / n2
}

// SegmentNearestPoint locates the point on the line segment uv that is closest to p,
// and returns the point and the square of the distance.
func SegmentNearestPoint(u, v, p I2) (I2, int64) {
	q, d := LineNearestPoint(u, v, p)
	if u == v {
		return q, d
	}
	// Already know q is somewhere on the line, so just do comparisons.
	qu, vu := q.Sub(u), v.Sub(u)
	qu.X *= Sgn(vu.X)
	qu.Y *= Sgn(vu.Y)
	if qu.X < 0 || qu.Y < 0 {
		pu := p.Sub(u)
		return u, pu.Dot(pu)
	}
	if qu.X > Abs(vu.X) || qu.Y > Abs(vu.Y) {
		vp := v.Sub(p)
		return v, vp.Dot(vp)
	}
	return q, d
}

// RectRange makes a list of integer points contained in the Cartesian product [ul.X, dr.X) * [ul.Y, dr.Y).
func RectRange(ul, dr I2) []I2 {
	r := make([]I2, 0, (dr.X-ul.X)*(dr.Y-ul.Y))
	for x := ul.X; x < dr.X; x++ {
		for y := ul.Y; y < dr.Y; y++ {
			r = append(r, I2{x, y})
		}
	}
	return r
}
