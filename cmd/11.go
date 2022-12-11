package cmd

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

var monkeys []monkey

type monkey struct {
	items        []int
	inspectCount int
	divisor      int
	true, false  int
	op           func(int) int
	rough        bool
}

func (m *monkey) doTurn() {
	for len(m.items) != 0 {
		m.inspect()
	}
}

func (m *monkey) inspect() {
	m.inspectCount++

	item := m.items[0]
	m.items = m.items[1:]
	item = m.op(item)
	if !m.rough {
		item /= 3
	}

	var o *monkey
	if item%m.divisor == 0 {
		o = &monkeys[m.true]
	} else {
		o = &monkeys[m.false]
	}
	o.items = append(o.items, item)
}

func day11(input *bufio.Reader) (partOne, partTwo any) {
	commonModulus := 1
	s := bufio.NewScanner(input)
	for s.Scan() {
		var m monkey

		s.Scan()
		for _, itemStr := range strings.Split(strings.TrimPrefix(s.Text(), "  Starting items: "), ", ") {
			item, err := strconv.ParseUint(itemStr, 10, 0)
			if err != nil {
				panic(err)
			}
			m.items = append(m.items, int(item))
		}

		s.Scan()
		var operator, rhsStr string
		if n, err := fmt.Sscanf(s.Text(), "  Operation: new = old %s %s", &operator, &rhsStr); n != 2 || err != nil {
			panic(err)
		}

		var rhs int
		if rhsStr != "old" {
			rhsVal, err := strconv.ParseUint(rhsStr, 10, 0)
			if err != nil {
				panic(err)
			}
			rhs = int(rhsVal)
		}

		switch operator {
		case "+":
			if rhsStr == "old" {
				m.op = func(i int) int { return (i << 1) % commonModulus }
			} else {
				m.op = func(i int) int { return (i + rhs) % commonModulus }
			}
		case "*":
			if rhsStr == "old" {
				m.op = func(i int) int { return (i * i) % commonModulus }
			} else {
				m.op = func(i int) int { return (i * rhs) % commonModulus }
			}
		default:
			panic("unexpected operator")
		}

		s.Scan()
		if n, err := fmt.Sscanf(s.Text(), "  Test: divisible by %d", &m.divisor); n != 1 || err != nil {
			panic(err)
		}
		commonModulus *= m.divisor

		s.Scan()
		if n, err := fmt.Sscanf(s.Text(), "    If true: throw to monkey %d", &m.true); n != 1 || err != nil {
			panic(err)
		}

		s.Scan()
		if n, err := fmt.Sscanf(s.Text(), "    If false: throw to monkey %d", &m.false); n != 1 || err != nil {
			panic(err)
		}

		s.Scan()
		monkeys = append(monkeys, m)
	}

	monkeysCopy := make([]monkey, len(monkeys))
	copy(monkeysCopy, monkeys)

	for round := 0; round < 20; round++ {
		for i := range monkeys {
			monkeys[i].doTurn()
		}
	}

	inspectCounts := make([]int, len(monkeys))
	for i, m := range monkeys {
		inspectCounts[i] = m.inspectCount
	}

	partOne = productSlice(maxSlice(inspectCounts, 2))

	monkeys = monkeysCopy
	for i := range monkeys {
		monkeys[i].rough = true
	}

	for round := 0; round < 10000; round++ {
		for i := range monkeys {
			monkeys[i].doTurn()
		}
	}

	for i, m := range monkeys {
		inspectCounts[i] = m.inspectCount
	}

	partTwo = productSlice(maxSlice(inspectCounts, 2))
	return
}
