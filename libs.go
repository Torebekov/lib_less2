package lib_less2

import (
	"fmt"
	"github.com/muyo/sno"
	"math"
)

func ChangeCase(s string) string {
	res := []rune(s)
	for i, v := range res {
		switch {
		case v >= 'a' && v <= 'z':
			res[i] -= 32
		case v >= 'A' && v <= 'Z':
			res[i] += 32
		}
	}
	return string(res)
}

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

func GetUUID(x byte) string {
	return fmt.Sprintf("%v", sno.New(x))
}
