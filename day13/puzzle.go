package main

import (
	"advent/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const day = "13"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	problems := make([]problem, 0)
	p := problem{}
	for _, l := range lines {
		if strings.HasPrefix(l, "Button A:") {
			xy := strings.Split(strings.Split(l, ":")[1], ",")
			p.x1 = utils.ParseInt(strings.Split(xy[0], "+")[1])
			p.y1 = utils.ParseInt(strings.Split(xy[1], "+")[1])
		} else if strings.HasPrefix(l, "Button B:") {
			xy := strings.Split(strings.Split(l, ":")[1], ",")
			p.x2 = utils.ParseInt(strings.Split(xy[0], "+")[1])
			p.y2 = utils.ParseInt(strings.Split(xy[1], "+")[1])
		} else if strings.HasPrefix(l, "Prize:") {
			xy := strings.Split(strings.Split(l, ":")[1], ",")
			p.xt = utils.ParseInt(strings.Split(xy[0], "=")[1])
			p.yt = utils.ParseInt(strings.Split(xy[1], "=")[1])
			problems = append(problems, p)
			p = problem{}
		}
	}
	fmt.Printf("Solving %d problems\n", len(problems))
	total := 0
	for _, pbm := range problems {
		if a, b, ok := pbm.solve(100); ok {
			fmt.Printf("Solving %v => A: %d, B: %d\n", pbm, a, b)
			total += 3*a + b
		} else {
			fmt.Printf("Solving %v => XXX (%d, %d)\n", pbm, a, b)
		}
	}
	return total
}

func part2(lines []string) int {
	problems := make([]problem, 0)
	p := problem{}
	for _, l := range lines {
		if strings.HasPrefix(l, "Button A:") {
			xy := strings.Split(strings.Split(l, ":")[1], ",")
			p.x1 = utils.ParseInt(strings.Split(xy[0], "+")[1])
			p.y1 = utils.ParseInt(strings.Split(xy[1], "+")[1])
		} else if strings.HasPrefix(l, "Button B:") {
			xy := strings.Split(strings.Split(l, ":")[1], ",")
			p.x2 = utils.ParseInt(strings.Split(xy[0], "+")[1])
			p.y2 = utils.ParseInt(strings.Split(xy[1], "+")[1])
		} else if strings.HasPrefix(l, "Prize:") {
			xy := strings.Split(strings.Split(l, ":")[1], ",")
			p.xt = utils.ParseInt(strings.Split(xy[0], "=")[1])
			p.yt = utils.ParseInt(strings.Split(xy[1], "=")[1])
			problems = append(problems, p)
			p = problem{}
		}
	}
	fmt.Printf("Solving %d problems\n", len(problems))
	total := 0
	for _, pbm := range problems {
		pbm.xt += 10000000000000
		pbm.yt += 10000000000000
		if a, b, ok := pbm.solve(math.MaxInt); ok {
			fmt.Printf("Solving %v => A: %d, B: %d\n", pbm, a, b)
			total += 3*a + b
		} else {
			fmt.Printf("Solving %v => XXX (%d, %d)\n", pbm, a, b)
		}
	}
	return total
}

type problem struct {
	x1, x2, y1, y2, xt, yt int
}

func (p problem) solve(max int) (int, int, bool) {
	if p.x2 == 0 {
		return 0, 0, false
	}
	x1y2, xty2, y1x2, ytx2 := p.x1*p.y2, p.xt*p.y2, p.y1*p.x2, p.yt*p.x2
	down := x1y2 - y1x2
	up := xty2 - ytx2
	if down == 0 || up%down != 0 {
		return 0, 0, false
	}
	a := up / down
	if a < 0 || a >= max {
		return a, 0, false
	}
	b := p.xt - p.x1*a
	if b%p.x2 != 0 {
		return a, 0, false
	}
	b = b / p.x2
	if b < 0 || b >= max {
		return a, b, false
	}
	if p.x1*a+p.x2*b != p.xt ||
		p.y1*a+p.y2*b != p.yt {
		return a, b, false
	}
	return a, b, true
}
