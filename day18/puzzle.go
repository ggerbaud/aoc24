package main

import (
	"advent/utils"
	"fmt"
	"slices"
	"strconv"
)

const day = "18"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines, 70, 1024)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	pt := part2(lines, 70, 1024)
	fmt.Printf("Day#%s.2 : %d", day, pt)
}

func part1(lines []string, w, steps int) int {
	dm := make(datamap, w+1)
	for i := 0; i < steps; i++ {
		coords := utils.ListOfNumbers(lines[i], ",")
		dm.set(coords[0], coords[1], true)
	}
	//dm.draw()
	return shortestPath(dm, utils.Point{}, utils.Point{X: w, Y: w})
}

func part2(lines []string, w, atleast int) utils.Point {
	dm := make(datamap, w+1)
	for i := 0; i < atleast; i++ {
		coords := utils.ListOfNumbers(lines[i], ",")
		dm.set(coords[0], coords[1], true)
	}
	for i := atleast; i < len(lines); i++ {
		// brute force is good enough for this
		// an alternative would be to verify if the new byte is on the last shortest path
		coords := utils.ListOfNumbers(lines[i], ",")
		dm.set(coords[0], coords[1], true)
		sp := shortestPath(dm, utils.Point{}, utils.Point{X: w, Y: w})
		if sp == -1 {
			return utils.Point{X: coords[0], Y: coords[1]}
		}
	}
	return utils.Point{}
}

func shortestPath(dm datamap, start, end utils.Point) int {
	visited := make(map[utils.Point]int)
	queue := utils.PriorityQueue[path]{}
	queue.GPush(path{start}, 0)
	for len(queue) > 0 {
		p, length := queue.GPop()
		//dm.drawWithPath(p)
		last := p[len(p)-1]
		if last == end {
			return length
		}
		if i, ok := visited[last]; ok && i <= length {
			continue
		}
		visited[last] = length
		u, d, r, l := last.Up(), last.Down(), last.Right(), last.Left()
		for _, pt := range []utils.Point{u, d, r, l} {
			if oqp, ok := dm.get(pt.X, pt.Y); ok && !oqp && !slices.Contains(p, pt) {
				pds := length + 1
				np := make(path, len(p)+1)
				copy(np, p)
				np[len(np)-1] = pt
				queue.GPush(np, pds)
			}
		}
	}
	return -1
}

type datamap [][]bool
type path []utils.Point

func (dm datamap) get(x, y int) (bool, bool) {
	if y < 0 || y >= len(dm) {
		return false, false
	}
	row := dm[y]
	if row == nil {
		row = make([]bool, len(dm))
		dm[y] = row
	}
	if x < 0 || x >= len(row) {
		return false, false
	}
	return row[x], true
}

func (dm datamap) set(x, y int, v bool) {
	row := dm[y]
	if row == nil {
		row = make([]bool, len(dm))
		dm[y] = row
	}
	row[x] = v
}

func (dm datamap) draw() {
	for _, row := range dm {
		for _, cell := range row {
			if cell {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (dm datamap) drawWithPath(p path) {
	out := make([][]rune, len(dm))
	for j, row := range dm {
		out[j] = make([]rune, len(row))
		for i, cell := range row {
			out[j][i] = '.'
			if cell {
				out[j][i] = '#'
			}
		}
	}
	for _, pt := range p {
		out[pt.Y][pt.X] = 'O'
	}
	for _, runes := range out {
		fmt.Println(string(runes))
	}
	println()
}
