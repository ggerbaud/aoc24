package main

import "testing"

func TestPart1(t *testing.T) {
	lines := testData()
	result := part1(lines)
	expect := 126384
	if result != expect {
		t.Fatalf("Part1 returns %d, we expect %d", result, expect)
	}
}

//func TestPart2(t *testing.T) {
//	lines := testData()
//	result := part2(lines)
//	expect := 0
//	if result != expect {
//		t.Fatalf("Part2 returns %d, we expect %d", result, expect)
//	}
//}

func TestDecode1_1(t *testing.T) {
	result := decode([]rune("029A"), 2, make(memory))
	expect := 68
	if result != expect {
		t.Fatalf("Decode returns %d, we expect %d", result, expect)
	}
}

func TestComplexity1_1(t *testing.T) {
	result := complexity([]rune("029A"), 2, make(memory))
	expect := 68 * 29
	if result != expect {
		t.Fatalf("Decode returns %d, we expect %d", result, expect)
	}
}

func TestDecode1_2(t *testing.T) {
	result := decode([]rune("980A"), 2, make(memory))
	expect := 60
	if result != expect {
		t.Fatalf("Decode returns %d, we expect %d", result, expect)
	}
}

func TestComplexity1_2(t *testing.T) {
	result := complexity([]rune("980A"), 2, make(memory))
	expect := 60 * 980
	if result != expect {
		t.Fatalf("Decode returns %d, we expect %d", result, expect)
	}
}

func TestDecode1_3(t *testing.T) {
	result := decode([]rune("179A"), 2, make(memory))
	expect := 68
	if result != expect {
		t.Fatalf("Decode returns %d, we expect %d", result, expect)
	}
}

func TestComplexity1_3(t *testing.T) {
	result := complexity([]rune("179A"), 2, make(memory))
	expect := 68 * 179
	if result != expect {
		t.Fatalf("Decode returns %d, we expect %d", result, expect)
	}
}

func TestDecode1_4(t *testing.T) {
	result := decode([]rune("456A"), 2, make(memory))
	expect := 64
	if result != expect {
		t.Fatalf("Decode returns %d, we expect %d", result, expect)
	}
}

func TestComplexity1_4(t *testing.T) {
	result := complexity([]rune("456A"), 2, make(memory))
	expect := 64 * 456
	if result != expect {
		t.Fatalf("Decode returns %d, we expect %d", result, expect)
	}
}

func TestDecode1_5(t *testing.T) {
	result := decode([]rune("379A"), 2, make(memory))
	expect := 64
	if result != expect {
		t.Fatalf("Decode returns %d, we expect %d", result, expect)
	}
}

func TestComplexity1_5(t *testing.T) {
	result := complexity([]rune("379A"), 2, make(memory))
	expect := 64 * 379
	if result != expect {
		t.Fatalf("Decode returns %d, we expect %d", result, expect)
	}
}

func TestBackwardNum1_1(t *testing.T) {
	step1 := decodeNum([]rune("029A"))
	back, ok := backwardNum(step1)
	expect := "029A"
	if !ok || string(back) != expect {
		t.Fatalf("Backward returns %s, we expect %s", string(back), expect)
	}
}

func TestBackwardNum1_2(t *testing.T) {
	data := "980A"
	step1 := decodeNum([]rune(data))
	back, ok := backwardNum(step1)
	expect := data
	if !ok || string(back) != expect {
		t.Fatalf("Backward returns %s, we expect %s", string(back), expect)
	}
}

func TestBackwardNum1_3(t *testing.T) {
	data := "179A"
	step1 := decodeNum([]rune(data))
	back, ok := backwardNum(step1)
	expect := data
	if !ok || string(back) != expect {
		t.Fatalf("Backward returns %s, we expect %s", string(back), expect)
	}
}

func TestBackwardNum1_4(t *testing.T) {
	data := "456A"
	step1 := decodeNum([]rune(data))
	back, ok := backwardNum(step1)
	expect := data
	if !ok || string(back) != expect {
		t.Fatalf("Backward returns %s, we expect %s", string(back), expect)
	}
}

func TestBackwardNum1_5(t *testing.T) {
	data := "379A"
	step1 := decodeNum([]rune(data))
	back, ok := backwardNum(step1)
	expect := data
	if !ok || string(back) != expect {
		t.Fatalf("Backward returns %s, we expect %s", string(back), expect)
	}
}

func testData() []string {
	return []string{
		"029A",
		"980A",
		"179A",
		"456A",
		"379A",
	}
}
