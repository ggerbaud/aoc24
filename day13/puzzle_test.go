package main

import "testing"

func TestPart1(t *testing.T) {
	lines := testData()
	result := part1(lines)
	expect := 480
	if result != expect {
		t.Fatalf("Part1 returns %d, we expect %d", result, expect)
	}
}

func TestPart2(t *testing.T) {
	lines := testData()
	result := part2(lines)
	expect := 875318608908
	if result != expect {
		t.Fatalf("Part2 returns %d, we expect %d", result, expect)
	}
}

func TestSolve1_1(t *testing.T) {
	p := problem{26, 36, 48, 11, 1946, 1220}
	a, b, ok := p.solve(100)
	ea, eb, eok := 0, 0, false
	if ok != eok || ea != a || eb != b {
		t.Fatalf("Solve1_1 returns %d, %d, %v, we expect %d, %d, %v", a, b, ok, ea, eb, eok)
	}
}

func testData() []string {
	return []string{
		"Button A: X+94, Y+34",
		"Button B: X+22, Y+67",
		"Prize: X=8400, Y=5400",
		"",
		"Button A: X+26, Y+66",
		"Button B: X+67, Y+21",
		"Prize: X=12748, Y=12176",
		"",
		"Button A: X+17, Y+86",
		"Button B: X+84, Y+37",
		"Prize: X=7870, Y=6450",
		"",
		"Button A: X+69, Y+23",
		"Button B: X+27, Y+71",
		"Prize: X=18641, Y=10279",
	}
}
