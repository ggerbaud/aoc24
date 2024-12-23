package main

import (
	"advent/utils"
	"fmt"
	"strconv"
)

const day = "21"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	// 134341709499296
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	total := 0
	for _, line := range lines {
		code := []rune(line)
		total += complexity(code, 2, gmem)
	}
	return total
}

func part2(lines []string) int {
	total := 0
	for _, line := range lines {
		code := []rune(line)
		total += complexity(code, 25, gmem)
	}
	return total
}

func complexity(code []rune, dir int, m memory) int {
	seq := decode(code, dir, m)
	return seq * numPart(code)
}

func numPart(code []rune) int {
	str := ""
	for _, r := range code {
		if r > '9' || r < '0' {
			continue
		}
		if r == '0' {
			if len(str) > 0 {
				str += "0"
			}
		} else {
			str += string(r)
		}
	}
	return utils.ParseInt(str)
}

func decodeOld(code []rune, dir int) []rune {
	//fmt.Printf("decode %s (%d)\n", string(code), len(code))
	result := decodeNum(code)
	for i := 0; i < dir; i++ {
		result = decodeDir(result)
	}
	//fmt.Printf("result %s (%d)\n", string(result), len(result))
	return result
}

func decode(code []rune, dir int, m memory) int {
	c := 'A'
	total := 0
	for _, r := range code {
		result := numFromTo(c, r)
		result += "A"
		c2 := 'A'
		for _, r2 := range result {
			total += decodeUnit(c2, r2, dir, m)
			c2 = r2
		}
		c = r
	}
	return total
}

func decodeUnit(from, to rune, level int, m memory) int {
	result, ok := m.get(from, to, level)
	if ok {
		return result
	}
	steps := dirFromTo(from, to)
	steps = append(steps, 'A')
	if level == 1 {
		res := len(steps)
		m.set(from, to, 1, res)
		return res
	}
	total := 0
	c := 'A'
	for _, step := range steps {
		total += decodeUnit(c, step, level-1, m)
		c = step
	}
	m.set(from, to, level, total)
	return total
}

func decodeDir(instr []rune) []rune {
	//fmt.Printf("decodeDir %s (%d)\n", string(instr), len(instr))
	start := 'A'
	out := make([]rune, 0)
	for _, r := range instr {
		out = append(out, dirFromTo(start, r)...)
		out = append(out, 'A')
		start = r
	}
	return out
}

func dirFromTo(from, to rune) []rune {
	switch from {
	case 'A':
		switch to {
		case '>':
			return []rune{'v'}
		case '<':
			return []rune("v<<")
		case '^':
			return []rune{'<'}
		case 'v':
			return []rune("v<")
		}
	case '>':
		switch to {
		case 'v':
			return []rune{'<'}
		case '<':
			return []rune("<<")
		case '^':
			return []rune{'<', '^'}
		case 'A':
			return []rune{'^'}
		}
	case '<':
		switch to {
		case 'v':
			return []rune{'>'}
		case '>':
			return []rune(">>")
		case 'A':
			return []rune(">>^")
		case '^':
			return []rune(">^")
		}
	case '^':
		switch to {
		case 'v':
			return []rune{'v'}
		case '>':
			return []rune(">v")
		case '<':
			return []rune("v<")
		case 'A':
			return []rune{'>'}
		}
	case 'v':
		switch to {
		case '>':
			return []rune{'>'}
		case '<':
			return []rune("<")
		case '^':
			return []rune{'^'}
		case 'A':
			return []rune(">^")
		}
	}
	return []rune{}
}

func backwardNum(instr []rune) ([]rune, bool) {
	//fmt.Printf("backwardNum %s (%d)\n", string(instr), len(instr))
	pt := utils.Point{X: 2, Y: 3}
	out := make([]rune, 0)
	for _, r := range instr {
		switch r {
		case '>':
			pt.X++
		case '<':
			pt.X--
		case '^':
			pt.Y--
		case 'v':
			pt.Y++
		case 'A':
			out = append(out, numpad[pt.Y][pt.X])
		}
		if pt.Y < 0 || pt.Y >= len(numpad) || pt.X < 0 || pt.X >= len(numpad[pt.Y]) || numpad[pt.Y][pt.X] == 0 {
			return nil, false
		}
	}
	return out, true
}

func decodeNum(code []rune) []rune {
	//fmt.Printf("decodeNum %s (%d)\n", string(code), len(code))
	start := 'A'
	out := make([]rune, 0)
	for _, r := range code {
		out = append(out, []rune(numFromTo(start, r))...)
		out = append(out, 'A')
		start = r
	}
	return out
}

func numFromTo(from, to rune) string {
	switch from {
	case 'A':
		return numATo(to)
	case '0':
		return num0To(to)
	case '1':
		return num1To(to)
	case '2':
		return num2To(to)
	case '3':
		return num3To(to)
	case '4':
		return num4To(to)
	case '5':
		return num5To(to)
	case '6':
		return num6To(to)
	case '7':
		return num7To(to)
	case '8':
		return num8To(to)
	case '9':
		return num9To(to)
	}
	return ""
}

func numATo(to rune) string {
	switch to {
	case '3':
		return "^"
	case '0':
		return "<"
	case '1':
		return "^<<"
	case '2':
		return "<^"
	case '4':
		return "^^<<"
	case '5':
		return "<^^"
	case '6':
		return "^^"
	case '7':
		return "^^^<<"
	case '8':
		return "<^^^"
	case '9':
		return "^^^"
	}
	return ""
}

func num0To(to rune) string {
	switch to {
	case 'A':
		return ">"
	case '1':
		return "^<"
	case '2':
		return "^"
	case '3':
		return "^>"
	case '4':
		return "^^<"
	case '5':
		return "^^"
	case '6':
		return "^^>"
	case '7':
		return "^^^<"
	case '8':
		return "^^^"
	case '9':
		return "^^^>"
	}
	return ""
}

func num1To(to rune) string {
	switch to {
	case 'A':
		return ">>v"
	case '0':
		return ">v"
	case '2':
		return ">"
	case '3':
		return ">>"
	case '4':
		return "^"
	case '5':
		return "^>"
	case '6':
		return "^>>"
	case '7':
		return "^^"
	case '8':
		return "^^>"
	case '9':
		return "^^>>"
	}
	return ""
}

func num2To(to rune) string {
	switch to {
	case 'A':
		return "v>"
	case '0':
		return "v"
	case '1':
		return "<"
	case '3':
		return ">"
	case '4':
		return "<^"
	case '5':
		return "^"
	case '6':
		return "^>"
	case '7':
		return "<^^"
	case '8':
		return "^^"
	case '9':
		return "^^>"
	}
	return ""
}

func num3To(to rune) string {
	switch to {
	case 'A':
		return "v"
	case '0':
		return "<v"
	case '1':
		return "<<"
	case '2':
		return "<"
	case '4':
		return "<<^"
	case '5':
		return "<^"
	case '6':
		return "^"
	case '7':
		return "<<^^"
	case '8':
		return "<^^"
	case '9':
		return "^^"
	}
	return ""
}

func num4To(to rune) string {
	switch to {
	case 'A':
		return ">>vv"
	case '0':
		return ">vv"
	case '1':
		return "v"
	case '2':
		return "v>"
	case '3':
		return "v>>"
	case '5':
		return ">"
	case '6':
		return ">>"
	case '7':
		return "^"
	case '8':
		return "^>"
	case '9':
		return "^>>"
	}
	return ""
}

func num5To(to rune) string {
	switch to {
	case 'A':
		return "vv>"
	case '0':
		return "vv"
	case '1':
		return "<v"
	case '2':
		return "v"
	case '3':
		return "v>"
	case '4':
		return "<"
	case '6':
		return ">"
	case '7':
		return "<^"
	case '8':
		return "^"
	case '9':
		return "^>"
	}
	return ""
}

func num6To(to rune) string {
	switch to {
	case 'A':
		return "vv"
	case '0':
		return "<vv"
	case '1':
		return "<<v"
	case '2':
		return "<v"
	case '3':
		return "v"
	case '4':
		return "<<"
	case '5':
		return "<"
	case '7':
		return "<<^"
	case '8':
		return "<^"
	case '9':
		return "^"
	}
	return ""
}

func num7To(to rune) string {
	switch to {
	case 'A':
		return ">>vvv"
	case '0':
		return ">vvv"
	case '1':
		return "vv"
	case '2':
		return "vv>"
	case '3':
		return "vv>>"
	case '4':
		return "v"
	case '5':
		return "v>"
	case '6':
		return "v>>"
	case '8':
		return ">"
	case '9':
		return ">>"
	}
	return ""
}

func num8To(to rune) string {
	switch to {
	case 'A':
		return "vvv>"
	case '0':
		return "vvv"
	case '1':
		return "<vv"
	case '2':
		return "vv"
	case '3':
		return "vv>"
	case '4':
		return "<v"
	case '5':
		return "v"
	case '6':
		return "v>"
	case '7':
		return "<"
	case '9':
		return ">"
	}
	return ""
}

func num9To(to rune) string {
	switch to {
	case 'A':
		return "vvv"
	case '0':
		return "<vvv"
	case '1':
		return "<<vv"
	case '2':
		return "<vv"
	case '3':
		return "vv"
	case '4':
		return "<<v"
	case '5':
		return "<v"
	case '6':
		return "v"
	case '7':
		return "<<"
	case '8':
		return "<"
	}
	return ""
}

var (
	gmem   = make(memory)
	numpad = [][]rune{{'7', '8', '9'}, {'4', '5', '6'}, {'1', '2', '3'}, {0, '0', 'A'}}
	dirpad = [][]rune{{0, '^', 'A'}, {'<', 'v', '>'}}
)

type memory map[rune]map[rune]map[int]int

func (m memory) set(from, to rune, level int, result int) {
	tomap, ok := m[from]
	if !ok {
		tomap = make(map[rune]map[int]int)
		m[from] = tomap
	}
	levelmap, ok := tomap[to]
	if !ok {
		levelmap = make(map[int]int)
		tomap[to] = levelmap
	}
	levelmap[level] = result
}

func (m memory) get(from, to rune, level int) (int, bool) {
	tomap, ok := m[from]
	if !ok {
		return -1, false
	}
	levelmap, ok := tomap[to]
	if !ok {
		return -1, false
	}
	result, ok := levelmap[level]
	return result, ok
}
