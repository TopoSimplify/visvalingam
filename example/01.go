package main

import (
	"simplex/visvalingam"
	"simplex/geom"
	"fmt"
)

func main() {
	var line = []*geom.Point{{3.0, 1.6}, {3.0, 2.0}, {2.4, 2.8}, {0.5, 3.0}, {1.2, 3.2}, {1.4, 2.6}, {2.0, 3.5}}
	var visva = visvalingam.NewVisvalingam(line)
	var res = 0.25

	var simplx = visva.Simplify(res)

	fmt.Println(geom.NewLineString(line).WKT())
	fmt.Println(geom.NewLineString(simplx).WKT())
}
