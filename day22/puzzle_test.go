package main

import "testing"

func TestPart1(t *testing.T) {
	lines := testData()
	result := part1(lines)
	expect := 37327623
	if result != expect {
		t.Fatalf("Part1 returns %d, we expect %d", result, expect)
	}
}

func TestPart2(t *testing.T) {
	lines := testData2()
	result := part2(lines)
	expect := 23
	if result != expect {
		t.Fatalf("Part2 returns %d, we expect %d", result, expect)
	}
}

func TestNextSecret(t *testing.T) {
	result := nextNthSecret(123, 1)
	expect := 15887950
	if result != expect {
		t.Fatalf("Part1 returns %d, we expect %d", result, expect)
	}
}

func TestNextSecret2(t *testing.T) {
	result := nextNthSecret(123, 2)
	expect := 16495136
	if result != expect {
		t.Fatalf("Part1 returns %d, we expect %d", result, expect)
	}
}

func TestNextSecret3(t *testing.T) {
	result := nextNthSecret(123, 10)
	expect := 5908254
	if result != expect {
		t.Fatalf("Part1 returns %d, we expect %d", result, expect)
	}
}

func TestNextSecret4(t *testing.T) {
	result := nextNthSecret(1, 2000)
	expect := 8685429
	if result != expect {
		t.Fatalf("Part1 returns %d, we expect %d", result, expect)
	}
}

func testData() []string {
	return []string{
		"1",
		"10",
		"100",
		"2024",
	}
}

func testData2() []string {
	return []string{
		"1",
		"2",
		"3",
		"2024",
	}
}
