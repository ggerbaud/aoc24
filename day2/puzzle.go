package main

import (
	"advent/utils"
	"fmt"
	"strconv"
)

const day = "2"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	total := 0
	for _, line := range lines {
		l := makeLevel(line)
		if l.isSafe() {
			total++
		}
	}
	return total
}

func part2(lines []string) int {
	total := 0
	for _, line := range lines {
		l := makeLevel(line)
		if l.isTolerablySafe() {
			total++
		}
	}
	return total
}

func makeLevel(line string) level {
	return utils.ListOfNumbers(line, " ")
}

type level []int

func (l level) isSafe() bool {
	if len(l) < 2 {
		return true
	}
	diff := l[1] - l[0]
	if diff == 0 || utils.Abs(diff) > 3 {
		return false
	}
	for i := 2; i < len(l); i++ {
		d := l[i] - l[i-1]
		if d == 0 || utils.Abs(d) > 3 || d*diff < 0 {
			return false
		}
	}
	return true
}

func (l level) isTolerablySafe() bool {
	if l.isSafe() {
		return true
	}
	for i := 0; i < len(l); i++ {
		l2 := make(level, len(l))
		copy(l2, l)
		var l3 = append(l2[0:i], l2[i+1:]...)
		if l3.isSafe() {
			return true
		}
	}
	return false
}
