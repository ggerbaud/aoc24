package main

import (
	"advent/utils"
	"fmt"
	"strconv"
)

const day = "25"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	keys, locks := makeData(lines)
	total := 0
	for _, key := range keys {
		for _, lock := range locks {
			if compatible(key, lock) {
				total++
			}
		}
	}
	return total
}

func part2(lines []string) int {
	return 0
}

func makeData(lines []string) ([][]int, [][]int) {
	fl := true
	keys := make([][]int, 0)
	locks := make([][]int, 0)
	isLock := true
	var data []int
	for _, line := range lines {
		if len(line) == 0 {
			if isLock {
				locks = append(locks, data)
			} else {
				for i := range data {
					data[i]--
				}
				keys = append(keys, data)
			}
			fl = true
			continue
		}
		if fl {
			isLock = line[0] == '#'
			data = make([]int, len(line))
			fl = false
			continue
		}
		for i, c := range line {
			if c == '#' {
				data[i] += 1
			}
		}
	}
	if isLock {
		locks = append(locks, data)
	} else {
		for i := range data {
			data[i]--
		}
		keys = append(keys, data)
	}
	return keys, locks
}

func compatible(key, lock []int) bool {
	for i := 0; i < len(key); i++ {
		if key[i]+lock[i] > 5 {
			return false
		}
	}
	return true
}
