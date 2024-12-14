package main

import (
	"advent/utils"
	"fmt"
	"strconv"
	"strings"
)

const day = "14"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines, 101, 103)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines, 101, 103)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string, w, h int) int {
	robots := make([]*robot, 0)
	for _, line := range lines {
		data := strings.Split(line, " ")
		poss := utils.ListOfNumbers(strings.Split(data[0], "=")[1], ",")
		speeds := utils.ListOfNumbers(strings.Split(data[1], "=")[1], ",")
		r := robot{pos: utils.Point{X: poss[0], Y: poss[1]}, speed: utils.Point{X: speeds[0], Y: speeds[1]}, max: utils.Point{X: w, Y: h}}
		robots = append(robots, &r)
	}
	hw, hh := w/2, h/2
	q1, q2, q3, q4 := 0, 0, 0, 0
	for _, r := range robots {
		r.Move(100)
		if r.pos.X > hw {
			if r.pos.Y > hh {
				q4++
			} else if r.pos.Y < hh {
				q2++
			}
		} else if r.pos.X < hw {
			if r.pos.Y > hh {
				q3++
			} else if r.pos.Y < hh {
				q1++
			}
		}
	}
	return q1 * q2 * q3 * q4
}

func part2(lines []string, w, h int) int {
	robots := make([]*robot, 0)
	for _, line := range lines {
		data := strings.Split(line, " ")
		poss := utils.ListOfNumbers(strings.Split(data[0], "=")[1], ",")
		speeds := utils.ListOfNumbers(strings.Split(data[1], "=")[1], ",")
		r := robot{pos: utils.Point{X: poss[0], Y: poss[1]}, speed: utils.Point{X: speeds[0], Y: speeds[1]}, max: utils.Point{X: w, Y: h}}
		robots = append(robots, &r)
	}
	//fmt.Println("Init")
	//draw(robots, w, h)
	sec := 1
	for {
		for _, r := range robots {
			r.Move(1)
		}
		out := draw(robots, w, h, false)
		if strings.Contains(out, "XXXXXXX") {
			//fmt.Printf("After %d seconds\n", sec)
			//print(robots, w, h)
			break
		}
		sec++
	}
	return sec
}

type robot struct {
	pos, speed, max utils.Point
}

func (r *robot) Move(steps int) {
	r.pos.X = r.pos.X + steps*r.speed.X
	for r.pos.X < 0 {
		r.pos.X += r.max.X
	}
	r.pos.X = r.pos.X % r.max.X
	r.pos.Y = r.pos.Y + steps*r.speed.Y
	for r.pos.Y < 0 {
		r.pos.Y += r.max.Y
	}
	r.pos.Y = r.pos.Y % r.max.Y
}

func draw(robots []*robot, w, h int, doubles bool) string {
	out := make([][]rune, h)
	for i := 0; i < h; i++ {
		out[i] = make([]rune, w)
	}
	for _, r := range robots {
		x := out[r.pos.Y][r.pos.X]
		if x == 0 || !doubles {
			out[r.pos.Y][r.pos.X] = 'X'
		} else {
			out[r.pos.Y][r.pos.X] = '*'
		}
	}
	str := ""
	for _, runes := range out {
		for _, r := range runes {
			c := string(r)
			if r == 0 {
				c = "."
			}
			str += c
		}
		str += "\n"
	}
	return str
}

func print(robots []*robot, w, h int) {
	fmt.Println(draw(robots, w, h, true))
}
