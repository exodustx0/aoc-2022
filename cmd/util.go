package cmd

import "golang.org/x/exp/constraints"

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

func abs[T constraints.Signed](x T) T {
	if x < 0 {
		return -x
	}
	return x
}
