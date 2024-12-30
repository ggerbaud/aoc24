package main

import (
	"advent/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const day = "24"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	result := part2(lines)
	fmt.Println("Day#" + day + ".2 : " + result)
}

func part1(lines []string) int {
	wires := make(map[string]bool)
	k := 0
	for i, line := range lines {
		if len(line) == 0 {
			k = i + 1
			break
		}
		parts := strings.Split(line, ":")
		wires[parts[0]] = utils.ParseInt(parts[1]) == 1
	}
	logicals := make([]door, 0)
	for i := k; i < len(lines); i++ {
		parts := strings.Split(lines[i], "->")
		output := strings.TrimSpace(parts[1])
		d := door{out: output}
		if strings.Contains(parts[0], "AND") {
			ws := strings.Split(strings.TrimSpace(parts[0]), " AND ")
			d.w1, d.w2 = ws[0], ws[1]
			d.kind = "AND"
		} else if strings.Contains(parts[0], "XOR") {
			ws := strings.Split(strings.TrimSpace(parts[0]), " XOR ")
			d.w1, d.w2 = ws[0], ws[1]
			d.kind = "XOR"
		} else if strings.Contains(parts[0], " OR ") {
			ws := strings.Split(strings.TrimSpace(parts[0]), " OR ")
			d.w1, d.w2 = ws[0], ws[1]
			d.kind = "OR"
		}
		logicals = append(logicals, d)
	}
	for len(logicals) > 0 {
		h := logicals[0]
		logicals = logicals[1:]
		if !h.calc(wires) {
			logicals = append(logicals, h)
		}
	}
	result := 0
	for n, v := range wires {
		if n[0] == 'z' {
			if v {
				idx := utils.ParseInt(n[1:])
				mask := 1 << idx
				result = result | mask
			}
		}
	}
	return result
}

func part2(lines []string) string {
	wires := make(map[string]bool)
	k := 0
	for i, line := range lines {
		if len(line) == 0 {
			k = i + 1
			break
		}
		parts := strings.Split(line, ":")
		wires[parts[0]] = utils.ParseInt(parts[1]) == 1
	}
	logicals := make([]door, 0)
	for i := k; i < len(lines); i++ {
		parts := strings.Split(lines[i], "->")
		output := strings.TrimSpace(parts[1])
		d := door{out: output}
		if strings.Contains(parts[0], "AND") {
			ws := strings.Split(strings.TrimSpace(parts[0]), " AND ")
			d.w1, d.w2 = ws[0], ws[1]
			d.kind = "AND"
		} else if strings.Contains(parts[0], "XOR") {
			ws := strings.Split(strings.TrimSpace(parts[0]), " XOR ")
			d.w1, d.w2 = ws[0], ws[1]
			d.kind = "XOR"
		} else if strings.Contains(parts[0], " OR ") {
			ws := strings.Split(strings.TrimSpace(parts[0]), " OR ")
			d.w1, d.w2 = ws[0], ws[1]
			d.kind = "OR"
		}
		logicals = append(logicals, d)
	}
	// found by hand
	inversions := []string{"hjm", "mcq"}
	logicals = doInvert(logicals, "hjm", "mcq")
	for {
		data := renaming(logicals)
		found, from, to := findInverted(data)
		if found {
			inversions = append(inversions, from)
			inversions = append(inversions, to)
			logicals = doInvert(logicals, from, to)
		} else {
			logicals = data
			break
		}
	}
	//sort.SliceStable(logicals, logicalSorter(logicals))
	//for len(logicals) > 0 {
	//	h := logicals[0]
	//	logicals = logicals[1:]
	//	fmt.Println(h.String())
	//}
	sort.Strings(inversions)
	return strings.Join(inversions, ",")
}

func findInverted(logicals []door) (bool, string, string) {
	n := 0
	for {
		sn := str(n)
		np1 := n + 1
		snp1 := str(np1)
		// V(N) = XOR(N) xor RETAIN(N-1)
		xor := "XOR" + snp1
		rtn := "RETAIN" + sn
		z := "z" + snp1
		found := false
		for _, d := range logicals {
			if d.kind == "XOR" && (d.w1 == xor || d.w2 == xor) && (d.w1 == rtn || d.w2 == rtn) {
				found = true
				if d.out != z {
					return true, d.out, z
				}
				break
			}
		}
		if found {
			n++
		} else {
			break
		}
	}
	return false, "", ""
}

func str(i int) string {
	res := strconv.Itoa(i)
	if i < 10 {
		res = "0" + res
	}
	return res
}

func renaming(logicals []door) []door {
	out := make([]door, len(logicals))
	copy(out, logicals)
	renames := make(map[string]string)
	changed := true
	for changed {
		changed = false
		for i, d := range out {
			if v, ok := renames[d.w1]; ok {
				changed = true
				d.w1 = v
			}
			if v, ok := renames[d.w2]; ok {
				changed = true
				d.w2 = v
			}
			compatible := d.isInput()
			if compatible {
				nn := d.kind + d.value()
				if nn == "XOR00" {
					continue
				}
				if nn == "AND00" {
					nn = "RETAIN00"
				}
				//if strings.HasPrefix(d.out, "z") {
				//	fmt.Printf("%s%s%s\n", utils.Red, d, utils.Reset)
				//}
				if x, ok := renames[d.out]; d.out != nn && (!ok || x != nn) {
					changed = true
					renames[d.out] = nn
					d.out = nn
				}
			} else if d.kind == "AND" && d.isRetainTmp() {
				nn := "RTMP"
				ww1, ww2 := wireValue(d.w1), wireValue(d.w2)
				if ww1 < ww2 {
					nn += ww2
				} else {
					nn += ww1
				}
				//if strings.HasPrefix(d.out, "z") {
				//	fmt.Printf("%s%s%s\n", utils.Red, d, utils.Reset)
				//}
				if x, ok := renames[d.out]; d.out != nn && (!ok || x != nn) {
					changed = true
					renames[d.out] = nn
					d.out = nn
				}
			} else if d.kind == "OR" && d.isRetain() {
				nn := "RETAIN" + wireValue(d.w1)
				//if strings.HasPrefix(d.out, "z") {
				//	fmt.Printf("%s%s%s\n", utils.Red, d, utils.Reset)
				//}
				if x, ok := renames[d.out]; d.out != nn && (!ok || x != nn) {
					changed = true
					renames[d.out] = nn
					d.out = nn
				}
			}
			out[i] = d
		}
	}
	return out
}

func doInvert(logicals []door, from, to string) []door {
	out := make([]door, len(logicals))
	copy(out, logicals)
	for i, d := range logicals {
		if d.out == from {
			d.out = to
			out[i] = d
		} else if d.out == to {
			d.out = from
			out[i] = d
		}
	}
	return out
}

type door struct {
	w1, w2, out string
	kind        string
}

func (d door) isInput() bool {
	if (strings.HasPrefix(d.w1, "x") && strings.HasPrefix(d.w2, "y")) ||
		(strings.HasPrefix(d.w1, "y") && strings.HasPrefix(d.w2, "x")) {
		return wireValue(d.w1) == wireValue(d.w2)
	}
	return false
}

func (d door) isRetainTmp() bool {
	return isRetainTmp(d.w1, d.w2) || isRetainTmp(d.w2, d.w1)
}

func (d door) isRetain() bool {
	return isRetain(d.w1, d.w2) || isRetain(d.w2, d.w1)
}

func (d door) isValue() (bool, bool, string) {
	v, g, to := isValue(d.w1, d.w2, d.out)
	if !v {
		v, g, to = isValue(d.w2, d.w1, d.out)
	}
	return v, g, to
}

func isRetainTmp(w1, w2 string) bool {
	if strings.HasPrefix(w1, "XOR") && strings.HasPrefix(w2, "RETAIN") {
		ww1, ww2 := wireValue(w1), wireValue(w2)
		if ww1 != ww2 {
			i1, err := strconv.Atoi(ww1)
			if err != nil {
				return false
			}
			i2, err := strconv.Atoi(ww2)
			if err != nil {
				return false
			}
			return i1 == i2+1
		}
	}
	return false
}

func isRetain(w1, w2 string) bool {
	if strings.HasPrefix(w1, "AND") && strings.HasPrefix(w2, "RTMP") {
		ww1, ww2 := wireValue(w1), wireValue(w2)
		return ww1 == ww2
	}
	return false
}

func isValue(w1, w2, out string) (bool, bool, string) {
	isOutLabel := strings.HasPrefix(out, "z")
	outV := wireValue(out)
	isFinalXOR := false
	ww1 := wireValue(w1)
	if strings.HasPrefix(w1, "XOR") && strings.HasPrefix(w2, "RETAIN") {
		ww2 := wireValue(w2)
		if ww1 != ww2 {
			i1, err := strconv.Atoi(ww1)
			if err != nil {
				return false, false, ""
			}
			i2, err := strconv.Atoi(ww2)
			if err != nil {
				return false, false, ""
			}
			isFinalXOR = i1 == i2+1
		}
	}
	goodValue := isFinalXOR && isOutLabel && outV == ww1
	return isFinalXOR, goodValue, "z" + ww1
}

func (d door) calc(wires map[string]bool) bool {
	w1, ok1 := wires[d.w1]
	w2, ok2 := wires[d.w2]
	if ok1 && ok2 {
		v := false
		switch d.kind {
		case "AND":
			v = w1 && w2
		case "OR":
			v = w1 || w2
		case "XOR":
			v = w1 != w2
		}
		wires[d.out] = v
		return true
	}
	return false
}

func (d door) String() string {
	return fmt.Sprintf("%s %s %s => %s", d.w1, d.kind, d.w2, d.out)
}

func logicalSorter(logicals []door) func(i, j int) bool {
	return func(i, j int) bool {
		return less(logicals[i], logicals[j])
	}
}

func (d door) value() string {
	ww1, ww2 := wireValue(d.w1), wireValue(d.w2)
	if ww1 == ww2 {
		return ww1
	}
	if ww1 < ww2 {
		return ww1 + ww2
	}
	return ww2 + ww1
}

func wireValue(name string) string {
	return name[len(name)-2:]
}

func less(d1, d2 door) bool {
	m1, m2 := d1.value(), d2.value()
	return m1 < m2
}
