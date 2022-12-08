package cmd

import (
	"bufio"
	"strconv"
	"strings"
)

func day04(input *bufio.Reader) (partOne, partTwo any) {
	getSectionRange := func(str string) [2]uint {
		startStr, endStr, ok := strings.Cut(str, "-")
		if !ok {
			panic("no range separator")
		}

		start, err := strconv.ParseUint(startStr, 10, 64)
		if err != nil {
			panic(err)
		}
		end, err := strconv.ParseUint(endStr, 10, 64)
		if err != nil {
			panic(err)
		}

		if start > end {
			panic("range not low-to-high")
		}

		return [2]uint{uint(start), uint(end)}
	}

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
		firstStr, secondStr, ok := strings.Cut(s.Text(), ",")
		if !ok {
			panic("no pair separator")
		}

		first, second := getSectionRange(firstStr), getSectionRange(secondStr)
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
