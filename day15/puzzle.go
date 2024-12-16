package main

import (
	"advent/utils"
	"fmt"
	"strconv"
)

const day = "15"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	d := datamap{}
	d.data = make([][]rune, len(lines))
	empty := 0
	for j, line := range lines {
		if len(line) == 0 {
			empty = j
			break
		}
		d.data[j] = make([]rune, len(line))
		for i, r := range line {
			d.data[j][i] = r
			if r == '@' {
				d.data[j][i] = '.'
				d.robot = utils.Point{X: i, Y: j}
			}
		}
	}
	//d.print()
	for _, row := range lines[empty:] {
		for _, r := range row {
			d.move(r)
			//d.print()
		}
	}
	total := 0
	for j, row := range d.data {
		for i, r := range row {
			if r == 'O' {
				total += 100*j + i
			}
		}
	}
	return total
}

func part2(lines []string) int {
	d := datamap{}
	d.data = make([][]rune, len(lines))
	empty := 0
	for j, line := range lines {
		if len(line) == 0 {
			empty = j
			break
		}
		d.data[j] = make([]rune, 2*len(line))
		for i, r := range line {
			d.data[j][2*i] = r
			d.data[j][2*i+1] = r
			if r == '@' {
				d.data[j][2*i] = '.'
				d.data[j][2*i+1] = '.'
				d.robot = utils.Point{X: 2 * i, Y: j}
			} else if r == 'O' {
				d.data[j][2*i] = '['
				d.data[j][2*i+1] = ']'
			}
		}
	}
	//d.print()
	for _, row := range lines[empty:] {
		for _, r := range row {
			d.move2(r)
			//d.print()
		}
	}
	total := 0
	for j, row := range d.data {
		for i, r := range row {
			if r == '[' {
				total += 100*j + i
			}
		}
	}
	return total
}

type datamap struct {
	robot utils.Point
	data  [][]rune
}

func (d *datamap) move(r rune) {
	var nextF func(point utils.Point) utils.Point
	switch r {
	case '>':
		nextF = func(point utils.Point) utils.Point { return point.Right() }
	case '<':
		nextF = func(point utils.Point) utils.Point { return point.Left() }
	case '^':
		nextF = func(point utils.Point) utils.Point { return point.Up() }
	case 'v':
		nextF = func(point utils.Point) utils.Point { return point.Down() }
	}
	if d.moveThing(d.robot, nextF) {
		d.robot = nextF(d.robot)
	}
}

func (d *datamap) move2(r rune) {
	var nextF utils.Mover
	updown := false
	switch r {
	case '>':
		nextF = utils.Righter
	case '<':
		nextF = utils.Lefter
	case '^':
		updown = true
		nextF = utils.Upper
	case 'v':
		updown = true
		nextF = utils.Downer
	}
	if !updown && d.moveThing(d.robot, nextF) {
		d.robot = nextF(d.robot)
	} else if updown && d.moveThing2(d.robot, nextF) {
		d.robot = nextF(d.robot)
	}
}

func (d *datamap) moveThing(pt utils.Point, dir func(point utils.Point) utils.Point) bool {
	nextP := dir(pt)
	nextR := d.data[nextP.Y][nextP.X]
	if nextR == '#' {
		return false
	}
	if nextR == '.' || d.moveThing(nextP, dir) {
		d.data[nextP.Y][nextP.X] = d.data[pt.Y][pt.X]
		d.data[pt.Y][pt.X] = '.'
		return true
	}
	return false
}

// up or down
func (d *datamap) moveThing2(pt utils.Point, dir utils.Mover) bool {
	nextP := dir(pt)
	nextR := d.data[nextP.Y][nextP.X]
	if nextR == '#' {
		return false
	}
	if nextR == '.' {
		d.data[nextP.Y][nextP.X] = d.data[pt.Y][pt.X]
		d.data[pt.Y][pt.X] = '.'
		return true
	}
	var nextP2 utils.Point
	if nextR == '[' {
		nextP2 = utils.Point{X: nextP.X + 1, Y: nextP.Y}
	} else {
		nextP2 = utils.Point{X: nextP.X - 1, Y: nextP.Y}
	}
	bck := backup(d.data)
	c1 := d.moveThing2(nextP, dir)
	c2 := d.moveThing2(nextP2, dir)
	if c1 && c2 {
		d.data[nextP.Y][nextP.X] = d.data[pt.Y][pt.X]
		d.data[pt.Y][pt.X] = '.'
		return true
	}
	if c1 || c2 {
		d.data = bck
	}
	return false
}

func (d *datamap) print() {
	for j, row := range d.data {
		for i, r := range row {
			if d.robot.Y == j && d.robot.X == i {
				fmt.Print("\033[32m@\033[0m")
			} else {
				fmt.Printf("%c", r)
			}
		}
		fmt.Println()
	}
}

func backup(data [][]rune) [][]rune {
	out := make([][]rune, len(data))
	for i, r := range data {
		out[i] = make([]rune, len(r))
		copy(out[i], r)
	}
	return out
}
