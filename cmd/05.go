package cmd

import (
	"bufio"
	"fmt"
	"strings"
)

type crateStack []rune

func (cs *crateStack) pop(n int) []rune {
	popped := (*cs)[len(*cs)-n:]
	*cs = (*cs)[:len(*cs)-n]
	return popped
}

func (cs *crateStack) push(p []rune) {
	for i := len(p) - 1; i >= 0; i-- {
		*cs = append(*cs, p[i])
	}
}

func (cs *crateStack) push9001(p []rune) {
	*cs = append(*cs, p...)
}

func day05(input *bufio.Reader) error {
	type rearrangement struct {
		count, from, to byte
	}

	var stackStrs []string
	var stacks []crateStack
	s := bufio.NewScanner(input)
	for s.Scan() {
		str := s.Text()
		if strings.Contains(str, "[") {
			stackStrs = append(stackStrs, str)
			continue
		}

		n := (len(str) + 1) / 4
		stacks = make([]crateStack, 0, n)
		for s := 0; s < n; s++ {
			stack := make(crateStack, 0, len(stackStrs))
			for i := len(stackStrs) - 1; i >= 0; i-- {
				c := stackStrs[i][1+s*4]
				if c == ' ' {
					break
				}
				stack = append(stack, rune(c))
			}
			stacks = append(stacks, stack)
		}

		s.Scan()
		break
	}

	stacksCopy := make([]crateStack, 0, len(stacks))
	for _, s := range stacks {
		stack := make(crateStack, len(s))
		copy(stack, s)
		stacksCopy = append(stacksCopy, stack)
	}

	var rearrangements []rearrangement
	for s.Scan() {
		var r rearrangement
		if n, err := fmt.Sscanf(s.Text(), "move %d from %d to %d", &r.count, &r.from, &r.to); n != 3 || err != nil {
			panic("bad rearrangement line")
		}
		rearrangements = append(rearrangements, r)
	}

	for _, r := range rearrangements {
		stacks[r.to-1].push(stacks[r.from-1].pop(int(r.count)))
		stacksCopy[r.to-1].push9001(stacksCopy[r.from-1].pop(int(r.count)))
	}

	var tops, tops9001 string
	for i := 0; i < len(stacks); i++ {
		n := len(stacks[i]) - 1
		tops += string(stacks[i][n])
		tops9001 += string(stacksCopy[i][n])
	}

	partOne(tops)
	partTwo(tops9001)

	return nil
}
