package main

import (
	"advent/utils"
	"fmt"
	"strconv"
	"strings"
)

const day = "19"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	towels := strings.Split(lines[0], ", ")
	designs := make([]string, 0)
	for _, line := range lines[2:] {
		designs = append(designs, line)
	}
	//fmt.Printf("%d designs to check with %d towels\n", len(designs), len(towels))
	memory := make(map[string]bool)
	total := 0
	for _, design := range designs {
		//fmt.Printf("Design: %s", design)
		if isValid(design, towels, memory) {
			//fmt.Printf(" - Valid\n")
			total++
		} else {
			//fmt.Printf(" - Invalid\n")
		}
	}
	return total
}

func part2(lines []string) int {
	towels := strings.Split(lines[0], ", ")
	designs := make([]string, 0)
	for _, line := range lines[2:] {
		designs = append(designs, line)
	}
	//fmt.Printf("%d designs to check with %d towels\n", len(designs), len(towels))
	memory := make(map[string]int)
	total := 0
	for _, design := range designs {
		//fmt.Printf("Design: %s", design)
		result := isValid2(design, towels, memory)
		//fmt.Printf(" => %d solutions\n", result)
		total += result
	}
	return total
}

func isValid(design string, towels []string, memory map[string]bool) bool {
	if len(design) == 0 {
		return true
	}
	if v, ok := memory[design]; ok {
		return v
	}
	for _, towel := range towels {
		if strings.HasPrefix(design, towel) && isValid(design[len(towel):], towels, memory) {
			memory[design] = true
			return true
		}
	}
	memory[design] = false
	return false
}

func isValid2(design string, towels []string, memory map[string]int) int {
	if len(design) == 0 {
		return 1
	}
	if v, ok := memory[design]; ok {
		return v
	}
	out := 0
	for _, towel := range towels {
		if strings.HasPrefix(design, towel) {
			out += isValid2(design[len(towel):], towels, memory)
		}
	}
	memory[design] = out
	return out
}
