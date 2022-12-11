package cmd

import "golang.org/x/exp/constraints"

func abs[T constraints.Signed](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func min[T constraints.Ordered](x, y T) T {
	if y < x {
		return y
	}
	return x
}

func max[T constraints.Ordered](x, y T) T {
	if y > x {
		return y
	}
	return x
}

func maxSlice[T constraints.Ordered](x []T, n int) []T {
	most := make([]T, n)
	for _, v := range x {
		for i, m := range most {
			if v > m {
				if i != len(most)-1 {
					copy(most[i+1:], most[i:])
				}
				most[i] = v
				break
			}
		}
	}
	return most
}

func productSlice[T constraints.Integer | constraints.Float | constraints.Complex](x []T) T {
	product := T(1)
	for _, v := range x {
		product *= v
	}
	return product
}
