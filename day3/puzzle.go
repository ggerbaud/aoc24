package main

import (
	"advent/utils"
	"fmt"
	"regexp"
	"strconv"
)

const day = "3"

var mulPattern = regexp.MustCompile(`mul\((\d+),(\d+)\)`)
var doPattern = regexp.MustCompile(`do\(\)`)
var dontPattern = regexp.MustCompile(`don't\(\)`)

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
		code := line
		allMult := mulPattern.FindAllString(code, -1)
		for _, mult := range allMult {
			total += calcMult(mult)
		}
	}
	return total
}

func part2(lines []string) int {
	total := 0
	enabled := true
	for _, line := range lines {
		code := line
		for len(code) > 0 {
			if enabled {
				mulLoc := mulPattern.FindStringIndex(code)
				if mulLoc == nil {
					break
				}
				dontLoc := dontPattern.FindStringIndex(code)
				if dontLoc != nil {
					if dontLoc[0] < mulLoc[0] {
						enabled = false
						code = code[dontLoc[1]:]
						continue
					}
				}
				total += calcMult(code[mulLoc[0]:mulLoc[1]])
				code = code[mulLoc[1]:]
			} else {
				doLoc := doPattern.FindStringIndex(code)
				if doLoc == nil {
					break
				}
				code = code[doLoc[1]:]
				enabled = true
			}
		}
	}
	return total
}

func calcMult(mult string) int {
	matches := mulPattern.FindStringSubmatch(mult)
	if len(matches) != 3 {
		return 0
	}
	a := utils.ParseInt(matches[1])
	b := utils.ParseInt(matches[2])
	return a * b
}
