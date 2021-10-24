package lib_less2

import (
	"fmt"
	"math"
)

func FindRoots(a, b, c float64) []string {
	d := b*b - 4*a*c
	var res []string
	switch {
	case d > 0:
		x1 := (-b + math.Sqrt(d)) / (2 * a)
		x2 := (-b - math.Sqrt(d)) / (2 * a)
		res = append(res, fmt.Sprintf("%v", x1), fmt.Sprintf("%v", x2))
	case d == 0:
		x := -b / (2 * a)
		res = append(res, fmt.Sprintf("%v", x))
	case d < 0:
		res = append(res, "No real roots")
	}
	return res
}
