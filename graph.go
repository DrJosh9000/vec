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
	"errors"
	"math"
)

// Length returns the length of u-v.
func Length(u, v I2) float64 {
	v = v.Sub(u)
	return math.Sqrt(float64(v.Dot(v)))
}

// Edge describes an edge between two I2 vertices.
type Edge struct {
	U, V I2
}

// Length is the length of the edge.
func (e Edge) Length() float64 {
	v := e.V.Sub(e.U)
	return math.Sqrt(float64(v.Dot(v)))
}

// Reverse is the edge flipped (V->U).
func (e Edge) Reverse() Edge {
	return Edge{e.V, e.U}
}

// VertexSet is a set of vertices.
type VertexSet map[I2]bool

// Graph is an adjacency-set implementation of a graph.
type Graph struct {
	V VertexSet        // vertices
	E map[I2]VertexSet // edges; E[u] = {v: u-v is an edge}
}

// NewGraph creates a new empty *Graph.
func NewGraph() *Graph { return &Graph{} }

// AddEdge adds the edge u-v to the graph.
func (g *Graph) AddEdge(u, v I2) {
	if g.V == nil {
		g.V = make(VertexSet)
	}
	if g.E == nil {
		g.E = make(map[I2]VertexSet)
	}
	g.V[u] = true
	g.V[v] = true
	if g.E[u] == nil {
		g.E[u] = make(VertexSet)
	}
	g.E[u][v] = true
}

// AllEdges runs a function for every edge.
func (g *Graph) AllEdges(f func(I2, I2) bool) bool {
	for u, l := range g.E {
		for v, y := range l {
			if !y {
				continue
			}
			if !f(u, v) {
				return false
			}
		}
	}
	return true
}

// AllEdgesFacing calls f for every edge that is somewhat facing p.
func (g *Graph) AllEdgesFacing(p I2, f func(I2, I2) bool) bool {
	return g.AllEdges(func(u, v I2) bool {
		if SignedArea2(p, u, v) > 0 {
			return f(u, v)
		}
		return true
	})
}

// Edges returns all the edges as a slice.
func (g *Graph) Edges() (edges []Edge) {
	g.AllEdges(func(u, v I2) bool {
		edges = append(edges, Edge{u, v})
		return true
	})
	return
}

// NumEdges counts the number of edges.
func (g *Graph) NumEdges() (n int) {
	for _, l := range g.E {
		n += len(l)
	}
	return
}

// Blocks determines if the graph intersects the straight line segment start-end.
// Only edges facing start are considered.
func (g *Graph) Blocks(start, end I2) bool {
	return !g.AllEdgesFacing(start, func(u, v I2) bool {
		_, y := SegmentIntersectI(u, v, start, end)
		return !y
	})
}

// FullyBlocks determines if the graph intersects the straight line segment
// start-end, including back-facing edges.
func (g *Graph) FullyBlocks(start, end I2) bool {
	return !g.AllEdges(func(u, v I2) bool {
		_, y := SegmentIntersectI(u, v, start, end)
		return !y
	})
}

// NearestBlock returns the intersection with the graph nearest to start.
func (g *Graph) NearestBlock(start, end I2) (I2, bool) {
	found := false
	min := math.Inf(1)
	var pos I2
	g.AllEdgesFacing(start, func(u, v I2) bool {
		p, y := SegmentIntersectI(u, v, start, end)
		if !y {
			return true
		}
		if t := Length(start, p); t < min {
			min = t
			pos = p
			found = true
		}
		return true
	})
	return pos, found
}

// NearestPoint finds the edge, and closest point along that edge, to the query point.
func (g *Graph) NearestPoint(p I2) (e Edge, q I2) {
	d := int64(1<<63 - 1)
	g.AllEdges(func(u, v I2) bool {
		if r, t := SegmentNearestPoint(u, v, p); t < d {
			d = t
			e = Edge{u, v}
			q = r
		}
		return true
	})
	return
}

// FindPath finds a path from start to end, even if start and end are not vertices.
// The path will only use vertices contained in the rectangle boundsUL-boundsDR.
func FindPath(obstacles, paths *Graph, start, end, boundsUL, boundsDR I2) ([]I2, error) {
	// Is there a straight-line path?
	if !obstacles.Blocks(start, end) {
		return []I2{end}, nil
	}

	// Which vertices are start and end linked to?
	// Lengths to the vertices visible from start are optimal, because
	// the space is Euclidean.
	dists := make(map[I2]float64)
	prev := make(map[I2]I2)
	endN := make(VertexSet)
	q := VertexSet{end: true}
	for v, y := range paths.V {
		if !y || !v.InRect(boundsUL, boundsDR) {
			continue
		}
		q[v] = true
		dists[v] = math.Inf(1)
		if !obstacles.Blocks(start, v) {
			dists[v] = Length(start, v)
			prev[v] = start
		}
		if !obstacles.Blocks(v, end) {
			endN[v] = true
		}
	}
	if len(prev) == 0 || len(endN) == 0 {
		return nil, errors.New("no paths possible")
	}

	// Dijkstra time.
	dists[start] = 0
	dists[end] = math.Inf(1)

	relax := func(u, v I2) {
		if !q[v] {
			return
		}
		if t := Length(u, v) + dists[u]; t < dists[v] {
			dists[v] = t
			prev[v] = u
		}
	}

	for len(q) > 0 {
		dist := math.Inf(1)
		var u I2
		for v := range q {
			if dists[v] < dist {
				dist = dists[v]
				u = v
			}
		}
		if u == end || math.IsInf(dist, 1) {
			break
		}
		delete(q, u)

		for v := range paths.E[u] {
			relax(u, v)
		}
		for v := range endN {
			relax(v, end)
		}
	}

	if _, ok := prev[end]; !ok {
		return nil, errors.New("no path")
	}

	// Traverse the optimal path and then put it in the right order.
	v := end
	var rpath []I2
	for v != start {
		rpath = append(rpath, v)
		v = prev[v]
	}
	path := make([]I2, len(rpath))
	for i := 0; i < len(rpath); i++ {
		path[i] = rpath[len(rpath)-i-1]
	}
	return path, nil
}
