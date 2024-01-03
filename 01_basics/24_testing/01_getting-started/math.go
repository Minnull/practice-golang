package math

func Adder(xs ...int) int {
	res := 0
	for _, x := range xs {
		res += x
	}
	return res
}
