package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Fil0x/AdventOfCode/internal"
)

var funcMap = map[string]interface{}{
	"1.1": Day1part1,
	"1.2": Day1part2,
}

func main() {
	lines := filereader.ReadFile(os.Args[2])
	res := funcMap[os.Args[1]].(func([]string) int)(lines)
	fmt.Println(res)
}

// day1part1 outputs correct 459
func Day1part1(lines []string) int {
	sum := 0
	for _, freq := range lines {
		num, _ := strconv.Atoi(freq)
		sum += num
	}
	return sum
}

func Day1part2(lines []string) int {
	seen := make(map[int]bool)
	sum := 0
	seen[sum] = true
	for i := 0; true; i++ {
		num, _ := strconv.Atoi(lines[i%len(lines)])
		sum += num
		if seen[sum] {
			break
		}
		seen[sum] = true
	}
	return sum
}
