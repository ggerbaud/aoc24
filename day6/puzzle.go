package main

import (
	"advent/utils"
	"fmt"
	"strconv"
)

const day = "6"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	data, g := makeMapAndGuard(lines)
	for out := false; !out; out, _ = g.move(data) {
	}
	total := 0
	for _, kinds := range data {
		for _, k := range kinds {
			if k != free && k != obstacle {
				total++
			}
		}
	}
	return total
}

func part2(lines []string) int {
	data, g := makeMapAndGuard(lines)
	gx, gy := g.x, g.y
	for out := false; !out; out, _ = g.move(data) {
	}
	points := make([]point, 0)
	for y, kinds := range data {
		for x, k := range kinds {
			if k != free && k != obstacle {
				points = append(points, point{x, y})
			}
		}
	}
	total := 0
	for _, pt := range points {
		if pt.x == gx && pt.y == gy {
			continue
		}
		g.x, g.y, g.d = gx, gy, up
		data.clear()
		data[pt.y][pt.x] = obstacle
		loop, out := false, false
		for ; !out && !loop; out, loop = g.move(data) {
		}
		if loop {
			total++
		}
		data[pt.y][pt.x] = free
	}
	return total
}

func makeMapAndGuard(lines []string) (room, *guard) {
	data := make(room, len(lines))
	g := guard{d: up}
	for j, line := range lines {
		row := make([]kind, len(line))
		for i, c := range line {
			switch c {
			case '.':
				row[i] = free
			case '#':
				row[i] = obstacle
			case '^':
				row[i] = visitedUp
				g.x = i
				g.y = j
			}
		}
		data[j] = row
	}
	return data, &g
}

type guard struct {
	x, y int
	d    dir
}

type point struct {
	x, y int
}

type room [][]kind
type dir int
type kind int

func (r room) isIn(x, y int) bool {
	return y >= 0 && y < len(r) && x >= 0 && x < len(r[0])
}

func (r room) isFree(x, y int) bool {
	return r[y][x] != obstacle
}

func (r room) clear() {
	for y, kinds := range r {
		for x, k := range kinds {
			if k != obstacle {
				r[y][x] = free
			}
		}
	}
}

func (g *guard) turnRight() {
	switch g.d {
	case up:
		g.d = right
	case right:
		g.d = down
	case down:
		g.d = left
	case left:
		g.d = up
	}
}

func (g *guard) move(r room) (bool, bool) {
	oldX, oldY := g.x, g.y
	switch g.d {
	case up:
		g.y--
	case right:
		g.x++
	case down:
		g.y++
	case left:
		g.x--
	}
	if !r.isIn(g.x, g.y) {
		return true, false
	}
	if r.isFree(g.x, g.y) {
		k0 := r[g.y][g.x]
		k := visitedUp
		switch g.d {
		case down:
			k = visitedDown
		case left:
			k = visitedLeft
		case right:
			k = visitedRight
		default:
			k = visitedUp
		}
		nk := k0 | k
		if nk == k0 {
			return false, true
		}
		r[g.y][g.x] = nk
	} else {
		g.x = oldX
		g.y = oldY
		g.turnRight()
		return g.move(r)
	}
	return false, false
}

const (
	up dir = iota
	down
	left
	right
)
const (
	free kind = iota
	visitedUp
	visitedDown
	visitedUD
	visitedLeft
	visitedUL
	visitedDL
	visitedUDL
	visitedRight
	visitedUR
	visitedDR
	visitedUDR
	visitedLR
	visitedULR
	visitedDLR
	visitedUDLR
	obstacle
)
