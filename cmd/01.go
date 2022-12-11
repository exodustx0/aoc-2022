package cmd

import (
	"bufio"
	"strconv"
)

func day01(input *bufio.Reader) (partOne, partTwo any) {
	s := bufio.NewScanner(input)

	var calories uint64
	var elfs []uint64
	for s.Scan() {
		str := s.Text()
		if str == "" {
			elfs = append(elfs, calories)
			calories = 0
			continue
		}

		c, err := strconv.ParseUint(str, 10, 64)
		if err != nil {
			panic(err)
		}

		calories += c
	}

	if calories != 0 {
		elfs = append(elfs, calories)
	}

	most := maxSlice(elfs, 3)
	partOne = most[0]
	partTwo = most[0] + most[1] + most[2]
	return
}
