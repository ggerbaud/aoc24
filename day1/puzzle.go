package main

import (
	"advent/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const day = "1"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part2(lines []string) int {
	lefts, rights := make([]int, 0), make(map[int]int)
	for _, line := range lines {
		digits := strings.Split(line, "   ")
		if len(digits) == 2 {
			lefts = append(lefts, utils.ParseInt(digits[0]))
			r := utils.ParseInt(digits[1])
			if n, ok := rights[r]; ok {
				rights[r] = n + 1
			} else {
				rights[r] = 1
			}
		}
	}
	total := 0
	for _, left := range lefts {
		if n, ok := rights[left]; ok {
			total += n * left
		}
	}
	return total
}

func part1(lines []string) int {
	lefts, rights := make([]int, 0), make([]int, 0)
	for _, line := range lines {
		digits := strings.Split(line, "   ")
		if len(digits) == 2 {
			lefts = append(lefts, utils.ParseInt(digits[0]))
			rights = append(rights, utils.ParseInt(digits[1]))
		}
	}
	sort.Ints(lefts)
	sort.Ints(rights)
	total := 0
	for i := 0; i < len(lefts); i++ {
		l, r := lefts[i], rights[i]
		dist := utils.Abs(l - r)
		total += dist
	}
	return total
}
