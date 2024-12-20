package main

import (
	"advent/utils"
	"fmt"
	"slices"
	"strconv"
)

const day = "20"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part(lines, 2, 100)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part(lines, 20, 100)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part(lines []string, cmax, ceil int) int {
	var start, end utils.Point
	d := make(datamap, len(lines))
	for j, line := range lines {
		d[j] = make([]bool, len(line))
		for i, c := range line {
			if c == 'S' {
				start = utils.Point{X: i, Y: j}
			} else if c == 'E' {
				end = utils.Point{X: i, Y: j}
			}
			d[j][i] = c == '#'
		}
	}
	return d.shortestsPath(start, end, cmax, ceil)
}

type datamap [][]bool
type path []utils.Point
type cpath struct {
	p               path
	cheat, cheating bool
	cpt             int
	start           utils.Point
}
type cheats map[utils.Point]map[utils.Point]bool

func (dm datamap) getPt(pt utils.Point) (bool, bool) {
	return dm.get(pt.X, pt.Y)
}

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

func (dm datamap) shortestPath(start, end utils.Point) path {
	visited := make(map[utils.Point]int)
	queue := utils.PriorityQueue[path]{}
	queue.GPush(path{start}, 0)
	for len(queue) > 0 {
		p, length := queue.GPop()
		//dm.drawWithPath(start, end, p, false, utils.EmptyPoint, utils.EmptyPoint)
		last := p[len(p)-1]
		if last == end {
			return p
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
	return path{}
}

func (dm datamap) shortestsPath(start, end utils.Point, cmax, ceil int) int {
	sp := dm.shortestPath(start, end)
	ref := len(sp) - 1
	visited := make(map[utils.Point]int)
	for i := 0; i < len(sp); i++ {
		visited[sp[len(sp)-i-1]] = i
	}
	total := 0
	for i, pt1 := range sp {
		for j, pt2 := range sp {
			if j <= i {
				continue
			}
			d := pt1.Dist(pt2)
			if d < 2 || d > cmax {
				continue
			}
			r, _ := visited[pt2]
			t := i + d + r
			if t+ceil <= ref {
				total++
			}
		}
	}
	return total
}

func (dm datamap) draw(start, end utils.Point) {
	for y, row := range dm {
		for x, cell := range row {
			if cell {
				fmt.Print("#")
			} else if start.Y == y && start.X == x {
				fmt.Print(utils.Blue + "S" + utils.Reset)
			} else if end.X == x && end.Y == y {
				fmt.Print(utils.Blue + "E" + utils.Reset)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
