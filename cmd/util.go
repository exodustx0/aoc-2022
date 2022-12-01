package cmd

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

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

func partOne(a any) { fmt.Println("Part one:", a) }
func partTwo(a any) { fmt.Println("Part two:", a) }
