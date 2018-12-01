package main

import "testing"

func TestDay1Part1Simple1(t *testing.T) {
	lines := []string{"+1", "-2", "+3", "+1"}
	result := Day1part1(lines)
	var expected = 3
	if result != expected {
		t.Fatalf("Expected %d, got %d", expected, result)
	}
}
func TestDay1Part1Simple2(t *testing.T) {
	lines := []string{"+1", "+1", "+1"}
	result := Day1part1(lines)
	var expected = 3
	if result != expected {
		t.Fatalf("Expected %d, got %d", expected, result)
	}
}
func TestDay1Part1Simple3(t *testing.T) {
	lines := []string{"+1", "+1", "-2"}
	result := Day1part1(lines)
	var expected = 0
	if result != expected {
		t.Fatalf("Expected %d, got %d", expected, result)
	}
}
func TestDay1Part1Simple4(t *testing.T) {
	lines := []string{"-1", "-2", "-3"}
	result := Day1part1(lines)
	var expected = -6
	if result != expected {
		t.Fatalf("Expected %d, got %d", expected, result)
	}
}
func TestDay1Part2Simple1(t *testing.T) {
	lines := []string{"+1", "-2", "+3", "+1"}
	result := Day1part2(lines)
	var expected = 2
	if result != expected {
		t.Fatalf("Expected %d, got %d", expected, result)
	}
}
func TestDay1Part2Simple2(t *testing.T) {
	lines := []string{"+1", "-1"}
	result := Day1part2(lines)
	var expected = 0
	if result != expected {
		t.Fatalf("Expected %d, got %d", expected, result)
	}
}
func TestDay1Part2Simple3(t *testing.T) {
	lines := []string{"+3", "+3", "+4", "-2", "-4"}
	result := Day1part2(lines)
	var expected = 10
	if result != expected {
		t.Fatalf("Expected %d, got %d", expected, result)
	}
}
func TestDay1Part2Simple4(t *testing.T) {
	lines := []string{"-6", "+3", "+8", "+5", "-6"}
	result := Day1part2(lines)
	var expected = 5
	if result != expected {
		t.Fatalf("Expected %d, got %d", expected, result)
	}
}
func TestDay1Part2Simple5(t *testing.T) {
	lines := []string{"+7", "+7", "-2", "-7", "-4"}
	result := Day1part2(lines)
	var expected = 14
	if result != expected {
		t.Fatalf("Expected %d, got %d", expected, result)
	}
}
