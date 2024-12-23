package main

import (
	"advent/utils"
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
)

const day = "23"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	str := part2(lines)
	fmt.Println("Day#" + day + ".2 : " + str)
}

func part1(lines []string) int {
	net := make(network)
	for _, line := range lines {
		parts := strings.Split(line, "-")
		c1 := net.getOrMake(parts[0])
		c2 := net.getOrMake(parts[1])
		c1.connect(c2)
		c2.connect(c1)
	}
	threeSomes := make(map[string]struct{})
	for _, c := range net {
		if len(c.lan) >= 2 {
			for i := 0; i < len(c.lan)-1; i++ {
				for j := i + 1; j < len(c.lan); j++ {
					if slices.Contains(c.lan[i].lan, c.lan[j]) && slices.Contains(c.lan[j].lan, c.lan[i]) {
						all3 := []string{c.name, c.lan[i].name, c.lan[j].name}
						sort.Strings(all3)
						combo := all3[0] + "-" + all3[1] + "-" + all3[2]
						threeSomes[combo] = struct{}{}
					}
				}
			}
		}
	}
	//fmt.Printf("total group of 3 : %d\n", len(threeSomes))
	total := 0
	for s := range threeSomes {
		//fmt.Println(s)
		if strings.HasPrefix(s, "t") || strings.Contains(s, "-t") {
			total++
		}
	}
	return total
}

func part2(lines []string) string {
	net := make(network)
	for _, line := range lines {
		parts := strings.Split(line, "-")
		c1 := net.getOrMake(parts[0])
		c2 := net.getOrMake(parts[1])
		c1.connect(c2)
		c2.connect(c1)
	}

	cList := make([]*computer, 0, len(net))

	for _, c := range net {
		cList = append(cList, c)
	}

	var maxLenClique []string
	var maxLen int

	BronKerbosch([]string{}, cList, []*computer{}, &maxLenClique, &maxLen, net)
	slices.Sort(maxLenClique)

	return strings.Join(maxLenClique, ",")
}

type network map[string]*computer
type computer struct {
	name string
	lan  []*computer
}

func (c *computer) connect(c2 *computer) {
	if !slices.Contains(c.lan, c2) {
		c.lan = append(c.lan, c2)
	}
}

func (n network) getOrMake(name string) *computer {
	c, ok := n[name]
	if !ok {
		c = &computer{name: name, lan: make([]*computer, 0)}
		n[name] = c
	}
	return c
}

func BronKerbosch(R []string, P, X []*computer, maxLenClique *[]string, maxLen *int, net network) {
	if len(P) == 0 && len(X) == 0 && *maxLen < len(R) {
		rCopy := slices.Clone(R)
		*maxLenClique = rCopy
		*maxLen = len(rCopy)
		return
	}

	pCopy := slices.Clone(P)
	for _, v := range pCopy {
		newR := append(R, v.name)
		neighbours := net[v.name].lan

		newP := intersect(P, neighbours)
		newX := intersect(X, neighbours)

		BronKerbosch(newR, newP, newX, maxLenClique, maxLen, net)

		vIdx := slices.Index(P, v)
		P = slices.Delete(P, vIdx, vIdx+1)

		X = append(X, v)
	}
}

func intersect(a, b []*computer) []*computer {
	t := make(map[*computer]bool)
	for _, val := range a {
		t[val] = true
	}
	out := make([]*computer, 0)
	for _, val := range b {
		if t[val] {
			out = append(out, val)
		}
	}
	return out
}
