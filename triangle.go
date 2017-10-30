package visvalingam

import (
	"github.com/intdxdt/geom"
	"github.com/intdxdt/math"
)

type Triangle struct {
	a, b, c *Pt
	prev    *Triangle
	next    *Triangle
}

func NewTriangle(pts []*Pt) *Triangle {
	var a, b, c = pts[0], pts[1], pts[2]
	b.area = Area(a.Point, b.Point, c.Point)
	return &Triangle{a: a, b: b, c: c, prev: nil, next: nil}
}

func TriangleAreaCompare(t, o interface{}) int {
	var self  = t.(*Triangle)
	var other = o.(*Triangle)
	dx := float64(self.b.area - other.b.area)
	if math.FloatEqual(dx, 0.0) {
		return 0
	} else if dx < 0 {
		return -1
	}
	return 1
}

func (t *Triangle) String() string {
	return geom.NewPolygon([]*geom.Point{t.a.Point, t.b.Point, t.c.Point}).WKT()
}
