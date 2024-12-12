package main

import (
	"advent/utils"
	"fmt"
	"strconv"
)

const day = "12"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	data := make([][]rune, 0)
	for _, line := range lines {
		row := make([]rune, 0)
		for _, r := range line {
			row = append(row, r)
		}
		data = append(data, row)
	}
	total := 0
	for _, f := range splitFields(data) {
		total += f.perimeter * f.area
	}
	return total
}

func part2(lines []string) int {
	data := make([][]rune, 0)
	for _, line := range lines {
		row := make([]rune, 0)
		for _, r := range line {
			row = append(row, r)
		}
		data = append(data, row)
	}
	total := 0
	for _, f := range splitFields(data) {
		total += f.corners * f.area
	}
	return total
}

func splitFields(data [][]rune) []*field {
	visited := make(map[utils.Point]bool)
	fields := make([]*field, 0)
	for j := 0; j < len(data); j++ {
		for i := 0; i < len(data[j]); i++ {
			pt := utils.Point{X: i, Y: j}
			if v, ok := visited[pt]; ok && v {
				continue
			}
			// start a new field
			f := &field{v: data[j][i]}
			fields = append(fields, f)
			fieldExploration(data, visited, f, pt)
		}
	}
	return fields
}

func fieldExploration(data [][]rune, visited map[utils.Point]bool, f *field, current utils.Point) {
	visited[current] = true
	f.area++
	f.perimeter += 4
	u, d, l, r := current.Up(), current.Down(), current.Left(), current.Right()
	ku, kd, kl, kr := isOfKind(data, u, f.v), isOfKind(data, d, f.v), isOfKind(data, l, f.v), isOfKind(data, r, f.v)
	if ku {
		f.perimeter--
		if !visited[u] {
			fieldExploration(data, visited, f, u)
		}
	}
	if kd {
		f.perimeter--
		if !visited[d] {
			fieldExploration(data, visited, f, d)
		}
	}
	if kl {
		f.perimeter--
		if !visited[l] {
			fieldExploration(data, visited, f, l)
		}
	}
	if kr {
		f.perimeter--
		if !visited[r] {
			fieldExploration(data, visited, f, r)
		}
	}
	if !ku && !kl {
		f.corners++
	}
	if !ku && !kr {
		f.corners++
	}
	if !kd && !kl {
		f.corners++
	}
	if !kd && !kr {
		f.corners++
	}
	if ku && kl && !isOfKind(data, u.Left(), f.v) {
		f.corners++
	}
	if ku && kr && !isOfKind(data, u.Right(), f.v) {
		f.corners++
	}
	if kd && kl && !isOfKind(data, d.Left(), f.v) {
		f.corners++
	}
	if kd && kr && !isOfKind(data, d.Right(), f.v) {
		f.corners++
	}
}

func isOfKind(data [][]rune, pt utils.Point, v rune) bool {
	return pt.Y >= 0 && pt.Y < len(data) && pt.X >= 0 && pt.X < len(data[pt.Y]) && data[pt.Y][pt.X] == v
}

type field struct {
	v                        rune
	area, perimeter, corners int
}
