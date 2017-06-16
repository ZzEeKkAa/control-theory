package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/ZzEeKkAa/numeric-methods/lab3/mid-rec"
)

func main() {
	rand.Seed(time.Now().Unix())
	var s mid_rec.Solver

	a1 := rand.Float64()
	a2 := rand.Float64()
	b1 := rand.Float64()
	b2 := rand.Float64()
	n1 := rand.Float64()
	n2 := rand.Float64()
	h1 := rand.Float64()
	h2 := rand.Float64()
	t := rand.Float64()
	_ns := 0.07
	_hs := 0.2
	it := 100

	calc(a1, a2, b1, b2, h1, h2, n1, n2, t)

	var l Points

	for _h1 := -h1; _h1 < h1+_hs; _h1 += _hs {
		if _h1 > h1 {
			_h1 = h1
		}
		for _h2 := -h2; _h2 < h2+_hs; _h2 += _hs {
			if _h2 > h2 {
				_h2 = h2
			}
			s.F(func(i int, x float64, y ...float64) float64 {
				switch i {
				case 0:
					return -a1*y[0] - b2*y[1] + _h1
				case 1:
					return -b1*y[0] - a2*y[1] + _h2
				}
				return 0
			}, 2)
			for _n1 := 0.; _n1 < n1+_ns; _n1 += _ns {
				if _n1 > n1 {
					_n1 = n1
				}
				for _n2 := 0.; _n2 < n2+_ns; _n2 += _ns {
					if _n2 > n2 {
						_n2 = n2
					}
					sol := s.Solve(t, it, 0, _n1, _n2)
					l = append(l, Point{sol[0], sol[1]})
				}
			}
		}
	}
	fmt.Println("p:=", l, ":;")

	//calc(1, 10, 20, 10, 10, 20, 1, 2, 0.43)
}

func calc(a1, a2, b1, b2, h1, h2, n1, n2, t float64) {
	fmt.Printf("a[1]:=%.3f;a[2]:=%.3f;b[1]:=%.3f;b[2]:=%.3f;h[1]:=%.3f;h[1]:=%.3f;N[1]:=%.3f;N[1]:=%.3f;T:=%.3f;\n", a1, a2, b1, b2, h1, h2, n1, n2, t)
	k := math.Sqrt(a1*a1 - 2*a1*a2 + a2*a2 + 4*b1*b2)
	q1 := 0.5 * (k + a1 + a2)
	q2 := 0.5 * (k - a1 - a2)
	p1 := 0.5 * (k - a1 + a2)
	p2 := 0.5 * (k + a1 - a2)

	a11 := (p1*math.Exp(q2*t) + p2*math.Exp(-q1*t)) / k
	a12 := b2 * (-math.Exp(q2*t) + math.Exp(-q1*t)) / k
	a21 := -p1 * p2 * (math.Exp(q2*t) - math.Exp(-q1*t)) / (k * b2)
	a22 := (p2*math.Exp(q2*t) + p1*math.Exp(-q1*t)) / k

	b11 := (p1*math.Exp(q2*t)*q1 - p2*math.Exp(-q1*t)*q2 - p1*q1 + p2*q2) / (k * q2 * q1)
	b12 := -b2 * (math.Exp(q2*t)*q1 + math.Exp(-q1*t)*q2 - q1 - q2) / (k * q2 * q1)
	b21 := -p1 * p2 * (math.Exp(q2*t)*q1 + math.Exp(-q1*t)*q2 - q1 - q2) / (b2 * k * q2 * q1)
	b22 := (p2*math.Exp(q2*t)*q1 - p1*math.Exp(-q1*t)*q2 + p1*q2 - p2*q1) / (k * q2 * q1)

	var l1 Points
	l1 = append(l1,
		Point{0, 0},
		Point{n1*a11 + n2*a12, n1*a21 + n2*a22},
		Point{n1 * a11, n1 * a21},
		Point{n2 * a12, n2 * a22},
	)
	var l2 Points
	l2 = append(l2,
		Point{h1*b11 + h2*b12, h1*b21 + h2*b22},
		Point{-h1*b11 + h2*b12, -h1*b21 + h2*b22},
		Point{h1*b11 - h2*b12, h1*b21 - h2*b22},
		Point{-h1*b11 - h2*b12, -h1*b21 - h2*b22},
	)
	var l Points
	for _, p1 := range l1 {
		for _, p2 := range l2 {
			l = append(l, p1.Add(p2))
		}
	}

	l = findConvexHull(l)

	fmt.Println("l:=", l, ":;")
}
