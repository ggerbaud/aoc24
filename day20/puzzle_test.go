package main

import "testing"

func TestPart1_1(t *testing.T) {
	lines := testData()
	result := part(lines, 2, 20)
	expect := 5
	if result != expect {
		t.Fatalf("Part1 returns %d, we expect %d", result, expect)
	}
}

func TestPart1_2(t *testing.T) {
	lines := testData()
	result := part(lines, 2, 10)
	expect := 10
	if result != expect {
		t.Fatalf("Part1 returns %d, we expect %d", result, expect)
	}
}

func TestPart1_3(t *testing.T) {
	lines := testData()
	result := part(lines, 2, 1)
	expect := 44
	if result != expect {
		t.Fatalf("Part1 returns %d, we expect %d", result, expect)
	}
}

func TestPart2_1(t *testing.T) {
	lines := testData()
	result := part(lines, 20, 75)
	expect := 3
	if result != expect {
		t.Fatalf("Part2 returns %d, we expect %d", result, expect)
	}
}

func TestPart2_2(t *testing.T) {
	lines := testData()
	result := part(lines, 20, 74)
	expect := 7
	if result != expect {
		t.Fatalf("Part2 returns %d, we expect %d", result, expect)
	}
}

func TestPart2_3(t *testing.T) {
	lines := testData()
	result := part(lines, 20, 71)
	expect := 29
	if result != expect {
		t.Fatalf("Part2 returns %d, we expect %d", result, expect)
	}
}

func testData() []string {
	return []string{
		"###############",
		"#...#...#.....#",
		"#.#.#.#.#.###.#",
		"#S#...#.#.#...#",
		"#######.#.#.###",
		"#######.#.#...#",
		"#######.#.###.#",
		"###..E#...#...#",
		"###.#######.###",
		"#...###...#...#",
		"#.#####.#.###.#",
		"#.#...#.#.#...#",
		"#.#.#.#.#.#.###",
		"#...#...#...###",
		"###############",
	}
}
