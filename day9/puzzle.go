package main

import (
	"advent/utils"
	"fmt"
	"strconv"
)

const day = "9"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	data := lines[0]
	blocks := expand(data)
	blocks = compress(blocks)
	return checksum(blocks)
}

func part2(lines []string) int {
	data := lines[0]
	fs := expandFS(data)
	blocks := fs.compressFS()
	return checksum(blocks)
}

func expand(data string) []block {
	blocks := make([]block, 0)
	isFile := true
	id := 0
	for _, d := range data {
		n := int(d - '0')
		x := -1
		if isFile {
			x = id
		}
		for i := 0; i < n; i++ {
			blocks = append(blocks, block{x})
		}
		if isFile {
			id++
		}
		isFile = !isFile
	}
	return blocks
}

func expandFS(data string) filesystem {
	blocks := make([]block, 0)
	files := make(map[int]int)
	isFile := true
	id := 0
	for _, d := range data {
		n := int(d - '0')
		x := -1
		if isFile {
			x = id
		}
		for i := 0; i < n; i++ {
			blocks = append(blocks, block{x})
		}
		if isFile {
			files[id] = n
			id++
		}
		isFile = !isFile
	}
	return filesystem{blocks, files}
}

func compress(blocks []block) []block {
	result := make([]block, 0)
	idx := len(blocks) - 1
	for i, b := range blocks {
		if b.id == -1 {
			result = append(result, blocks[idx])
			idx--
			for blocks[idx].id == -1 {
				idx--
			}
		} else {
			result = append(result, b)
		}
		if i >= idx {
			break
		}
	}
	return result
}

func (fs filesystem) compressFS() []block {
	total := len(fs.blocks)
	i := total - 1
	for i >= 0 {
		b := fs.blocks[i]
		if b.id == -1 {
			i--
			continue
		}
		size := fs.files[b.id]
		for k := 0; k < i; k++ {
			b2 := fs.blocks[k]
			if b2.id != -1 {
				continue
			}
			n := 0
			for x := k; x < i && fs.blocks[x].id == -1; x++ {
				n++
			}
			if n < size {
				k += n - 1
				continue
			}
			for j := 0; j < size; j++ {
				fs.blocks[k+j] = block{b.id}
				fs.blocks[i-j] = block{-1}
			}
			break
		}
		i -= size
	}
	return fs.blocks
}

func checksum(blocks []block) int {
	sum := 0
	for i, b := range blocks {
		if b.id != -1 {
			sum += i * b.id
		}
	}
	return sum
}

func toString(blocks []block) string {
	result := ""
	for _, b := range blocks {
		if b.id == -1 {
			result += "."
		} else {
			result += strconv.Itoa(b.id)
		}
	}
	return result
}

type block struct {
	id int
}

type filesystem struct {
	blocks []block
	files  map[int]int
}
