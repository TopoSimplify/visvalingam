package visvalingam

import (
	"simplex/geom"
	"simplex/struct/heap"
)

//type visvalingam
type Visvalingam struct {
	coords     []*Pt
	triangles  []*Triangle
	heap       *heap.Heap
	buildState bool
}

//new visvalingam
func NewVisvalingam(coords []*geom.Point) *Visvalingam {
	pts := make([]*Pt, len(coords))
	for i, coord := range coords {
		pts[i] = &Pt{Point: coord.Clone(), area: 0.0}
	}
	return &Visvalingam{coords: pts}
}

//build
func (vis *Visvalingam) build() *Visvalingam {
	vis.heap = heap.NewHeap(heap.NewHeapType().AsMin())
	vis.createTriangles().updateTrianglePtrs()
	vis.buildState = true
	return vis
}

func (vis *Visvalingam) Simplify(threshold float64) []*geom.Point {
	if !vis.buildState {
		vis.build()
	}

	var maxArea float64
	for {
		triangle := vis.heap.Pop()
		if triangle == nil {
			break
		}
		t := triangle.(*Triangle)

		// If the area of the current point is less than that of the previous point
		// to be eliminated, use the latter’s area instead. This ensures that the
		// current point cannot be eliminated without eliminating previously-
		// eliminated points.
		if t.b.area < maxArea {
			t.b.area = maxArea
		} else {
			maxArea = t.b.area
		}

		if t.prev != nil {
			t.prev.next = t.next
			t.prev.c = t.c
			vis.update(t.prev)
		} else {
			t.a.area = t.b.area
		}

		if t.next != nil {
			t.next.prev = t.prev
			t.next.a = t.a
			vis.update(t.next)
		} else {
			t.c.area = t.b.area
		}

	}

	simplx := make([]*geom.Point, 0)
	for _, pt := range vis.coords {
		if pt.area >= threshold {
			simplx = append(simplx, pt.Point)
		}
	}
	return simplx
}

//create triangles
func (vis *Visvalingam) createTriangles() *Visvalingam {
	vis.triangles = make([]*Triangle, 0)
	for i, n := 1, len(vis.coords)-1; i < n; i++ {
		coords := vis.coords[i-1: i+2]
		t := NewTriangle(coords)

		if t.b.area > 0 {
			vis.triangles = append(vis.triangles, t)
			vis.heap.Push(t)
		}
	}
	return vis
}

//update triangles pointers
func (vis *Visvalingam) updateTrianglePtrs() *Visvalingam {
	for i, n := 0, len(vis.triangles); i < n; i++ {
		t := vis.triangles[i]
		t.next, t.prev = nil, nil

		if i-1 >= 0 {
			t.prev = vis.triangles[i-1]
		}

		if i+1 < n {
			t.next = vis.triangles[i+1]
		}
	}
	return vis
}

func (vis *Visvalingam) update(t *Triangle) {
	vis.heap.Remove(t)
	t.b.area = Area(t.a.Point, t.b.Point, t.c.Point)
	vis.heap.Push(t)
}
