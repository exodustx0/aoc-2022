package cmd

import (
	"bufio"
	"image"
	"strconv"
	"strings"
)

func day14(input *bufio.Reader) (partOne, partTwo any) {
	const (
		_rock = '#'
		_sand = 'o'
	)
	cave := make(map[image.Point]byte)
	bounds := image.Rect(500, 0, 500, 0)
	/* printCave := func() {
		pt := bounds.Min
		pt.X--
		for ; pt.Y <= bounds.Max.Y+1; pt.Y++ {
			for ; pt.X <= bounds.Max.X+1; pt.X++ {
				if elem, ok := cave[pt]; ok {
					print(string(elem))
				} else {
					print(".")
				}
			}
			println()
			pt.X = bounds.Min.X - 1
		}
		for ; pt.X <= bounds.Max.X+1; pt.X++ {
			print(string(_rock))
		}
		println()
	} */

	s := bufio.NewScanner(input)
	for s.Scan() {
		pointStrs := strings.Split(s.Text(), " -> ")
		points := make([]image.Point, 0, len(pointStrs))
		for _, pointStr := range pointStrs {
			xStr, yStr, ok := strings.Cut(pointStr, ",")
			if !ok {
				panic("no x,y separator")
			}
			x, err := strconv.ParseUint(xStr, 10, 0)
			if err != nil {
				panic(err)
			}
			y, err := strconv.ParseUint(yStr, 10, 0)
			if err != nil {
				panic(err)
			}
			points = append(points, image.Pt(int(x), int(y)))
		}

		current := points[0]
		points = points[1:]
		bounds.Min.X = min(bounds.Min.X, current.X)
		bounds.Max.X = max(bounds.Max.X, current.X)
		bounds.Max.Y = max(bounds.Max.Y, current.Y)
		cave[current] = _rock

		var add image.Point
		for _, pt := range points {
			switch {
			case current.X < pt.X:
				add = image.Pt(1, 0)
			case current.X > pt.X:
				add = image.Pt(-1, 0)
			case current.Y < pt.Y:
				add = image.Pt(0, 1)
			case current.Y > pt.Y:
				add = image.Pt(0, -1)
			}

			for current != pt {
				current = current.Add(add)
				bounds.Min.X = min(bounds.Min.X, current.X)
				bounds.Max.X = max(bounds.Max.X, current.X)
				bounds.Max.Y = max(bounds.Max.Y, current.Y)
				cave[current] = _rock
			}
		}
	}

	sandStart := image.Pt(500, 0)
	sandAdds := [3]image.Point{
		image.Pt(0, 1),
		image.Pt(-1, 1),
		image.Pt(1, 1),
	}
	dropSand := func() image.Point {
		sand := sandStart
	fallLoop:
		for sand.Y != bounds.Max.Y+1 {
			for _, sandAdd := range sandAdds {
				newSand := sand.Add(sandAdd)
				if _, ok := cave[newSand]; !ok {
					sand = newSand
					continue fallLoop
				}
			}
			return sand
		}
		return sand
	}

	i := 0
	for ; ; i++ {
		sand := dropSand()
		cave[sand] = _sand
		if sand.Y >= bounds.Max.Y {
			break
		}
	}

	partOne = i

	i++
	for ; ; i++ {
		sand := dropSand()
		bounds.Min.X = min(bounds.Min.X, sand.X)
		bounds.Max.X = max(bounds.Max.X, sand.X)
		cave[sand] = _sand
		if sand == sandStart {
			break
		}
	}

	partTwo = i + 1
	return
}
