package main

import (
	"advent/utils"
	"fmt"
	"strconv"
)

const day = "10"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	data := make(datamap, len(lines))
	starts := make([]utils.Point, 0)
	for j, line := range lines {
		row := make([]int, len(line))
		for i, d := range line {
			row[i] = int(d - '0')
			if row[i] == 0 {
				starts = append(starts, utils.Point{X: i, Y: j})
			}
		}
		data[j] = row
	}
	total := 0
	for _, start := range starts {
		total += trailheadScore(data, start)
	}
	return total
}

func part2(lines []string) int {
	data := make(datamap, len(lines))
	starts := make([]utils.Point, 0)
	for j, line := range lines {
		row := make([]int, len(line))
		for i, d := range line {
			row[i] = int(d - '0')
			if row[i] == 0 {
				starts = append(starts, utils.Point{X: i, Y: j})
			}
		}
		data[j] = row
	}
	total := 0
	for _, start := range starts {
		total += trailheadScore2(data, start)
	}
	return total
}

func trailheadScore(data datamap, start utils.Point) int {
	ends := make(map[utils.Point]interface{})
	works := make(map[work]interface{})
	queue := make([]work, 0)
	queue = append(queue, work{start, start, 1})
	for len(queue) > 0 {
		w := queue[0]
		queue = queue[1:]
		if _, ok := works[w]; ok {
			continue
		}
		works[w] = nil
		for _, p := range nextStep(data, w.end, w.s) {
			if w.s == 9 {
				ends[p] = nil
			} else {
				queue = append(queue, work{w.start, p, w.s + 1})
			}
		}
	}
	return len(ends)
}

func trailheadScore2(data datamap, start utils.Point) int {
	score := 0
	queue := make([]work, 0)
	queue = append(queue, work{start, start, 1})
	for len(queue) > 0 {
		w := queue[0]
		queue = queue[1:]
		for _, p := range nextStep(data, w.end, w.s) {
			if w.s == 9 {
				score++
			} else {
				queue = append(queue, work{w.start, p, w.s + 1})
			}
		}
	}
	return score
}

func nextStep(data datamap, lp utils.Point, step int) []utils.Point {
	results := make([]utils.Point, 0)
	if data.check(lp.X, lp.Y-1, step) {
		results = append(results, utils.Point{X: lp.X, Y: lp.Y - 1})
	}
	if data.check(lp.X, lp.Y+1, step) {
		results = append(results, utils.Point{X: lp.X, Y: lp.Y + 1})
	}
	if data.check(lp.X-1, lp.Y, step) {
		results = append(results, utils.Point{X: lp.X - 1, Y: lp.Y})
	}
	if data.check(lp.X+1, lp.Y, step) {
		results = append(results, utils.Point{X: lp.X + 1, Y: lp.Y})
	}
	return results
}

type datamap [][]int
type path []utils.Point
type work struct {
	start, end utils.Point
	s          int
}
type work2 struct {
	p path
	s int
}

func (d datamap) check(x, y, v int) bool {
	return y >= 0 && y < len(d) && x >= 0 && x < len(d[y]) && d[y][x] == v
}
