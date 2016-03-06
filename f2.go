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

import "math"

// Epsilon is a small quantity used for floating point comparisons.
// TODO: do it a better way.
const Epsilon = 0.0000000001

// F2 is a tuple of float64 numbers, (X,Y).
type F2 struct{ X, Y float64 }

// NewF2 is a convenience function for creating an F2.
func NewF2(x, y float64) F2 { return F2{x, y} }

// Unit returns the unit vector at angle t.
func Unit(t float64) F2 { return F2{math.Cos(t), math.Sin(t)} }

// C returns the components of v (X,Y).
func (v F2) C() (float64, float64) { return v.X, v.Y }

// I2 rounds this F2 to an I2, which will generally result in loss of precision.
func (v F2) I2() I2 {
	return I2{int(v.X + 0.5), int(v.Y + 0.5)}
}

// Add returns v + w.
func (v F2) Add(w F2) F2 { return F2{v.X + w.X, v.Y + w.Y} }

// Sub returns v - w.
func (v F2) Sub(w F2) F2 { return F2{v.X - w.X, v.Y - w.Y} }

// Mul returns the scalar product k * v.
func (v F2) Mul(k float64) F2 { return F2{v.X * k, v.Y * k} }

// Div returns the componentwise division by k.
func (v F2) Div(k float64) F2 { return F2{v.X / k, v.Y / k} }

// Dot returns the dot product, v dot w.
func (v F2) Dot(w F2) float64 { return v.X*w.X + v.Y*w.Y }

// Norm returns the length of v (the square root of v dot v).
func (v F2) Norm() float64 { return math.Sqrt(v.Dot(v)) }

// Unit returns the unit vector pointing in the same direction as v.
func (v F2) Unit() F2 { return v.Mul(1 / v.Norm()) }

// Normal returns a vector perpendicular to v of the same length.
func (v F2) Normal() F2 { return F2{-v.Y, v.X} }

// Arg returns the angle between the X-axis and the vector.
func (v F2) Arg() float64 { return math.Atan2(v.Y, v.X) }

// Cmul returns the complex product v * w, where the X components are treated as real
// and the Y components as imaginary.
func (v F2) Cmul(w F2) F2 { return F2{w.X*v.X - w.Y*v.Y, w.Y*v.X + w.X*v.Y} }

// Rot rotates the vector by the angle t.
func (v F2) Rot(t float64) F2 { return v.Cmul(Unit(t)) }

// RotAbout rotates the vector by the angle t around the vector b.
func (v F2) RotAbout(t float64, b F2) F2 { return v.Sub(b).Rot(t).Add(b) }

// Dir returns the general direction of v (Up, Down, Left, Right).
func (v F2) Dir() Direction {
	switch {
	case v.X >= v.Y && v.X >= -v.Y:
		return Right
	case v.Y > v.X && v.Y > -v.X:
		return Down
	case v.Y < v.X && v.Y < -v.X:
		return Up
	default:
		return Left
	}
}

// InRect tests if v is in the rectangle with topleft corner (x0, y0) and bottomright corner (x1, y1).
func (v F2) InRect(x0, y0, x1, y1 float64) bool {
	return v.X >= x0 && v.X <= x1 && v.Y >= y0 && v.Y <= y1
}

// LineIntersect finds the intersection of the lines (infinite) through p,q and a,b,
// or returns false if they are parallel.
func LineIntersect(p, q, a, b F2) (F2, bool) {
	dx1, dx2, dy1, dy2 := p.X-q.X, a.X-b.X, p.Y-q.Y, a.Y-b.Y
	det := dx2*dy1 - dx1*dy2
	if -Epsilon < det && det < Epsilon {
		return F2{}, false
	}
	pq := (q.X*p.Y - p.X*q.Y)
	ab := (b.X*a.Y - a.X*b.Y)
	return F2{pq*dx2 - ab*dx1, pq*dy2 - ab*dy1}.Div(det), true
}

// SegmentIntersect tests for the intersection of the line segments p-q, a-b;
// if there is an intersection it returns how far along a-b the
// intersection occurs.
func SegmentIntersect(p, q, a, b F2) (float64, bool) {
	qmpn, bma := q.Sub(p).Normal(), b.Sub(a)
	det := bma.Dot(qmpn)
	if -Epsilon < det && det < Epsilon { // Singularish
		return 0, false
	}
	// Is our result (t,t') is within bounds?
	pma := p.Sub(a)
	t := pma.Dot(bma.Normal()) / det
	if !(0 <= t && t < 1) {
		return 0, false
	}
	t = pma.Dot(qmpn) / det
	return t, 0 <= t && t < 1
}
