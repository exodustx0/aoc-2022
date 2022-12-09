package cmd

import (
	"bufio"
	"strconv"
	"strings"
)

type knotPoint struct{ x, y int }

func (p *knotPoint) add(o knotPoint) {
	p.x += o.x
	p.y += o.y
}

func (p *knotPoint) follow(o knotPoint) {
	dx := p.x - o.x
	dy := p.y - o.y
	adx := abs(dx)
	ady := abs(dy)
	if adx == 2 && ady == 2 {
		p.x -= dx >> 1
		p.y -= dy >> 1
	} else if adx == 2 {
		p.y = o.y
		p.x -= dx >> 1
	} else if ady == 2 {
		p.x = o.x
		p.y -= dy >> 1
	}
}

func day09(input *bufio.Reader) (partOne, partTwo any) {
	type nothing struct{}

	var h knotPoint
	var t [9]knotPoint
	smallTailPts := map[knotPoint]nothing{h: {}}
	bigTailPts := map[knotPoint]nothing{h: {}}

	s := bufio.NewScanner(input)
	for s.Scan() {
		moveStr, stepsStr, ok := strings.Cut(s.Text(), " ")
		if !ok {
			panic("no space-separation")
		}

		var move knotPoint
		switch moveStr {
		case "L":
			move = knotPoint{-1, 0}
		case "R":
			move = knotPoint{1, 0}
		case "U":
			move = knotPoint{0, -1}
		case "D":
			move = knotPoint{0, 1}
		default:
			panic("unknown move")
		}

		steps, err := strconv.ParseUint(stepsStr, 10, 8)
		if err != nil {
			panic(err)
		}

		for ; steps > 0; steps-- {
			h.add(move)

			f := h
			for i := range t {
				t[i].follow(f)
				f = t[i]
			}

			smallTailPts[t[0]] = nothing{}
			bigTailPts[t[8]] = nothing{}
		}
	}

	partOne = len(smallTailPts)
	partTwo = len(bigTailPts)

	return
}
