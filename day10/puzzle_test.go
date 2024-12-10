package main

import "testing"

func TestPart1(t *testing.T) {
	lines := testData()
	result := part1(lines)
	expect := 36
	if result != expect {
		t.Fatalf("Part1 returns %d, we expect %d", result, expect)
	}
}

func TestPart2(t *testing.T) {
	lines := testData()
	result := part2(lines)
	expect := 81
	if result != expect {
		t.Fatalf("Part2 returns %d, we expect %d", result, expect)
	}
}

func TestPart2_2(t *testing.T) {
	lines := testData2()
	result := part2(lines)
	expect := 3
	if result != expect {
		t.Fatalf("Part2_2 returns %d, we expect %d", result, expect)
	}
}

func TestPart2_3(t *testing.T) {
	lines := testData3()
	result := part2(lines)
	expect := 13
	if result != expect {
		t.Fatalf("Part2_3 returns %d, we expect %d", result, expect)
	}
}

func testData() []string {
	return []string{
		"89010123",
		"78121874",
		"87430965",
		"96549874",
		"45678903",
		"32019012",
		"01329801",
		"10456732",
	}
}

func testData2() []string {
	return []string{
		"8888403",
		"8843218",
		"8858828",
		"8865438",
		"6673848",
		"8887658",
		"8893888",
	}
}

func testData3() []string {
	return []string{
		"8890779",
		"8881798",
		"8882837",
		"6543456",
		"7658987",
		"8768333",
		"9873888",
	}
}
