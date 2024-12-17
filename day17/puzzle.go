package main

import (
	"advent/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const day = "17"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	str := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + str)
	total := part2Bis(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) string {
	prog, a, b, c := makeData(lines)
	_, _, _, output := compute(prog, 0, a, b, c, []int{})
	return display(output)
}

func part2(lines []string) int {
	prog, _, _, _ := makeData(lines)
	for i := 0; ; i++ {
		//fmt.Printf("Trying %d\n", i)
		if computeBF(prog, 0, i, 0, 0, []int{}) {
			return i
		}
	}
}

func part2Bis(lines []string) int {
	prog, _, _, _ := makeData(lines)
	result, ok := computeRec(prog, 0)
	if ok {
		return result
	}
	return -1
}

func computeRec(prog []int, x int) (int, bool) {
	for j := 0; j < 8; j++ {
		k := x<<3 + j
		_, _, _, output := compute(prog, 0, k, 0, 0, []int{})
		do := display(output)
		dt := display(prog[len(prog)-len(output):])
		if do == dt {
			if len(output) == len(prog) {
				return k, true
			}
			result, ok := computeRec(prog, k)
			if ok {
				return result, true
			}
		}
	}
	return 0, false
}

func compute(prog []int, pt, a, b, c int, output []int) (int, int, int, []int) {
	if pt >= len(prog)-1 {
		return a, b, c, output
	}
	opcode := prog[pt]
	operand := prog[pt+1]
	switch opcode {
	case adv:
		a = a / int(math.Pow(2, float64(combo(operand, a, b, c))))
		return compute(prog, pt+2, a, b, c, output)
	case bxl:
		b = b ^ operand
		return compute(prog, pt+2, a, b, c, output)
	case bst:
		b = combo(operand, a, b, c) % 8
		return compute(prog, pt+2, a, b, c, output)
	case jnz:
		if a != 0 {
			return compute(prog, operand, a, b, c, output)
		}
		return compute(prog, pt+2, a, b, c, output)
	case bxc:
		b = b ^ c
		return compute(prog, pt+2, a, b, c, output)
	case out:
		output = append(output, combo(operand, a, b, c)%8)
		return compute(prog, pt+2, a, b, c, output)
	case bdv:
		b = a / int(math.Pow(2, float64(combo(operand, a, b, c))))
		return compute(prog, pt+2, a, b, c, output)
	case cdv:
		c = a / int(math.Pow(2, float64(combo(operand, a, b, c))))
		return compute(prog, pt+2, a, b, c, output)
	}
	return a, b, c, output
}

func computeBF(prog []int, pt, a, b, c int, output []int) bool {
	if pt >= len(prog)-1 {
		return len(prog) == len(output)
	}
	opcode := prog[pt]
	operand := prog[pt+1]
	switch opcode {
	case adv:
		a = a / int(math.Pow(2, float64(combo(operand, a, b, c))))
		return computeBF(prog, pt+2, a, b, c, output)
	case bxl:
		b = b ^ operand
		return computeBF(prog, pt+2, a, b, c, output)
	case bst:
		b = combo(operand, a, b, c) % 8
		return computeBF(prog, pt+2, a, b, c, output)
	case jnz:
		if a != 0 {
			return computeBF(prog, operand, a, b, c, output)
		}
		return computeBF(prog, pt+2, a, b, c, output)
	case bxc:
		b = b ^ c
		return computeBF(prog, pt+2, a, b, c, output)
	case out:
		last := combo(operand, a, b, c) % 8
		if prog[len(output)] != last {
			return false
		}
		output = append(output, last)
		return computeBF(prog, pt+2, a, b, c, output)
	case bdv:
		b = a / int(math.Pow(2, float64(combo(operand, a, b, c))))
		return computeBF(prog, pt+2, a, b, c, output)
	case cdv:
		c = a / int(math.Pow(2, float64(combo(operand, a, b, c))))
		return computeBF(prog, pt+2, a, b, c, output)
	}
	return false
}

func combo(op, a, b, c int) int {
	if op == 4 {
		return a
	}
	if op == 5 {
		return b
	}
	if op == 6 {
		return c
	}
	return op
}

func makeData(lines []string) ([]int, int, int, int) {
	a, b, c := 0, 0, 0
	var prog []int
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		splits := strings.Split(line, ":")
		if strings.HasPrefix(line, "Register A") {
			a = utils.ParseInt(strings.TrimSpace(splits[1]))
		} else if strings.HasPrefix(line, "Register B") {
			b = utils.ParseInt(strings.TrimSpace(splits[1]))
		} else if strings.HasPrefix(line, "Register C") {
			c = utils.ParseInt(strings.TrimSpace(splits[1]))
		} else if strings.HasPrefix(line, "Program") {
			prog = utils.ListOfNumbers(strings.TrimSpace(splits[1]), ",")
		}
	}
	return prog, a, b, c
}

func display(output []int) string {
	str := ""
	for _, i := range output {
		str += strconv.Itoa(i) + ","
	}
	if len(str) > 0 {
		str = str[:len(str)-1]
	}
	return str
}

const (
	adv = iota
	bxl
	bst
	jnz
	bxc
	out
	bdv
	cdv
)
