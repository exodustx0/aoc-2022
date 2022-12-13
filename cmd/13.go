package cmd

import (
	"bufio"
	"sort"
	"strconv"
)

func day13(input *bufio.Reader) (partOne, partTwo any) {
	var parseList func(in string, i int) ([]any, int)
	parseList = func(in string, i int) ([]any, int) {
		var list []any
		for i < len(in) {
			switch in[i] {
			case ',':
				i++

			case '[':
				var subList []any
				subList, i = parseList(in, i+1)
				list = append(list, subList)

			case ']':
				return list, i + 1

			default:
				d := 0
				for in[i+d] >= '0' && in[i+d] <= '9' {
					d++
				}

				if d == 0 {
					panic("expected integer")
				}

				integer, err := strconv.ParseUint(in[i:i+d], 10, 0)
				if err != nil {
					panic(err)
				}

				list = append(list, int(integer))
				i += d
			}
		}

		panic("unexpected end of packet")
	}

	var packets [][]any
	s := bufio.NewScanner(input)
	for s.Scan() {
		p, _ := parseList(s.Text(), 1)
		packets = append(packets, p)
		if !s.Scan() {
			panic("unexpected end of input")
		}

		p, _ = parseList(s.Text(), 1)
		packets = append(packets, p)
		s.Scan()
	}

	var cmpPackets func(a, b any) int
	cmpPackets = func(a, b any) int {
		aInt, aIntOK := a.(int)
		aList, aListOK := a.([]any)
		bInt, bIntOK := b.(int)
		bList, bListOK := b.([]any)

		switch {
		case aIntOK && bIntOK:
			switch {
			case aInt < bInt:
				return -1
			case aInt > bInt:
				return 1
			default:
				return 0
			}

		default:
			if aIntOK {
				aList = []any{aInt}
			} else {
				bList = []any{bInt}
			}

			fallthrough
		case aListOK && bListOK:
			for i := 0; i < min(len(aList), len(bList)); i++ {
				if result := cmpPackets(aList[i], bList[i]); result != 0 {
					return result
				}
			}

			switch {
			case len(aList) < len(bList):
				return -1
			case len(aList) > len(bList):
				return 1
			default:
				return 0
			}
		}
	}

	correctOrderIndexSum := 0
	for i, n := 0, len(packets)/2; i < n; i++ {
		if cmpPackets(packets[i*2], packets[i*2+1]) != 1 {
			correctOrderIndexSum += i + 1
		}
	}

	partOne = correctOrderIndexSum

	two := []any{[]any{2}}
	six := []any{[]any{6}}
	packets = append(packets, two)
	packets = append(packets, six)

	sort.Slice(packets, func(i, j int) bool { return cmpPackets(packets[i], packets[j]) == -1 })

	key := 1
	i := 0
	for {
		i++
		if cmpPackets(packets[i-1], two) == 0 {
			key *= i
			break
		}
	}
	for {
		i++
		if cmpPackets(packets[i-1], six) == 0 {
			key *= i
			break
		}
	}

	partTwo = key
	return
}
