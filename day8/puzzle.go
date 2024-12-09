package main

import (
	"advent/utils"
	"fmt"
	"strconv"
)

const day = "8"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	maxY, maxX := len(lines)-1, len(lines[0])-1
	_, antennas := makeData(lines)
	antinodes := make(map[pt]interface{})
	for _, ants := range antennas {
		if len(ants) == 1 {
			continue
		}
		for i := 0; i < len(ants); i++ {
			for j := i + 1; j < len(ants); j++ {
				a1 := ants[i]
				a2 := ants[j]
				v := pt{a2.p.x - a1.p.x, a2.p.y - a1.p.y}
				an := v.add(a2.p)
				if an.inbound(0, maxX, 0, maxY) {
					antinodes[an] = true
				}
				an = a1.p.sub(v)
				if an.inbound(0, maxX, 0, maxY) {
					antinodes[an] = true
				}
			}
		}
	}
	return len(antinodes)
}

func part2(lines []string) int {
	maxY, maxX := len(lines)-1, len(lines[0])-1
	_, antennas := makeData(lines)
	antinodes := make(map[pt]interface{})
	for _, ants := range antennas {
		if len(ants) == 1 {
			continue
		}
		for i := 0; i < len(ants); i++ {
			for j := i + 1; j < len(ants); j++ {
				a1 := ants[i]
				a2 := ants[j]
				v := pt{a2.p.x - a1.p.x, a2.p.y - a1.p.y}
				an := v.add(a1.p)
				for an.inbound(0, maxX, 0, maxY) {
					antinodes[an] = true
					an = an.add(v)
				}
				an = a2.p.sub(v)
				for an.inbound(0, maxX, 0, maxY) {
					antinodes[an] = true
					an = an.sub(v)
				}
			}
		}
	}
	return len(antinodes)
}

func makeData(lines []string) ([][]rune, map[rune][]antenna) {
	data := make([][]rune, 0)
	antennas := make(map[rune][]antenna)
	for y, line := range lines {
		row := make([]rune, 0)
		for x, c := range line {
			row = append(row, c)
			if c != '.' {
				a := antenna{pt{x, y}, c}
				var ants []antenna
				if ats, ok := antennas[c]; ok {
					ants = append(ats, a)
				} else {
					ants = []antenna{a}
				}
				antennas[c] = ants
			}
		}
		data = append(data, row)
	}
	return data, antennas
}

type antenna struct {
	p    pt
	kind rune
}

type pt struct {
	x, y int
}

func (a pt) add(b pt) pt {
	return pt{a.x + b.x, a.y + b.y}
}

func (a pt) sub(b pt) pt {
	return pt{a.x - b.x, a.y - b.y}
}

func (a pt) inbound(minX, maxX, minY, maxY int) bool {
	return a.x >= minX && a.x <= maxX &&
		a.y >= minY && a.y <= maxY
}

func printData(data [][]rune, nodes map[pt]interface{}) {
	failures := make([]pt, 0)
	for p := range nodes {
		old := data[p.y][p.x]
		if old == '.' {
			data[p.y][p.x] = '#'
		} else if old == '#' || old == '*' {
			failures = append(failures, p)
		} else {
			data[p.y][p.x] = '*'
		}
	}
	for _, row := range data {
		for _, r := range row {
			fmt.Printf("%c", r)
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println("Failures: ", len(failures))
	for _, f := range failures {
		fmt.Println(f)
	}
}
