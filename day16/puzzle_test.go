package main

import "testing"

func TestPart1_1(t *testing.T) {
	lines := testData()
	result := part1(lines)
	expect := 7036
	if result != expect {
		t.Fatalf("Part1 returns %d, we expect %d", result, expect)
	}
}

func TestPart1_2(t *testing.T) {
	lines := testData2()
	result := part1(lines)
	expect := 11048
	if result != expect {
		t.Fatalf("Part1 returns %d, we expect %d", result, expect)
	}
}

func TestPart2_1(t *testing.T) {
	lines := testData()
	result := part2(lines)
	expect := 45
	if result != expect {
		t.Fatalf("Part2 returns %d, we expect %d", result, expect)
	}
}

func TestPart2_2(t *testing.T) {
	lines := testData2()
	result := part2(lines)
	expect := 64
	if result != expect {
		t.Fatalf("Part2 returns %d, we expect %d", result, expect)
	}
}

func testData() []string {
	return []string{
		"###############",
		"#.......#....E#",
		"#.#.###.#.###.#",
		"#.....#.#...#.#",
		"#.###.#####.#.#",
		"#.#.#.......#.#",
		"#.#.#####.###.#",
		"#...........#.#",
		"###.#.#####.#.#",
		"#...#.....#.#.#",
		"#.#.#.###.#.#.#",
		"#.....#...#.#.#",
		"#.###.#.#.#.#.#",
		"#S..#.....#...#",
		"###############",
	}
}

func testData2() []string {
	return []string{
		"#################",
		"#...#...#...#..E#",
		"#.#.#.#.#.#.#.#.#",
		"#.#.#.#...#...#.#",
		"#.#.#.#.###.#.#.#",
		"#...#.#.#.....#.#",
		"#.#.#.#.#.#####.#",
		"#.#...#.#.#.....#",
		"#.#.#####.#.###.#",
		"#.#.#.......#...#",
		"#.#.###.#####.###",
		"#.#.#...#.....#.#",
		"#.#.#.#####.###.#",
		"#.#.#.........#.#",
		"#.#.#.#########.#",
		"#S#.............#",
		"#################",
	}
}
