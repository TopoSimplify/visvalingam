package visvalingam

import (
	"simplex/geom"
	"math"
)

//triangle area
func  Area(a, b, c *geom.Point) float64 {
	return 0.5 * math.Abs((a[0]-c[0])*(b[1]-a[1])-(a[0]-b[0])*(c[1]-a[1]))
}
