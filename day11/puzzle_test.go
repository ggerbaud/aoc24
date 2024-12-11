package main

import "testing"

func TestPart1(t *testing.T) {
	lines := testData()
	result := part1(lines)
	expect := 55312
	if result != expect {
		t.Fatalf("Part1 returns %d, we expect %d", result, expect)
	}
}

func TestPart2(t *testing.T) {
	lines := testData()
	result := part2(lines)
	expect := 0
	if result != expect {
		t.Fatalf("Part2 returns %d, we expect %d", result, expect)
	}
}

func TestPartApplyRulesAllNTimes(t *testing.T) {
	data := []int{125, 17}
	result := applyRulesNTimes4All(data, 6)
	expect := 22
	if result != expect {
		t.Fatalf("ApplyRulesAllNTimes returns %d, we expect %d", result, expect)
	}
}

func TestPartApplyRulesAllNTimes2(t *testing.T) {
	data := []int{125, 17}
	result := applyRulesNTimes4All(data, 1)
	expect := 3
	if result != expect {
		t.Fatalf("ApplyRulesAllNTimes returns %d, we expect %d", result, expect)
	}
}

func testData() []string {
	return []string{"125 17"}
}
