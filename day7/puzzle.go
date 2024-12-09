package main

import (
	"advent/utils"
	"fmt"
	"strconv"
	"strings"
)

const day = "7"

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
		splits := strings.Split(line, ":")
		result := utils.ParseInt(splits[0])
		ops := utils.ListOfNumbers(strings.TrimSpace(splits[1]), " ")
		if solve(result, ops) {
			total += result
		}
	}
	return total
}

func part2(lines []string) int {
	total := 0
	for _, line := range lines {
		splits := strings.Split(line, ":")
		result := utils.ParseInt(splits[0])
		ops := utils.ListOfNumbers(strings.TrimSpace(splits[1]), " ")
		if solve2(result, ops) {
			//fmt.Printf("%d is solved by %d\n", result, ops)
			total += result
		}
	}
	return total
}

func solve(result int, data []int) bool {
	if len(data) == 2 {
		return data[0]*data[1] == result ||
			data[0]+data[1] == result
	}
	t, q := data[len(data)-1], data[0:len(data)-1]
	if !solve(result-t, q) {
		return (result/t)*t == result && solve(result/t, q)
	}
	return true
}

func solve2(result int, data []int) bool {
	if len(data) == 2 {
		return data[0]*data[1] == result ||
			data[0]+data[1] == result ||
			concat(data[0], data[1]) == result
	}
	a, b, q := data[0], data[1], data[2:]
	return solve2(result, append([]int{a + b}, q...)) ||
		solve2(result, append([]int{a * b}, q...)) ||
		solve2(result, append([]int{concat(a, b)}, q...))

}

func concat(a, b int) int {
	return utils.ParseInt(strconv.Itoa(a) + strconv.Itoa(b))
}
