package main

import (
	"advent/utils"
	"fmt"
	"strconv"
)

const day = "22"

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
		init := utils.ParseInt(line)
		total += nextNthSecret(init, 2000)
	}
	return total
}

func part2(lines []string) int {
	m := make(memory)
	buyers := make([][]data, len(lines))
	for i, line := range lines {
		init := utils.ParseInt(line)
		buyers[i] = allNSecret(init, 2000, i, m)
	}
	maxB := 0
	for i := 9; i >= 0; i-- {
		for _, c := range m[i] {
			sliceP := buyers[c.buyer][c.rank-4 : c.rank]
			pattern := [4]int{sliceP[0].diff, sliceP[1].diff, sliceP[2].diff, sliceP[3].diff}
			bns := bananas(pattern, buyers, maxB)
			//fmt.Printf("Pattern %d gives %d\n", pattern, bns)
			if bns > maxB {
				fmt.Printf(utils.Red+"New max : %d with pattern %d\n"+utils.Reset, bns, pattern)
				maxB = bns
			}
		}
	}
	return maxB
}

type data struct {
	n, digit, diff int
}
type coords struct {
	buyer, rank int
}
type memory map[int][]coords

func (m memory) set(price, buyer, rank int) {
	v, ok := m[price]
	if !ok {
		v = make([]coords, 0)
		m[price] = v
	}
	v = append(v, coords{buyer, rank})
	m[price] = v
}

func bananas(pattern [4]int, buyers [][]data, maxB int) int {
	total := 0
	nb := len(buyers)
	for i, buyer := range buyers {
		expect := total + 9*(nb-i)
		if expect <= maxB {
			// kill switch
			return -1
		}
		total += findPattern(pattern, buyer)
	}
	return total
}

func findPattern(pattern [4]int, buyer []data) int {
	for i := 3; i < len(buyer); i++ {
		if buyer[i-3].diff == pattern[0] &&
			buyer[i-2].diff == pattern[1] &&
			buyer[i-1].diff == pattern[2] &&
			buyer[i].diff == pattern[3] {
			return buyer[i].digit
		}
	}
	return 0
}

func nextSecret(secret int) int {
	out := secret
	a := out * 64
	out = out ^ a
	out = out % 16777216
	b := out / 32
	out = out ^ b
	out = out % 16777216
	c := out * 2048
	out = out ^ c
	out = out % 16777216
	return out
}

func nextNthSecret(secret int, n int) int {
	out := secret
	for i := 0; i < n; i++ {
		out = nextSecret(out)
	}
	return out
}

func allNSecret(secret int, n, buyer int, m memory) []data {
	nxt := data{n: secret, digit: secret % 10, diff: 0}
	out := []data{nxt}
	for i := 0; i < n; i++ {
		nxtN := nextSecret(nxt.n)
		d := nxtN % 10
		diff := d - nxt.digit
		nxt = data{n: nxtN, digit: d, diff: diff}
		out = append(out, nxt)
		if len(out) >= 5 {
			m.set(d, buyer, len(out)-1)
		}
	}
	return out
}
