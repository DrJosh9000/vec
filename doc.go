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

// Package vec is a simplistic, non-optimised 2D vector math library and
// some routines for segment intersections and graphs in the plane.
//
// It has two basic types, F2 and I2, representing points in the plane.
// The components of F2 are float64, and the components of I2 are int.
// Functions for I2 use int operations and avoid division where possible,
// though sometimes it is necessary to cast into int64 to avoid overflow.
package vec
