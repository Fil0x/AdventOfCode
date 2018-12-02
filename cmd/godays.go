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
	"2.1": Day2part1,
	"2.2": Day2part2,
}

func main() {
	lines := filereader.ReadFile(os.Args[2])
	resInt, resStr := funcMap[os.Args[1]].(func([]string) (int, string))(lines)
	fmt.Println(resInt, resStr)
}

func Day1part1(lines []string) (int, string) {
	sum := 0
	for _, freq := range lines {
		num, _ := strconv.Atoi(freq)
		sum += num
	}
	return sum, ""
}

func Day1part2(lines []string) (int, string) {
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
	return sum, ""
}

func Day2part1(lines []string) (int, string) {
	twoCount, threeCount := 0, 0
	for _, boxids := range lines {
		charMap := make(map[byte]int)
		candidates := make([]byte, 0)
		n := len(boxids)
		for i := 0; i < n; i++ {
			if _, ok := charMap[boxids[i]]; ok {
				charMap[boxids[i]]++
				candidates = append(candidates, boxids[i])
			} else {
				charMap[boxids[i]] = 1
			}
		}
		// loop through the candidates and find one exactly 2 and 3
		twoFound, threeFound := false, false
		for _, candidate := range candidates {
			if !twoFound && charMap[candidate] == 2 {
				twoCount++
				twoFound = true
			}
			if !threeFound && charMap[candidate] == 3 {
				threeCount++
				threeFound = true
			}
			if twoFound && threeFound {
				break
			}
		}
	}

	return twoCount * threeCount, ""
}

func Day2part2(lines []string) (int, string) {
	for i := 0; i < len(lines)-1; i++ {
		for j := i + 1; j < len(lines); j++ {
			differences, pos := diffStrings(lines[i], lines[j])
			if differences == 1 {
				return 0, (lines[i][:pos[0]] + lines[i][pos[0]+1:])
			}
		}
	}
	return 0, ""
}

func diffStrings(str1 string, str2 string) (int, []int) {
	counter := 0
	diffPositions := make([]int, 0)
	// str1 and str2 are same length
	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			counter++
			diffPositions = append(diffPositions, i)
		}
	}
	return counter, diffPositions
}
