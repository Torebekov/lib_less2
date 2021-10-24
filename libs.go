package lib_less2

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
