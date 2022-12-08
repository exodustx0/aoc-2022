package cmd

import "bufio"

func day08(input *bufio.Reader) (partOne, partTwo any) {
	var treemap []byte
	var w int
	s := bufio.NewScanner(input)
	for s.Scan() {
		str := s.Text()
		if treemap == nil {
			w = len(str)
		}
		treemap = append(treemap, []byte(str)...)
	}
	l := len(treemap)
	h := l / w

	lineOfSight := func(z byte, i, step, end int) (viewingDistance uint, visible bool) {
		for {
			viewingDistance++
			if treemap[i] >= z {
				return
			}
			if i == end {
				visible = true
				return
			}
			i += step
		}
	}

	var visibleCount uint
	var highestScenicScore uint
	for i, z := range treemap {
		x := i % w
		if x == 0 || x == w-1 {
			// On west/east edge.
			visibleCount++
			continue
		}

		y := i / w
		if y == 0 || y == h-1 {
			// On north/south edge.
			visibleCount++
			continue
		}

		var isVisible bool
		x0 := y * w
		west, ok := lineOfSight(z, i-1, -1, x0)
		if ok {
			isVisible = true
		}

		east, ok := lineOfSight(z, i+1, 1, x0+w-1)
		if ok {
			isVisible = true
		}

		north, ok := lineOfSight(z, i-w, -w, x)
		if ok {
			isVisible = true
		}

		south, ok := lineOfSight(z, i+w, w, l-w+x)
		if ok || isVisible {
			visibleCount++
		}
		highestScenicScore = max(highestScenicScore, north*east*south*west)
	}

	partOne = visibleCount
	partTwo = highestScenicScore
	return
}
