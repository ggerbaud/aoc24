package main

import (
	"advent/utils"
	"fmt"
	"strconv"
)

const day = "11"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	data := utils.ListOfNumbers(lines[0], " ")
	return applyRulesNTimes4All(data, 25)
}

func part2(lines []string) int {
	data := utils.ListOfNumbers(lines[0], " ")
	return applyRulesNTimes4All(data, 75)
}

func applyRulesAllNTimes(stones []int, n int) []int {
	out := stones
	for i := 0; i < n; i++ {
		out = applyRulesAll(out)
	}
	return out
}

func applyRulesNTimes4All(stones []int, n int) int {
	total := 0
	for _, stone := range stones {
		total += applyRulesNTimes(stone, n)
	}
	return total
}

func applyRulesNTimes(stone, n int) int {
	if x, ok := getMemo(stone, n); ok {
		return x
	}
	stones := applyRules(stone)
	if n == 1 {
		setMemo(stone, n, len(stones))
		return len(stones)
	}
	total := 0
	for _, s := range stones {
		total += applyRulesNTimes(s, n-1)
	}
	setMemo(stone, n, total)
	return total
}

func applyRulesAll(stones []int) []int {
	out := make([]int, 0)
	for _, stone := range stones {
		out = append(out, applyRules(stone)...)
	}
	return out
}

func applyRules(stone int) []int {
	if stone == 0 {
		return []int{1}
	}
	str := strconv.Itoa(stone)
	if len(str)%2 == 0 {
		n := len(str) / 2
		a, b := utils.ParseInt(str[:n]), utils.ParseInt(str[n:])
		return []int{a, b}
	}
	return []int{stone * 2024}
}

var memo = map[int]map[int]int{}

func getMemo(a, n int) (int, bool) {
	if _, ok := memo[a]; !ok {
		memo[a] = map[int]int{}
	}
	if v, ok := memo[a][n]; ok {
		return v, true
	}
	return -1, false
}

func setMemo(a, n, v int) {
	if _, ok := memo[a]; !ok {
		memo[a] = map[int]int{}
	}
	memo[a][n] = v
}
