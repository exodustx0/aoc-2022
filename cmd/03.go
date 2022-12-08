package cmd

import "bufio"

func day03(input *bufio.Reader) (partOne, partTwo any) {
	var rucksacks [][2]map[byte]uint

	s := bufio.NewScanner(input)
	for s.Scan() {
		str := s.Text()
		rucksack := [2]map[byte]uint{make(map[byte]uint), make(map[byte]uint)}
		for i := range rucksack {
			items := str[len(str)/2*i : len(str)/2*(i+1)]
			for ii := 0; ii < len(items); ii++ {
				rucksack[i][items[ii]]++
			}
		}
		rucksacks = append(rucksacks, rucksack)
	}

	priority := func(i byte) uint {
		if i&0x20 != 0 {
			return uint(i) - 'a' + 1
		} else {
			return uint(i) - 'A' + 27
		}
	}

	var prioritySum uint
	for _, r := range rucksacks {
		for i := range r[0] {
			if r[1][i] != 0 {
				prioritySum += priority(i)
			}
		}
	}

	partOne = prioritySum

	prioritySum = 0
groupLoop:
	for g, n := 0, len(rucksacks)/3; g < n; g++ {
		for _, rc := range rucksacks[g*3] {
			for i := range rc {
				if rucksacks[g*3+1][0][i]+rucksacks[g*3+1][1][i] == 0 {
					continue
				}
				if rucksacks[g*3+2][0][i]+rucksacks[g*3+2][1][i] == 0 {
					continue
				}

				prioritySum += priority(i)
				continue groupLoop
			}
		}
	}

	partTwo = prioritySum
	return
}
