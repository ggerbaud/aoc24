package main

import "testing"

func TestPart1(t *testing.T) {
	lines := testData()
	result := part1(lines)
	expect := 1928
	if result != expect {
		t.Fatalf("Part1 returns %d, we expect %d", result, expect)
	}
}

func TestPart2(t *testing.T) {
	lines := testData()
	result := part2(lines)
	expect := 2858
	if result != expect {
		t.Fatalf("Part2 returns %d, we expect %d", result, expect)
	}
}

func TestExpand1(t *testing.T) {
	data := "12345"
	result := toString(expand(data))
	expect := "0..111....22222"
	if result != expect {
		t.Fatalf("Expand returns %s, we expect %s", result, expect)
	}
}

func TestExpand2(t *testing.T) {
	data := "2333133121414131402"
	result := toString(expand(data))
	expect := "00...111...2...333.44.5555.6666.777.888899"
	if result != expect {
		t.Fatalf("Expand returns %s, we expect %s", result, expect)
	}
}

func TestExpandAndCompress1(t *testing.T) {
	data := "12345"
	result := toString(compress(expand(data)))
	expect := "022111222."
	if result != expect {
		t.Fatalf("Expand&Compress returns %s, we expect %s", result, expect)
	}
}

func TestExpandAndCompress2(t *testing.T) {
	data := "2333133121414131402"
	result := toString(compress(expand(data)))
	expect := "0099811188827773336446555566"
	if result != expect {
		t.Fatalf("Expand&Compress returns %s, we expect %s", result, expect)
	}
}

func TestExpandFS1(t *testing.T) {
	data := "12345"
	result := toString(expandFS(data).blocks)
	expect := "0..111....22222"
	if result != expect {
		t.Fatalf("Expand returns %s, we expect %s", result, expect)
	}
}

func TestExpandFS2(t *testing.T) {
	data := "2333133121414131402"
	result := toString(expandFS(data).blocks)
	expect := "00...111...2...333.44.5555.6666.777.888899"
	if result != expect {
		t.Fatalf("Expand returns %s, we expect %s", result, expect)
	}
}

func TestExpandAndCompressFS1(t *testing.T) {
	data := "12345"
	result := toString(expandFS(data).compressFS())
	expect := "0..111....22222"
	if result != expect {
		t.Fatalf("Expand&Compress returns %s, we expect %s", result, expect)
	}
}

func TestExpandAndCompressFS2(t *testing.T) {
	data := "2333133121414131402"
	result := toString(expandFS(data).compressFS())
	expect := "00992111777.44.333....5555.6666.....8888.."
	if result != expect {
		t.Fatalf("Expand&Compress returns %s, we expect %s", result, expect)
	}
}

func testData() []string {
	return []string{
		"2333133121414131402",
	}
}
