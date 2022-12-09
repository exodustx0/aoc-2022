package cmd

import (
	"bufio"
	"fmt"
)

func day04(input *bufio.Reader) (partOne, partTwo any) {
	contains := func(first, second [2]uint) bool {
		return first[0] >= second[0] && first[1] <= second[1]
	}

	overlaps := func(point uint, sRange [2]uint) bool {
		return point >= sRange[0] && point <= sRange[1]
	}

	var redundancyCount uint
	var overlapCount uint
	s := bufio.NewScanner(input)
	for s.Scan() {
		var start1, end1, start2, end2 uint
		if n, err := fmt.Sscanf(s.Text(), "%d-%d,%d-%d", &start1, &end1, &start2, &end2); n != 4 || err != nil {
			panic(err)
		}

		first, second := [2]uint{start1, end1}, [2]uint{start2, end2}
		if contains(first, second) || contains(second, first) {
			redundancyCount++
			overlapCount++
			continue
		}

		if overlaps(first[0], second) || overlaps(first[1], second) {
			overlapCount++
		}
	}

	partOne = redundancyCount
	partTwo = overlapCount
	return
}
