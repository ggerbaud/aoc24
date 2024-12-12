package main

import "testing"

func TestPart1(t *testing.T) {
	lines := testData()
	result := part1(lines)
	expect := 1930
	if result != expect {
		t.Fatalf("Part1 returns %d, we expect %d", result, expect)
	}
}

func TestPart1_2(t *testing.T) {
	lines := testData2()
	result := part1(lines)
	expect := 140
	if result != expect {
		t.Fatalf("Part1 returns %d, we expect %d", result, expect)
	}
}

func TestPart1_3(t *testing.T) {
	lines := testData3()
	result := part1(lines)
	expect := 772
	if result != expect {
		t.Fatalf("Part1 returns %d, we expect %d", result, expect)
	}
}

func TestPart2(t *testing.T) {
	lines := testData()
	result := part2(lines)
	expect := 1206
	if result != expect {
		t.Fatalf("Part2 returns %d, we expect %d", result, expect)
	}
}

func TestPart2_2(t *testing.T) {
	lines := testData2()
	result := part2(lines)
	expect := 80
	if result != expect {
		t.Fatalf("Part2 returns %d, we expect %d", result, expect)
	}
}

func TestPart2_3(t *testing.T) {
	lines := testData3()
	result := part2(lines)
	expect := 436
	if result != expect {
		t.Fatalf("Part2 returns %d, we expect %d", result, expect)
	}
}

func TestPart2_4(t *testing.T) {
	lines := testData4()
	result := part2(lines)
	expect := 236
	if result != expect {
		t.Fatalf("Part2 returns %d, we expect %d", result, expect)
	}
}

func TestPart2_5(t *testing.T) {
	lines := testData5()
	result := part2(lines)
	expect := 368
	if result != expect {
		t.Fatalf("Part2 returns %d, we expect %d", result, expect)
	}
}

func testData() []string {
	return []string{
		"RRRRIICCFF",
		"RRRRIICCCF",
		"VVRRRCCFFF",
		"VVRCCCJFFF",
		"VVVVCJJCFE",
		"VVIVCCJJEE",
		"VVIIICJJEE",
		"MIIIIIJJEE",
		"MIIISIJEEE",
		"MMMISSJEEE",
	}
}

func testData2() []string {
	return []string{
		"AAAA",
		"BBCD",
		"BBCC",
		"EEEC",
	}
}

func testData3() []string {
	return []string{
		"OOOOO",
		"OXOXO",
		"OOOOO",
		"OXOXO",
		"OOOOO",
	}
}

func testData4() []string {
	return []string{
		"EEEEE",
		"EXXXX",
		"EEEEE",
		"EXXXX",
		"EEEEE",
	}
}

func testData5() []string {
	return []string{
		"AAAAAA",
		"AAABBA",
		"AAABBA",
		"ABBAAA",
		"ABBAAA",
		"AAAAAA",
	}
}
