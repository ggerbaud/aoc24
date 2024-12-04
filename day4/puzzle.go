package main

import (
	"advent/utils"
	"fmt"
	"strconv"
)

const day = "4"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	total := 0
	for i, line := range lines {
		for j, c := range line {
			if c != 'X' {
				continue
			}
			total += getXMAS(lines, i, j)
		}
	}
	return total
}

func part2(lines []string) int {
	total := 0
	for i, line := range lines {
		for j, c := range line {
			if c != 'A' {
				continue
			}
			if isXMAS(lines, i, j) {
				total++
			}
		}
	}
	return total
}

func getXMAS(lines []string, i, j int) int {
	xmas := 0
	if isH(lines, i, j) {
		xmas++
	}
	if isHB(lines, i, j) {
		xmas++
	}
	if isV(lines, i, j) {
		xmas++
	}
	if isVB(lines, i, j) {
		xmas++
	}
	if isDR(lines, i, j) {
		xmas++
	}
	if isDL(lines, i, j) {
		xmas++
	}
	if isDLB(lines, i, j) {
		xmas++
	}
	if isDRB(lines, i, j) {
		xmas++
	}
	return xmas
}
func isH(lines []string, i, j int) bool {
	return checkLetter(lines, i, j+1, 'M') &&
		checkLetter(lines, i, j+2, 'A') &&
		checkLetter(lines, i, j+3, 'S')
}
func isHB(lines []string, i, j int) bool {
	return checkLetter(lines, i, j-1, 'M') &&
		checkLetter(lines, i, j-2, 'A') &&
		checkLetter(lines, i, j-3, 'S')
}
func isV(lines []string, i, j int) bool {
	return checkLetter(lines, i+1, j, 'M') &&
		checkLetter(lines, i+2, j, 'A') &&
		checkLetter(lines, i+3, j, 'S')
}
func isVB(lines []string, i, j int) bool {
	return checkLetter(lines, i-1, j, 'M') &&
		checkLetter(lines, i-2, j, 'A') &&
		checkLetter(lines, i-3, j, 'S')
}
func isDR(lines []string, i, j int) bool {
	return checkLetter(lines, i+1, j+1, 'M') &&
		checkLetter(lines, i+2, j+2, 'A') &&
		checkLetter(lines, i+3, j+3, 'S')
}
func isDLB(lines []string, i, j int) bool {
	return checkLetter(lines, i-1, j-1, 'M') &&
		checkLetter(lines, i-2, j-2, 'A') &&
		checkLetter(lines, i-3, j-3, 'S')
}
func isDL(lines []string, i, j int) bool {
	return checkLetter(lines, i+1, j-1, 'M') &&
		checkLetter(lines, i+2, j-2, 'A') &&
		checkLetter(lines, i+3, j-3, 'S')
}
func isDRB(lines []string, i, j int) bool {
	return checkLetter(lines, i-1, j+1, 'M') &&
		checkLetter(lines, i-2, j+2, 'A') &&
		checkLetter(lines, i-3, j+3, 'S')
}

func isXMAS(lines []string, i, j int) bool {
	ltr := (checkLetter(lines, i-1, j-1, 'M') && checkLetter(lines, i+1, j+1, 'S')) ||
		(checkLetter(lines, i-1, j-1, 'S') && checkLetter(lines, i+1, j+1, 'M'))
	if ltr {
		return (checkLetter(lines, i-1, j+1, 'M') && checkLetter(lines, i+1, j-1, 'S')) ||
			(checkLetter(lines, i-1, j+1, 'S') && checkLetter(lines, i+1, j-1, 'M'))
	}
	return false
}

func checkLetter(lines []string, i, j int, expected uint8) bool {
	if i < 0 || j < 0 || len(lines) <= i || len(lines[i]) <= j {
		return false
	}
	return lines[i][j] == expected
}
