package cmd

import (
	"bufio"
	"strconv"
)

func day01(input *bufio.Reader) error {
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
			return err
		}

		calories += c
	}

	if calories != 0 {
		elfs = append(elfs, calories)
	}

	var most [3]uint64
	for _, c := range elfs {
		for i, m := range most {
			if c > m {
				if i != len(most)-1 {
					copy(most[i+1:], most[i:])
				}
				most[i] = c
				break
			}
		}
	}

	partOne(most[0])
	partTwo(most[0] + most[1] + most[2])

	return nil
}
