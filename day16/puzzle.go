package main

import (
	"advent/utils"
	"fmt"
	"math"
	"strconv"
)

const day = "16"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	m := newMaze(lines)
	d0 := deer{east, m.start, 0, []utils.Point{m.start}}
	queue, visited := utils.PriorityQueue[deer]{}, make(map[utils.Point]int)
	queue.GPush(d0, 0)
	for len(queue) > 0 {
		d, _ := queue.GPop()
		//fmt.Printf("%d, %v: %d\n", d.pos, d.face, score)
		//m.data[d.pos.Y][d.pos.X] = d.direction()
		//printMaze(m, d, d.score)
		if d.pos == m.end {
			return d.score
		}
		if md, ok := d.moveForward(m, visited, false); ok {
			queue.GPush(md, md.score)
		}
		if md, ok := d.rotate(m, true, visited); ok {
			queue.GPush(md, md.score)
		}
		if md, ok := d.rotate(m, false, visited); ok {
			queue.GPush(md, md.score)
		}
	}
	return 0
}

func part2(lines []string) int {
	m := newMaze(lines)
	d0 := deer{east, m.start, 0, []utils.Point{m.start}}
	queue, visited, bestPlaces := utils.PriorityQueue[deer]{}, make(map[utils.Point]int), make(map[utils.Point]struct{})
	queue.GPush(d0, 0)
	minP := math.MaxInt
	for len(queue) > 0 {
		d, _ := queue.GPop()
		//m.data[d.pos.Y][d.pos.X] = d.direction()
		//printMaze(m, d, d.score)
		if d.pos == m.end {
			if d.score <= minP {
				minP = d.score
				for _, pt := range d.path {
					bestPlaces[pt] = struct{}{}
				}
			}
		} else if d.score < minP {
			if md, ok := d.moveForward(m, visited, false); ok {
				queue.GPush(md, md.score)
			}
			if md, ok := d.rotate(m, true, visited); ok {
				queue.GPush(md, md.score)
			}
			if md, ok := d.rotate(m, false, visited); ok {
				queue.GPush(md, md.score)
			}
		}
	}
	//printMazeBP(m, bestPlaces)
	return len(bestPlaces)
}

type maze struct {
	data       [][]rune
	start, end utils.Point
}

func newMaze(lines []string) maze {
	m := maze{}
	m.data = make([][]rune, len(lines))
	for j, line := range lines {
		m.data[j] = []rune(line)
		for i, c := range line {
			if c == 'S' {
				m.start = utils.Point{X: i, Y: j}
			} else if c == 'E' {
				m.end = utils.Point{X: i, Y: j}
			}
		}
	}
	return m
}

func (m maze) isValid(pt utils.Point) bool {
	return pt.Y >= 0 && pt.Y < len(m.data) && pt.X >= 0 && pt.X < len(m.data[pt.Y]) &&
		m.data[pt.Y][pt.X] != '#'
}

type dir uint8
type deer struct {
	face  dir
	pos   utils.Point
	score int
	path  []utils.Point
}

func (d deer) moveForward(m maze, visited map[utils.Point]int, dry bool) (deer, bool) {
	var pt utils.Point
	switch d.face {
	case east:
		pt = utils.Point{X: d.pos.X + 1, Y: d.pos.Y}
	case south:
		pt = utils.Point{X: d.pos.X, Y: d.pos.Y + 1}
	case west:
		pt = utils.Point{X: d.pos.X - 1, Y: d.pos.Y}
	case north:
		pt = utils.Point{X: d.pos.X, Y: d.pos.Y - 1}
	}
	if m.isValid(pt) {
		score := d.score + 1
		if old, ok := visited[pt]; !ok || old+1000 >= score {
			if !dry {
				visited[pt] = score
			}
			path := make([]utils.Point, len(d.path)+1)
			copy(path, d.path)
			path[len(d.path)] = pt
			return deer{d.face, pt, score, path}, true
		}
	}
	return deer{}, false
}

func (d deer) rotate(m maze, clockwise bool, visited map[utils.Point]int) (deer, bool) {
	nd := d.face + 4 - 1
	if clockwise {
		nd += 2
	}
	nd = nd % 4
	neodeer := deer{nd, d.pos, d.score + 1000, d.path}
	if _, ok := neodeer.moveForward(m, visited, true); ok {
		path := make([]utils.Point, len(d.path))
		copy(path, d.path)
		neodeer.path = path
		return neodeer, true
	}
	return deer{}, false
}

func (d deer) direction() rune {
	switch d.face {
	case east:
		return '>'
	case south:
		return 'v'
	case west:
		return '<'
	case north:
		return '^'
	}
	return '*'
}

func printMaze(m maze, d deer, score int) {
	for j, row := range m.data {
		for i, r := range row {
			if d.pos.X == i && d.pos.Y == j {
				fmt.Print("\033[32m" + string(r) + "\033[0m")
			} else if m.start.X == i && m.start.Y == j {
				fmt.Print("S")
			} else if m.end.X == i && m.end.Y == j {
				fmt.Print("E")
			} else {
				fmt.Print(string(r))
			}
		}
		fmt.Println()
	}
	fmt.Printf("deer at %d facing %s (score: %d)\n", d.pos, string(d.direction()), score)
	fmt.Println()
}

func printMazeBP(m maze, bp map[utils.Point]struct{}) {
	for j, row := range m.data {
		for i, r := range row {
			if m.start.X == i && m.start.Y == j {
				fmt.Print("S")
			} else if m.end.X == i && m.end.Y == j {
				fmt.Print("E")
			} else if _, ok := bp[utils.Point{X: i, Y: j}]; ok {
				fmt.Print("O")
			} else {
				fmt.Print(string(r))
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

const (
	east dir = iota
	south
	west
	north
)
