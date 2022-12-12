package cmd

import (
	"bufio"
	"math"
	"sort"
)

type hillGrid struct {
	nodes      []*hillNode
	test       []*hillNode
	start, end *hillNode
	w, h       int
}

func (g *hillGrid) initialise() {
	for i := range g.nodes {
		g.nodes[i].local = math.MaxInt
		g.nodes[i].global = math.MaxInt
		g.nodes[i].visited = false
		g.nodes[i].parent = nil
	}
	g.start.local = 0
}

func (g *hillGrid) solve() {
	g.test = append(g.test, g.start)
	for len(g.test) != 0 {
		sort.Slice(g.test, func(i, j int) bool { return g.test[i].global < g.test[j].global })

		i := 0
		for len(g.test) > i && g.test[i].visited {
			i++
		}
		if i != 0 {
			g.test = g.test[:copy(g.test, g.test[i:])]
			if len(g.test) == 0 {
				break
			}
		}

		cur := g.test[0]
		cur.visited = true
		for _, n := range cur.neighbours {
			if !n.visited {
				g.test = append(g.test, n)
			}

			if l := cur.local + 1; l < n.local {
				n.parent = cur
				n.local = l
				n.global = l + abs(n.x-g.end.x) + abs(n.y-g.end.y)
			}
		}
	}
}

func (g *hillGrid) steps() int {
	steps := 0
	for cur := g.end; cur.parent != nil; cur = cur.parent {
		steps++
	}
	return steps
}

func (g *hillGrid) printPath() {
	m := make([][]byte, g.h)
	for y := range m {
		m[y] = make([]byte, g.w)
		for x := range m[y] {
			m[y][x] = '.'
		}
	}
	for cur := g.end; cur.parent != nil; cur = cur.parent {
		var arrow byte
		switch {
		case cur.y == cur.parent.y+1:
			arrow = 'v'
		case cur.x == cur.parent.x-1:
			arrow = '<'
		case cur.y == cur.parent.y-1:
			arrow = '^'
		case cur.x == cur.parent.x+1:
			arrow = '>'
		}
		m[cur.parent.y][cur.parent.x] = arrow
	}
	for _, row := range m {
		println(string(row))
	}
}

type hillNode struct {
	height     byte
	x, y       int
	neighbours []*hillNode
	local      int
	global     int
	visited    bool
	parent     *hillNode
}

func (n *hillNode) tryAddNeighbour(o *hillNode) {
	if n.height+1 < o.height {
		return
	}
	n.neighbours = append(n.neighbours, o)
}

func day12(input *bufio.Reader) (partOne, partTwo any) {
	var grid hillGrid

	s := bufio.NewScanner(input)
	for y := 0; s.Scan(); y++ {
		str := s.Text()
		for x, ch := range str {
			n := &hillNode{
				x:      x,
				y:      y,
				local:  math.MaxInt,
				global: math.MaxInt,
			}

			switch ch {
			case 'S':
				n.local = 0
				ch = 'a'
				grid.start = n
			case 'E':
				ch = 'z'
				grid.end = n
			}

			n.height = byte(ch - 'a')
			grid.nodes = append(grid.nodes, n)
		}
		grid.w = len(str)
	}
	grid.h = len(grid.nodes) / grid.w

	for y := 0; y < grid.h; y++ {
		for x := 0; x < grid.w; x++ {
			i := x + y*grid.w
			n := grid.nodes[i]
			if y != 0 {
				n.tryAddNeighbour(grid.nodes[i-grid.w])
			}
			if x != grid.w-1 {
				n.tryAddNeighbour(grid.nodes[i+1])
			}
			if y != grid.h-1 {
				n.tryAddNeighbour(grid.nodes[i+grid.w])
			}
			if x != 0 {
				n.tryAddNeighbour(grid.nodes[i-1])
			}
		}
	}

	grid.solve()
	partOne = grid.steps()

	return
}
