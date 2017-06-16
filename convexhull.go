package main

import (
	"fmt"
	"sort"
	"strings"
)

// Point is a struct that holds the X Y cooridinates
// of a specific point in the Ecliden plane or space.
type Point struct {
	X, Y float64
}

func (p Point) Add(p2 Point) Point {
	p.X += p2.X
	p.Y += p2.Y
	return p
}

// Points is a slice built up of Point structs.
type Points []Point

func (points Points) String() string {
	var sl []string
	for _, p := range points {
		sl = append(sl, fmt.Sprintf("[%.3f,%.3f]", p.X, p.Y))
	}
	return fmt.Sprintf("[%s]", strings.Join(sl, ","))
}

func (points Points) Swap(i, j int) {
	points[i], points[j] = points[j], points[i]
}

func (points Points) Len() int {
	return len(points)
}

// lets sort our Points by x and, if equal, by y
func (points Points) Less(i, j int) bool {
	if points[i].X == points[j].X {
		return points[i].Y < points[j].Y
	}
	return points[i].X < points[j].X
}

// returns the modulo (and sign) of the cross product between vetors OA and OB
func crossProduct(O, A, B Point) float64 {
	return (A.X-O.X)*(B.Y-O.Y) - (A.Y-O.Y)*(B.X-O.X)
}

// findConvexHull returns a slice of Point with a convex hull
// it is counterclockwise and starts and ends at the same point
// i.e. the same point is repeated at the beginning and at the end
func findConvexHull(points Points) Points {
	n := len(points)  // number of points to find convex hull
	var result Points // final result
	count := 0        // size of our convex hull (number of points added)

	// lets sort our points by x and if equal by y
	sort.Sort(points)

	if n == 0 {
		return result
	}

	// add the first element:
	result = append(result, points[0])
	count++

	// find the lower hull
	for i := 1; i < n; i++ {
		// remove points which are not part of the lower hull
		for count > 1 && crossProduct(result[count-2], result[count-1], points[i]) < 0 {
			count--
			result = result[:count]
		}

		// add a new better point than the removed ones
		result = append(result, points[i])
		count++
	}

	count0 := count // our base counter for the upper hull

	// find the upper hull
	for i := n - 2; i >= 0; i-- {
		// remove points which are not part of the upper hull
		for count-count0 > 0 && crossProduct(result[count-2], result[count-1], points[i]) < 0 {
			count--
			result = result[:count]
		}

		// add a new better point than the removed ones
		result = append(result, points[i])
		count++
	}

	return result
}
