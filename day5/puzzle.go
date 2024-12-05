package main

import (
	"advent/utils"
	"fmt"
	"sort"
	"strconv"
)

const day = "5"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	pages, cpt := makePages(lines)
	lines = lines[cpt:]
	updates := make([][]int, 0)
	for _, updateStr := range lines {
		update := utils.ListOfNumbers(updateStr, ",")
		if verifyUpdate(update, pages) {
			updates = append(updates, update)
		}
	}
	total := 0
	for _, update := range updates {
		total += update[len(update)/2]
	}
	return total
}

func part2(lines []string) int {
	pages, cpt := makePages(lines)
	lines = lines[cpt:]
	updates := make([][]int, 0)
	for _, updateStr := range lines {
		update := utils.ListOfNumbers(updateStr, ",")
		if !verifyUpdate(update, pages) {
			updates = append(updates, update)
		}
	}
	total := 0
	for _, update := range updates {
		fixUpdate(update, pages)
		total += update[len(update)/2]
	}
	return total
}

func makePages(lines []string) (map[int]page, int) {
	pages := make(map[int]page)
	cpt := 1
	line := lines[0]
	for len(line) > 0 {
		beforeAfter := utils.ListOfNumbers(line, "|")
		bef, aft := beforeAfter[0], beforeAfter[1]
		befP, ok := pages[bef]
		if !ok {
			befP = newPage()
			pages[bef] = befP
		}
		aftP, ok := pages[aft]
		if !ok {
			aftP = newPage()
			pages[aft] = aftP
		}
		befP.canBeAfter[aft] = true
		aftP.canBeAfter[bef] = false
		cpt++
		lines = lines[1:]
		line = lines[0]
	}
	return pages, cpt
}

func verifyUpdate(update []int, pages map[int]page) bool {
	for i := 0; i < len(update); i++ {
		p := update[i]
		for j := i + 1; j < len(update); j++ {
			b, ok := pages[p].canBeAfter[update[j]]
			if ok && !b {
				return false
			}
		}
	}
	return true
}

func fixUpdate(update []int, pages map[int]page) {
	sort.SliceStable(update, func(i, j int) bool {
		b, ok := pages[update[i]].canBeAfter[update[j]]
		return !ok || b
	})
}

type page struct {
	canBeAfter map[int]bool
}

func newPage() page {
	return page{make(map[int]bool)}
}
