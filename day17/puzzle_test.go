package main

import "testing"

func TestPart1(t *testing.T) {
	lines := testData()
	result := part1(lines)
	expect := "4,6,3,5,6,3,5,2,1,0"
	if result != expect {
		t.Fatalf("Part1 returns %s, we expect %s", result, expect)
	}
}

func TestPart2(t *testing.T) {
	lines := testData2()
	result := part2(lines)
	expect := 117440
	if result != expect {
		t.Fatalf("Part2 returns %d, we expect %d", result, expect)
	}
}

func TestCompute1(t *testing.T) {
	_, b, _, _ := compute([]int{2, 6}, 0, 0, 0, 9, []int{})
	expect := 1
	if b != expect {
		t.Fatalf("TestCompute1 returns %d, we expect %d", b, expect)
	}
}

func TestCompute2(t *testing.T) {
	_, _, _, output := compute([]int{5, 0, 5, 1, 5, 4}, 0, 10, 0, 0, []int{})
	expect := "0,1,2"
	if display(output) != expect {
		t.Fatalf("TestCompute2 returns %d, we expect %s", output, expect)
	}
}

func TestCompute3(t *testing.T) {
	a, _, _, output := compute([]int{0, 1, 5, 4, 3, 0}, 0, 2024, 0, 0, []int{})
	expectOutput := "4,2,5,6,7,7,7,7,3,1,0"
	expectA := 0
	if display(output) != expectOutput || expectA != a {
		t.Fatalf("TestCompute3 returns %d, we expect %s", output, expectOutput)
	}
}

func TestCompute4(t *testing.T) {
	_, b, _, _ := compute([]int{1, 7}, 0, 0, 29, 0, []int{})
	expect := 26
	if b != expect {
		t.Fatalf("TestCompute4 returns %d, we expect %d", b, expect)
	}
}

func TestCompute5(t *testing.T) {
	_, b, _, _ := compute([]int{4, 0}, 0, 0, 2024, 43690, []int{})
	expect := 44354
	if b != expect {
		t.Fatalf("TestCompute5 returns %d, we expect %d", b, expect)
	}
}

func TestCompute6(t *testing.T) {
	_, _, _, output := compute([]int{2, 4, 1, 3, 7, 5, 0, 3, 1, 5, 4, 4, 5, 5, 3, 0}, 0, 6, 0, 0, []int{})
	expect := "0"
	if display(output) != expect {
		t.Fatalf("TestCompute6 returns %d, we expect %s", output, expect)
	}
}

func TestCompute7(t *testing.T) {
	_, _, _, output := compute([]int{2, 4, 1, 3, 7, 5, 0, 3, 1, 5, 4, 4, 5, 5, 3, 0}, 0, 0b110001001, 0, 0, []int{})
	expect := "5,3,0"
	if display(output) != expect {
		t.Fatalf("TestCompute7 returns %d, we expect %s", output, expect)
	}
}

func TestComputeBF1(t *testing.T) {
	result := computeBF([]int{0, 3, 5, 4, 3, 0}, 0, 117440, 0, 0, []int{})
	if !result {
		t.Fatalf("TestComputeBF1 returns false, we expect true obviously")
	}
}

func testData() []string {
	return []string{
		"Register A: 729",
		"Register B: 0",
		"Register C: 0",
		"",
		"Program: 0,1,5,4,3,0",
	}
}

func testData2() []string {
	return []string{
		"Register A: 2024",
		"Register B: 0",
		"Register C: 0",
		"",
		"Program: 0,3,5,4,3,0",
	}
}
