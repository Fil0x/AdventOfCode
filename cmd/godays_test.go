package main

import "testing"

func TestDay1Part1Simple1(t *testing.T) {
	lines := []string{"+1", "-2", "+3", "+1"}
	rI, _ := Day1part1(lines)
	var expected = 3
	if rI != expected {
		t.Fatalf("Expected %d, got %d", expected, rI)
	}
}
func TestDay1Part1Simple2(t *testing.T) {
	lines := []string{"+1", "+1", "+1"}
	rI, _ := Day1part1(lines)
	var expected = 3
	if rI != expected {
		t.Fatalf("Expected %d, got %d", expected, rI)
	}
}
func TestDay1Part1Simple3(t *testing.T) {
	lines := []string{"+1", "+1", "-2"}
	rI, _ := Day1part1(lines)
	var expected = 0
	if rI != expected {
		t.Fatalf("Expected %d, got %d", expected, rI)
	}
}
func TestDay1Part1Simple4(t *testing.T) {
	lines := []string{"-1", "-2", "-3"}
	rI, _ := Day1part1(lines)
	var expected = -6
	if rI != expected {
		t.Fatalf("Expected %d, got %d", expected, rI)
	}
}
func TestDay1Part2Simple1(t *testing.T) {
	lines := []string{"+1", "-2", "+3", "+1"}
	rI, _ := Day1part2(lines)
	var expected = 2
	if rI != expected {
		t.Fatalf("Expected %d, got %d", expected, rI)
	}
}
func TestDay1Part2Simple2(t *testing.T) {
	lines := []string{"+1", "-1"}
	rI, _ := Day1part2(lines)
	var expected = 0
	if rI != expected {
		t.Fatalf("Expected %d, got %d", expected, rI)
	}
}
func TestDay1Part2Simple3(t *testing.T) {
	lines := []string{"+3", "+3", "+4", "-2", "-4"}
	rI, _ := Day1part2(lines)
	var expected = 10
	if rI != expected {
		t.Fatalf("Expected %d, got %d", expected, rI)
	}
}
func TestDay1Part2Simple4(t *testing.T) {
	lines := []string{"-6", "+3", "+8", "+5", "-6"}
	rI, _ := Day1part2(lines)
	var expected = 5
	if rI != expected {
		t.Fatalf("Expected %d, got %d", expected, rI)
	}
}
func TestDay1Part2Simple5(t *testing.T) {
	lines := []string{"+7", "+7", "-2", "-7", "-4"}
	rI, _ := Day1part2(lines)
	var expected = 14
	if rI != expected {
		t.Fatalf("Expected %d, got %d", expected, rI)
	}
}
func TestDay2Part1Simple(t *testing.T) {
	lines := []string{"abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"}
	rI, _ := Day2part1(lines)
	var expected = 12
	if rI != expected {
		t.Fatalf("Expected %d, got %d", expected, rI)
	}
}
func TestDay2Part2Simple(t *testing.T) {
	lines := []string{"abcde", "fghij", "klmno", "pqrst", "fguij", "axcye", "wvxyz"}
	_, rS := Day2part2(lines)
	var expected = "fgij"
	if rS != expected {
		t.Fatalf("Expected %s, got %s", expected, rS)
	}
}
