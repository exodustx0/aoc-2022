package cmd

import (
	"bufio"
	"strconv"
	"strings"
)

func day10(input *bufio.Reader) (partOne, partTwo any) {
	var cycle, signalStrengthSum int
	var image string
	cyclesUntilCheck := 20
	x := 1
	clockStep := func() {
		if cyclesUntilCheck == 20 {
			image += "\n"
		}

		if abs(cycle%40-x) > 1 {
			image += "."
		} else {
			image += "#"
		}

		cycle++
		cyclesUntilCheck--
		if cyclesUntilCheck == 0 {
			cyclesUntilCheck = 40
			signalStrengthSum += cycle * x
		}
	}

	s := bufio.NewScanner(input)
	for s.Scan() {
		cmd, arg, ok := strings.Cut(s.Text(), " ")
		switch cmd {
		case "noop":
			clockStep()
		case "addx":
			if !ok {
				panic("missing addx argument")
			}

			add, err := strconv.Atoi(arg)
			if err != nil {
				panic(err)
			}

			clockStep()
			clockStep()
			x += add
		}
	}

	partOne = signalStrengthSum
	partTwo = image
	return
}
